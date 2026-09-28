package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/web-rabis/ebook-client/client"
	"github.com/web-rabis/ebook-client/model"
	"github.com/web-rabis/ebook-client/protobuf"
)

type BaseClient struct {
	address  string
	conn     *grpc.ClientConn
	dialOpts []grpc.DialOption

	ebookSvc     client.EbookService
	ebookFormSvc client.EbookFormService
	dictSvc      client.DictionaryService
	worksheetSvc client.WorksheetService
}

var _ client.Base = &BaseClient{}

func NewClient(config *model.ConnectionConfig) (client.Base, error) {
	var grpcOpts []grpc.DialOption
	if config.Insecure == true {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &BaseClient{
		address:  config.Address,
		dialOpts: grpcOpts,
	}, nil
}

func (c *BaseClient) Connect() (err error) {
	c.conn, err = grpc.Dial(c.address, c.dialOpts...)
	if err != nil {
		return err
	}

	return err
}

func (c *BaseClient) Close() error {
	return c.conn.Close()
}

func (c *BaseClient) EbookSvc() client.EbookService {
	if c.ebookSvc == nil {
		c.ebookSvc = NewEbookServiceClient(protobuf.NewEbookSvcClient(c.conn))
	}
	return c.ebookSvc
}
func (c *BaseClient) EbookFormSvc() client.EbookFormService {
	if c.ebookFormSvc == nil {
		c.ebookFormSvc = NewEbookFormServiceClient(protobuf.NewEbookFormSvcClient(c.conn))
	}
	return c.ebookFormSvc
}
func (c *BaseClient) DictionarySvc() client.DictionaryService {
	if c.dictSvc == nil {
		c.dictSvc = NewDictionaryServiceClient(protobuf.NewDictionarySvcClient(c.conn))
	}
	return c.dictSvc
}
func (c *BaseClient) WorksheetSvc() client.WorksheetService {
	if c.worksheetSvc == nil {
		c.worksheetSvc = NewWorksheetServiceClient(protobuf.NewWorksheetSvcClient(c.conn))
	}
	return c.worksheetSvc
}
