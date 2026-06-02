package searcher_proxy

import (
	"errors"

	"github.com/web-rabis/searcher-proxy/client"
	"github.com/web-rabis/searcher-proxy/client/grpc"
	"github.com/web-rabis/searcher-proxy/model"
)

type FactoryMethod func(conf *model.ConnectionConfig) (client.Base, error)

var (
	factories = map[string]FactoryMethod{
		"grpc": grpc.NewClient,
	}
	ErrInvalidProtocol = errors.New("invalid protocol")
)

func NewSearcherProxyClient(config *model.ConnectionConfig) (client.Base, error) {
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
