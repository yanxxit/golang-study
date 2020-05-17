package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// go get github.com/boombuler/barcode
// https://github.com/boombuler/barcode
func main() {

	qrCode, _ := qr.Encode("http://blog.csdn.net/wangshubo1989", qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create("qr2.png")
	defer file.Close()

	png.Encode(file, qrCode)
}
