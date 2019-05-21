package main

import "github.com/spf13/pflag"

var (
	addr     = pflag.StringP("addr", "h", "localhost:9090", "Address to listen to")
	protocol = pflag.StringP("protocol", "p", "json", "Specify the protocol (binary, compact, json, simplejson)")
)

func main() {

}
