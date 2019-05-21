package main

import (
	"context"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/melkor/test-thrift/client"
	"github.com/melkor/test-thrift/payment"
)

var (
	addr       = pflag.StringP("addr", "h", "localhost:9090", "Address to listen to")
	protocol   = pflag.StringP("protocol", "p", "json", "Specify the protocol (binary, compact, json, simplejson)")
	number     = pflag.Int64P("number", "c", 0, "Credit Card number")
	name       = pflag.StringP("name", "n", "", "Card owner's Name)")
	crypto     = pflag.Int32P("crypto", "", 0, "Ctryptogram")
	date       = pflag.StringP("date", "d", "", "Expiration Date")
	defaultCtx = context.Background()
)

const layout = "2006/01/02"

func main() {
	pflag.Parse()

	client := client.New(*addr, *protocol)

	err := client.Init()

	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	t, err := time.Parse(layout, *date)

	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	client.Card = payment.CreditCard{
		Number:     *number,
		Cryptogram: *crypto,
		Holder:     *name,
		Date:       t.Unix(),
	}

	spew.Dump(client.Authorize(defaultCtx))

}
