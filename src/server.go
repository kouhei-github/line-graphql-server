package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/meetup/graph/repository"
	"github.com/meetup/route"
)

func main() {
	repository.ConnectDatabase()

	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	route.GetRouter()
}
