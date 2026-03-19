package db

import (
	"context"
	"errors"
	"fmt"

	ebookDrv "github.com/web-rabis/db/internal/adapter/database/ebook/drivers"
	ebookPgsql "github.com/web-rabis/db/internal/adapter/database/ebook/drivers/pgsql"
	orderDrv "github.com/web-rabis/db/internal/adapter/database/eorder/drivers"
	orderPgsql "github.com/web-rabis/db/internal/adapter/database/eorder/drivers/pgsql"
	readerDrv "github.com/web-rabis/db/internal/adapter/database/reader/drivers"
	readerPgsql "github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql"
	userDrv "github.com/web-rabis/db/internal/adapter/database/user/drivers"
	userPgsql "github.com/web-rabis/db/internal/adapter/database/user/drivers/pgsql"
)

func SetupReaderDatabase(ctx context.Context, url, dsName, dbName string) (readerDrv.DataStore, error) {
	ds, err := readerPgsql.New(readerDrv.DataStoreConfig{
		URL:           url,
		DataStoreName: dsName,
		DataBaseName:  dbName,
	})
	if err != nil {
		return nil, err
	}

	if err := ds.Connect(ctx); err != nil {
		errText := fmt.Sprintf("[ERROR] cannot connect to datastore %s: %v", dsName, err)
		return nil, errors.New(errText)
	}

	fmt.Println("Connected to", ds.Name())

	return ds, nil
}
func SetupUserDatabase(ctx context.Context, url, dsName, dbName string) (userDrv.DataStore, error) {
	ds, err := userPgsql.New(userDrv.DataStoreConfig{
		URL:           url,
		DataStoreName: dsName,
		DataBaseName:  dbName,
	})
	if err != nil {
		return nil, err
	}

	if err := ds.Connect(ctx); err != nil {
		errText := fmt.Sprintf("[ERROR] cannot connect to datastore %s: %v", dsName, err)
		return nil, errors.New(errText)
	}

	fmt.Println("Connected to", ds.Name())

	return ds, nil
}
func SetupEOrderDatabase(ctx context.Context, url, dsName, dbName string) (orderDrv.DataStore, error) {
	ds, err := orderPgsql.New(orderDrv.DataStoreConfig{
		URL:           url,
		DataStoreName: dsName,
		DataBaseName:  dbName,
	})
	if err != nil {
		return nil, err
	}

	if err := ds.Connect(ctx); err != nil {
		errText := fmt.Sprintf("[ERROR] cannot connect to datastore %s: %v", dsName, err)
		return nil, errors.New(errText)
	}

	fmt.Println("Connected to", ds.Name())

	return ds, nil
}
func SetupEbookDatabase(ctx context.Context, url, dsName, dbName string) (ebookDrv.DataStore, error) {
	ds, err := ebookPgsql.New(ebookDrv.DataStoreConfig{
		URL:           url,
		DataStoreName: dsName,
		DataBaseName:  dbName,
	})
	if err != nil {
		return nil, err
	}

	if err := ds.Connect(ctx); err != nil {
		errText := fmt.Sprintf("[ERROR] cannot connect to datastore %s: %v", dsName, err)
		return nil, errors.New(errText)
	}

	fmt.Println("Connected to", ds.Name())

	return ds, nil
}
