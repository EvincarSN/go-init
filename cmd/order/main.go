package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/EvincarSN/go-init/internal/infra/database"
	"github.com/EvincarSN/go-init/internal/usecase"
	"github.com/EvincarSN/go-init/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	defer db.Close() // defer espera tudo rodar depois executa o close
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) //escutando a fila e trava por isso fazer um T2
	rabbitmqWorker(msgRabbitmqChannel, uc)      //T1

	// input := usecase.OrderInput{
	// 	ID:    "123",
	// 	Price: 10.0,
	// 	Tax:   1.0,
	// }
	// output, err := uc.Execute(input)
	// if err != nil {
	// 	panic(err)
	// }
	// //println(output)
	// fmt.Println(output)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Mensagem processada e salva no banco:", output)
	}
}
