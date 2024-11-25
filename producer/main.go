package main

import (
	"asynq/dto"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	// Buat koneksi ke Redis

	if len(os.Args) < 2 {
		fmt.Println("Harap masukkan parameter tambahan.")
		return
	}
	param := os.Args[1]

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
	})
	defer client.Close()

	// Data yang akan diproses

	payload, err := json.Marshal(dto.DataSubscription{
		Msisdn: param,
	})
	if err != nil {
		panic(err)
	}

	// Jadwalkan tugas untuk diproses 7 hari kemudian
	task := asynq.NewTask("charge_user", payload)

	// Menjadwalkan tugas menggunakan ProcessIn
	_, err = client.Enqueue(task, asynq.ProcessIn(5*time.Second)) // 7 hari
	if err != nil {
		panic(err)
	}

	fmt.Println("Tugas berhasil dijadwalkan untuk diproses 7 hari kemudian!")
}
