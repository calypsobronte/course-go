package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
  randomQuote()
}

func randomQuote()  {
  fmt.Println(quote.Hello())
  fmt.Println(quote.Go())
  fmt.Println(quote.Glass())
}
