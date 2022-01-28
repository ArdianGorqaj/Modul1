package main

import (
	"example.com/ardian/Rivercrossing"
	"fmt"
)
import "example.com/ardian/MyQuote"

func main() {
	fmt.Println(MyQuote.GetGlass())
	fmt.Println(MyQuote.GetGo())
	fmt.Println(MyQuote.GetOpt())
	fmt.Println(MyQuote.GetHello())
	fmt.Println(Rivercrossing.ViewState())
	Rivercrossing.Putinboat()
	fmt.Println(Rivercrossing.ViewState())
	Rivercrossing.Putinrev()
	fmt.Println(Rivercrossing.ViewState())

}
