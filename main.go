package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
)

func main() {
	err := qrcode.WriteColorFile("qr code in golang", qrcode.Medium, 256, color.Black, color.White, "myfile.png")
	if err != nil {
		fmt.Printf("sorry we coldnt create qrcode:%v", err)
	}
}
