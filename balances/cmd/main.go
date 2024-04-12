package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com.br/devfullcycle/fc-ms-balances/internal/database"
	"github.com.br/devfullcycle/fc-ms-balances/internal/event"
	"github.com.br/devfullcycle/fc-ms-balances/internal/event/handler"
	checkbalance "github.com.br/devfullcycle/fc-ms-balances/internal/usecase/check_balance"
	updatebalance "github.com.br/devfullcycle/fc-ms-balances/internal/usecase/update_balance"
	"github.com.br/devfullcycle/fc-ms-balances/internal/web"
	"github.com.br/devfullcycle/fc-ms-balances/internal/web/webserver"
	"github.com.br/devfullcycle/fc-ms-balances/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func kafkaConsumer(wg *sync.WaitGroup, updateBalanceUseCase *updatebalance.UpdateBalanceUseCase) {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	consumer, err := ckafka.NewConsumer(&configMap)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()
	err = consumer.SubscribeTopics([]string{"balances"}, nil)
	if err != nil {
		panic(err)
	}

	// Signal that the consumer goroutine is ready
	wg.Done()
	fmt.Println("Kafka consumer is running")

	// Process messages
	run := true
	for run {
		select {
		default:
			ev, err := consumer.ReadMessage(100 * time.Millisecond)

			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}

			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

			balanceUpdated := event.BalanceUpdated{}
			err2 := json.Unmarshal([]byte(ev.Value), &balanceUpdated)
			if err2 != nil {
				fmt.Println(err2)
				continue
			}

			var updateBalanceKafkaHandler = handler.NewUpdateBalanceKafkaHandler(&balanceUpdated, updateBalanceUseCase)
			updateBalanceKafkaHandler.Handle()

		}
	}

	fmt.Println("Kafka consumer finished")
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-balances", "3306", "balances"))
	if err != nil {
		panic(err)
	}

	defer db.Close()
	balanceDb := database.NewBalanceDB(db)
	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("BalanceDB", func(tx *sql.Tx) interface{} {
		return database.NewBalanceDB(db)
	})

	checkBalanceUseCase := checkbalance.NewCheckBalanceUseCase(balanceDb)
	updateBalanceUseCase := updatebalance.NewUpdateBalanceUseCase(balanceDb)

	// Create a WaitGroup to synchronize goroutines.
	wg := &sync.WaitGroup{}
	// Add 1 to the WaitGroup, indicating one goroutine to wait for.
	wg.Add(1)
	// Launch the Kafka consumer goroutine in the background, passing the WaitGroup for synchronization.
	go kafkaConsumer(wg, updateBalanceUseCase)
	// Wait for the consumer goroutine to be ready
	wg.Wait()

	webServer := webserver.NewWebServer(":8080")
	checkBalanceHandler := web.NewWebCheckBalanceHandler(*checkBalanceUseCase)
	webServer.AddHandler("/balances/{account_id}", checkBalanceHandler.CheckBalance)
	fmt.Println("Server is running")
	webServer.Start()
}
