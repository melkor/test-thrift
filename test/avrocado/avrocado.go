package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	avro "github.com/leboncoin/avrocado"
)

type Someone struct {
	Name string `avro:"name"`
	Age  int32  `avro:"age"`
}

func main() {
	val := Someone{"MyName", 3}
	var decoded Someone

	schema := `{
          "type": "record",
          "name": "Someone",
          "fields": [
            {
              "name": "name",
              "type": "string"
            }, {
              "name": "age",
              "type": "int"
            }
          ]
        }`

	codec, err := avro.NewCodec(schema)
	if err != nil {
		panic(fmt.Sprintf("wrong schema: %s", err))
	}

	avrovro, err := codec.Marshal(&val)
	if err != nil {
		panic(fmt.Sprintf("unable to serialize to avro: %s", err))
	}

	spew.Dump(val)

	err = codec.Unmarshal(avrovro, &decoded)
	if err != nil {
		panic(fmt.Sprintf("unable to deserialize from avro: %s", err))
	}
	spew.Dump(avrovro)
	spew.Dump(decoded)

}
