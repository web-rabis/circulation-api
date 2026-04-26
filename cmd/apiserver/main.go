package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/logutils"
	"github.com/web-rabis/circulation-api/internal/domain/manager/ebook"

	"golang.org/x/sync/errgroup"

	"github.com/web-rabis/circulation-api/internal/config"
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/dictionary"
	"github.com/web-rabis/circulation-api/internal/domain/manager/order"
	"github.com/web-rabis/circulation-api/internal/server/http"
	"github.com/web-rabis/db"
	ebookCli "github.com/web-rabis/ebook-client"
	ebookModel "github.com/web-rabis/ebook-client/model"
	orderCli "github.com/web-rabis/order-client"
	orderModel "github.com/web-rabis/order-client/model"
	readerCli "github.com/web-rabis/reader-client"
	readerModel "github.com/web-rabis/reader-client/model"
	ssoCli "github.com/web-rabis/sso-client"
	ssoModel "github.com/web-rabis/sso-client/model"
)

var (
	version = "unknown"
)

func main() {
	fmt.Printf("reader %s\n", version)

	opts := config.Parse(new(config.APIServer)).(*config.APIServer)

	setupLogUtils(opts.Dbg)

	appCtx, appCtxCancel := context.WithCancel(context.Background())
	defer appCtxCancel()

	go catchForTermination(appCtxCancel, os.Interrupt, syscall.SIGTERM)

	userDs, err := db.SetupUserDatabase(appCtx, opts.DSURL, opts.DSName, opts.DSDB)
	if err != nil {
		log.Println(err)
		return
	}
	defer userDs.Close(appCtx)

	authMan := auth.NewAuthManager(
		[]byte(opts.JWTKey),
		time.Minute*time.Duration(opts.AccessTokenTTL),
		time.Hour*time.Duration(opts.RefreshTokenTTL))
	orderGrpcLient, err := orderCli.NewOrderClient(&orderModel.ConnectionConfig{
		Address:  opts.OrderConfig.GrpcAddress,
		Protocol: "grpc",
		Insecure: true,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("[INFO] starting order grpc listener")
	readerGrpcLient, err := readerCli.NewReaderClient(&readerModel.ConnectionConfig{
		Address:  opts.ReaderConfig.GrpcAddress,
		Protocol: "grpc",
		Insecure: true,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("[INFO] starting reader grpc listener")
	ebookGrpcLient, err := ebookCli.NewEbookClient(&ebookModel.ConnectionConfig{
		Address:  opts.EbookConfig.GrpcAddress,
		Protocol: "grpc",
		Insecure: true,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("[INFO] starting ebook grpc listener")
	ssoGrpcClient, err := ssoCli.NewSsoClient(&ssoModel.ConnectionConfig{
		Address:  opts.SsoConfig.GrpcAddress,
		Protocol: "grpc",
		Insecure: true,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("[INFO] starting sso grpc listener")
	orderMan := order.NewOrderManager(orderGrpcLient.Order(), readerGrpcLient.ReaderSvc(), ssoGrpcClient.User(), ebookGrpcLient.EbookSvc())
	dictMan := dictionary.NewManager(orderGrpcLient.Dictionary())
	ebookMan := ebook.NewManager(ebookGrpcLient.EbookSvc())

	servers, serversCtx := errgroup.WithContext(appCtx)

	servers.Go(func() error {
		return http.Run(serversCtx, opts, authMan, orderMan, dictMan, ebookMan, ssoGrpcClient.User(), version)
	})

	if err := servers.Wait(); err != nil {
		log.Printf("[INFO] process terminated, %s", err)
		return
	}

}

func setupLogUtils(inDebugMode bool) {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("INFO"),
		Writer:   os.Stdout,
	}

	log.SetFlags(log.Ldate | log.Ltime)

	if inDebugMode {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		filter.MinLevel = "DEBUG"
	}

	log.SetOutput(filter)
}

func catchForTermination(cancel context.CancelFunc, signals ...os.Signal) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, signals...)
	<-stop
	log.Print("[WARN] interrupt signal")
	cancel()
}
