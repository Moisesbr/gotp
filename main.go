package main

import (
	"github.com/Moisesbr/gotp/escpos"
)

func simpleprint() {
	p := escpos.New(false, "/dev/ttyUSB0", 9600)
	// p := escpos.New(false, "/dev/ttyS101", 9600)
	p.Begin()
	// p.SetCodePage("PC437")
	// p.Linefeed()
	p.SetAlign("center")
	p.SetFontSize("medium")
	// p.SetBold(true)
	p.WriteText("Seja bem vindo ao condominio")
	p.Linefeed()
	p.WriteText("Vila Di Capri")
	p.Linefeed()
	p.Linefeed()
	p.SetBarcodeHeight(100)
	// p.Debug = false
	// p.WriteText("teste123456789")
	p.BarCode("EAN13", "1234567")
	p.Linefeed()
	p.Linefeed()
	p.Linefeed()
	p.Linefeed()
	p.Linefeed()
	p.Linefeed()
	// p.Feed(16)
	p.Cut()
}

func main() {
	simpleprint()
}
