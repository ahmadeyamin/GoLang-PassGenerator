package main

import (
	"math/rand"
	"strings"
	"time"

	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/pass", GeneratePassword)

	r.Run(":8080")

}

func GeneratePassword(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" + "@$%&#@%$$^&$")

	var p = strings.Builder{}

	for i := 0; i < 20; i++ {
		p.WriteRune(chars[rand.Intn(len(chars))])
	}

	ctx.String(200, p.String())
}
