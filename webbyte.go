package main

import (
    wb "github.com/NovusEdge/web-byte/src"
)

func main() {
    bs, _ := wb.MakeByteServer(8080, nil)

    bs.AddHandler("/", wb.DefaultServe)
    bs.Init()
    log.Fatal(bs.Serve())
}
