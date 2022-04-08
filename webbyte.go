package main

import (
	wb "github.com/NovusEdge/web-byte/src"
	"log"
)

func main() {
	bs, _ := wb.MakeByteServer(8080, nil)

	bs.AddHandler("/", wb.DefaultServe)
	log.Fatal(bs.Serve())
}
