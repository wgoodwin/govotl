package main

import (
    "fmt"
    "log"

    "github.com/wgoodwin/govotl"
)

func main() {
    vd, err := govotl.NewVOTLDoc("../test.otl")
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Printf("%+v\n", vd)
}

