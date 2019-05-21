package main

import (
	"context"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/melkor/test-thrift/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	addr       = pflag.StringP("addr", "h", "localhost:9090", "Address to listen to")
	protocol   = pflag.StringP("protocol", "p", "json", "Specify the protocol (binary, compact, json, simplejson)")
	number     = pflag.IntP("number", "c", 0, "Credit Card number")
	name       = pflag.StringP("name", "n", "", "Card owner's Name)")
	crypto     = pflag.IntP("crypto", "", 0, "Ctryptogram")
	date       = pflag.StringP("date", "d", "", "Expiration Date")
	defaultCtx = context.Background()
)

func main() {
	pflag.Parse()

	client := client.New(*addr, *protocol)

	err := client.Init()

	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	spew.Dump(client.Authorize(defaultCtx))

}
