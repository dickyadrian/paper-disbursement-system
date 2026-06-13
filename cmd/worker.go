package cmd

import (
	"context"
	"fmt"

	"github.com/dickyadrian/paper-disbursement-system/application"
	"github.com/dickyadrian/paper-disbursement-system/internal/task"
	"github.com/hibiken/asynq"
)

func RunWorker(ctx context.Context, app *application.App) {
	srv := asynq.NewServer(app.AsynqRedisOpt, asynq.Config{
		Concurrency: 10,
	})

	mux := asynq.NewServeMux()

	th := task.NewHandler(app)
	mux.HandleFunc(task.TypeDisbursementProcess, th.ProcessDisbursement)

	if err := srv.Start(mux); err != nil {
		fmt.Printf("Error starting worker: %v\n", err)
		return
	}

	<-ctx.Done()
	srv.Shutdown()
	fmt.Printf("Worker shut down")
}
