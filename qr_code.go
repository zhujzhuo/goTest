package main

import qrcode "github.com/skip2/go-qrcode"
import (
    "fmt"
    "image/color"
)
//qrcode.Low   qrcode.Medium   qrcode.High   qrcode.Highest
func main() {
    err1 := qrcode.WriteColorFile("username   password", qrcode.Low, 256, color.Black, color.White,"qr1.png")
    err2 := qrcode.WriteColorFile("username   password", qrcode.Medium, 256, color.Black, color.White,"qr2.png")
    err3 := qrcode.WriteColorFile("username   password", qrcode.High, 256, color.Black, color.White,"qr3.png")
    err4 := qrcode.WriteColorFile("username   password", qrcode.Highest, 256, color.Black, color.White,"qr4.png")
    if err1 != nil {
        fmt.Println("write error")
    }
    if err2 != nil {
        fmt.Println("write error")
    }
    if err3 != nil {
        fmt.Println("write error")
    }
    if err4 != nil {
        fmt.Println("write error")
    }
}
