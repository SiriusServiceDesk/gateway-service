package main

import (
	"context"
	"github.com/SiriusServiceDesk/gateway-service/internal/config"
	"github.com/SiriusServiceDesk/gateway-service/internal/initializers"
	"log"
	"sync"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfig()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := initializers.StartGrpcServer(cfg); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()

		if err := initializers.StartHttpServer(cfg, ctx); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
