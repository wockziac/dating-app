package main

import (
	"datingapp/infrastructure/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	http.Setup(r)

	r.Run()
	fmt.Println("Hello")
}
