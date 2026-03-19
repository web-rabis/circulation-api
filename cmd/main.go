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
	"github.com/web-rabis/circulation-api/internal/adapter/database/drivers"
	"github.com/web-rabis/circulation-api/internal/adapter/database/drivers/pgsql"
	"github.com/web-rabis/circulation-api/internal/config"
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/order"
	"github.com/web-rabis/circulation-api/internal/domain/manager/user"
	"github.com/web-rabis/circulation-api/internal/server/http"
	"github.com/web-rabis/db"
	dbuserman "github.com/web-rabis/db/user/manager/user"
	orderCli "github.com/web-rabis/order-client"
	orderModel "github.com/web-rabis/order-client/model"
	"golang.org/x/sync/errgroup"
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

	ds, err := pgsql.New(drivers.DataStoreConfig(drivers.DataStoreConfig{
		//URL: "postgres://postgres:nlrk$postgres@192.168.7.241:5432/nlrk",
		//URL:           "postgres://postgres:postgres@localhost:5432/nlrk",
		//URL:           "postgres://postgres:darmenkus@localhost:5432/nlrk",
		URL:           opts.DSURL,
		DataStoreName: opts.DSName,
		DataBaseName:  opts.DSDB,
	}))
	if err != nil {
		panic(err)
	}

	err = ds.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	authMan := auth.NewAuthManager(
		[]byte(opts.JWTKey),
		time.Minute*time.Duration(opts.AccessTokenTTL),
		time.Hour*time.Duration(opts.RefreshTokenTTL))
	userMan := user.NewUserManager(dbuserman.NewManager(userDs), authMan)
	orderGrpcLient, err := orderCli.NewOrderClient(&orderModel.ConnectionConfig{
		Address:  ":4000",
		Protocol: "grpc",
		Insecure: true,
	})
	if err != nil {
		panic(err)
	}
	orderMan := order.NewOrderManager(orderGrpcLient.Order())

	servers, serversCtx := errgroup.WithContext(appCtx)

	servers.Go(func() error {
		return http.Run(serversCtx, opts, authMan, userMan, orderMan, version)
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
