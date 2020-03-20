package main

import (
	"github.com/Moisesbr/gotp/escpos"
)

func simpleprint() {
	p := escpos.New(false, "/dev/ttyUSB0", 9600)
	p.Begin()
	p.SetCodePage("PC437")
	p.SetAlign("center")
	p.SetFontSize("medium")
	// p.SetBold(true)
	p.WriteText("Seja bem vindo ao condominio")
	p.Linefeed()
	p.WriteText("Vila Di Capri")
	p.Linefeed()
	p.Linefeed()
	p.SetBarcodeHeight(100)
	p.BarCode("EAN13", "1234567")
	p.Linefeed()
	p.Linefeed()
	p.SetFontSize("small")
	p.WriteText("3SR Sistemas e Automacao.")
	p.Linefeed()
	p.WriteText("Tel: (16) xxxx-xxxx")
	count := 6
	for i := 0; i < count; i++ {
		p.Linefeed()
	}
	p.Cut()
}

func main() {
	simpleprint()
}
