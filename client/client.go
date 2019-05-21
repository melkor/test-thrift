package client

import (
	"errors"

	//"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/melkor/test-thrift/payment"
)

type Client struct {
	Addr            string
	Card            payment.CreditCard
	Protocol        string
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
}

func New(addr, proto string) *Client {
	return &Client{
		Addr:     addr,
		Protocol: proto,
	}
}

func (c *Client) Init(transportFactory thrift.TTransportFactory) error {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket(c.Addr)

	if err != nil {
		return err
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}

	switch c.Protocol {
	case "compact":
		c.ProtocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		c.ProtocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		c.ProtocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		c.ProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		return errors.New("Invalid protocol specified: " + c.Protocol)
	}

	c.Transport = transport

	return nil
}
