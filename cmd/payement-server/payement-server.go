package main

import (
	"fmt"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/melkor/test-thrift/server"
	"github.com/spf13/pflag"
)

var (
	addr     = pflag.StringP("addr", "h", "localhost:9090", "Address to listen to")
	protocol = pflag.StringP("protocol", "p", "json", "Specify the protocol (binary, compact, json, simplejson)")
)

func main() {
	pflag.Parse()

	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	if err := server.Run(transportFactory, protocolFactory, *addr); err != nil {
		fmt.Println("error running server:", err)
	}
}
