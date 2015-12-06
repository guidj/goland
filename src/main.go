package main

import (
//	"encoding/json"
	"fmt"
//	"io"
//	"log"
//	"strings"
	"./oracle"
)

const jsonStream = `
    {"bvu62fu6dq": {
        "name": "john",
        "age": 23,
        "xyz": "weu33s"}
    }`

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Info map[string]Person

func Hello(fn func() string) {
	fmt.Println("Hello " + fn())
}


func main() {

//	dec := json.NewDecoder(strings.NewReader(jsonStream))
//	for {
//		var info Info
//		if err := dec.Decode(&info); err == io.EOF {
//			break
//		} else if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("%s: %d\n", info["bvu62fu6dq"].Name, info["bvu62fu6dq"].Age)
//	}

	Hello(oracle.Colombia)
	Hello(oracle.DotGo)
}
