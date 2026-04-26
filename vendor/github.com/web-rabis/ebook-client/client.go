package order_client

import (
	"errors"

	"github.com/web-rabis/ebook-client/client"
	"github.com/web-rabis/ebook-client/client/grpc"
	"github.com/web-rabis/ebook-client/model"
)

type FactoryMethod func(conf *model.ConnectionConfig) (client.Base, error)

var (
	factories = map[string]FactoryMethod{
		"grpc": grpc.NewClient,
	}
	ErrInvalidProtocol = errors.New("invalid protocol")
)

func NewEbookClient(config *model.ConnectionConfig) (client.Base, error) {
	factory, ok := factories[config.Protocol]
	if !ok {
		return nil, ErrInvalidProtocol
	}

	c, err := factory(config)
	if err != nil {
		return nil, err
	}

	err = c.Connect()

	if err != nil {
		return nil, err
	}

	return c, nil
}
