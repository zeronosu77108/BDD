package main

import (
    "fmt"
)

type bdd struct {
    one *bdd
    zero *bdd
}

func main() {
    bdd := new(bdd)
    bdd.one = nil
    bdd.zero = nil
    fmt.Println("test")
}
