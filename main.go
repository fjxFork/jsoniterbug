package main

import (
	encjson "encoding/json"
	"fmt"
	"os"

	"github.com/json-iterator/go"
)

type str struct {
	F func()
}

func main() {
	var s *str = nil

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if _, err := json.Marshal(s); err != nil {
		fmt.Fprintf(os.Stderr, "json-iterator: marshalling failed: %s\n", err)
	} else {
		fmt.Println("json-iterator: marshalling succeeded.")
	}

	if _, err := encjson.Marshal(s); err != nil {
		fmt.Fprintf(os.Stderr, "encoding/json: marshalling failed: %s\n", err)
	} else {
		fmt.Println("encoding/json: marshalling succeeded.")
	}
}
