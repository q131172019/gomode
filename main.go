package main

import (
  "github.com/fatih/color"
  "rsc.io/quote"
  "github.com/q131172019/hellomod"
)

func main() {
  color.Cyan(quote.Hello())
  color.Red(hellomod.Hello())
}
