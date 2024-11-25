package main

import (
	"asynq/dto"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	// Buat server untuk menjalankan worker
	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: "localhost:6379",
		},
		asynq.Config{
			Concurrency: 10, // Jumlah worker yang berjalan secara paralel
		},
	)

	// Definisikan handler untuk tugas "charge_user"
	mux := asynq.NewServeMux()
	mux.HandleFunc("charge_user", func(ctx context.Context, task *asynq.Task) error {
		var p dto.DataSubscription
		if err := json.Unmarshal(task.Payload(), &p); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		}

		fmt.Println(p)
		
		return nil
	})

	// Jalankan worker
	if err := server.Run(mux); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
