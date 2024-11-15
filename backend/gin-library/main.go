package main

import (
	"ginius/gin-library/router"
)

func main() {
	r := router.SetupRouter()
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
