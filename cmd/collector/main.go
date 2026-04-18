package main

import (
	"context"
	"fmt"

	"github.com/B0GuS1/lol-winrate-iron/internal/api"
)

const (
	apiKey = "RGAPI-12edfc6b-56ea-442a-aefe-24724a79f49b"
)

func main() {
	// Создаем клиент
	client := api.NewRiotClient(apiKey)
	fmt.Println(client)

	ctx := context.Background()

	client.GetAllIron4Players(ctx)

}
