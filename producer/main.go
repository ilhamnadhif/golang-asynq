package main

import (
	"asynq/dto"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	// Buat koneksi ke Redis
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
	})
	defer client.Close()

	// Data yang akan diproses

	payload, err := json.Marshal(dto.DataSubscription{
		Msisdn: "234234",
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
