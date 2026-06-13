package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/dickyadrian/paper-disbursement-system/application"
	"github.com/dickyadrian/paper-disbursement-system/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	app, err := application.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Go(func() {
		cmd.RunServer(ctx, app)
	})
	wg.Go(func() {
		cmd.RunWorker(ctx, app)
	})

	wg.Wait()
	cancel()
}
