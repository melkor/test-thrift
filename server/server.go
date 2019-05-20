package server

import (
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/melkor/test-thrift/payment"
)

//Run function launch server
func Run(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error

	// TODO: SSL
	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	handler := NewPaymentHandler()
	processor := payment.NewPaymentProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
