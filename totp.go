package main

import (
	"image/png"
	"log"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@exmaple.com",
	})
	f, _ := os.Create("qr.png")
	img, _ := key.Image(200, 200)
	png.Encode(f, img)

	code, _ := totp.GenerateCode(key.Secret(), time.Now())
	log.Println("code", code)

	valid := totp.Validate(code, key.Secret())
	if valid {
		log.Println("success")
	}
}
