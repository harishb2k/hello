package hello

import (
    "fmt"
    "rsc.io/quote"
)

func SayHi() {
    fmt.Println("This is change 1 on version 0")
    fmt.Println(quote.Hello())
}
