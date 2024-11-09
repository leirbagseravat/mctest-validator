package main

import (
	"fmt"
	"mctest-agent/internal/http"
)

func main() {
	err := http.Init()

	if err != nil {
		fmt.Println(err)
	}
}
