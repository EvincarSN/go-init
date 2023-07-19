package main

import (
	"database/sql"
	"fmt"

	"github.com/EvincarSN/go-init/internal/infra/database"
	"github.com/EvincarSN/go-init/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

// type Car struct {
// 	Model string
// 	Color string
// }

// // Método
// func (c Car) Start() {
// 	println(c.Model + " has been started")
// }

// func (c *Car) ChangeColor(color string) {
// 	//Sem * cria cópia. com * altera o valor original na memória.
// 	c.Color = color //Duplicando o valor de c.color na memória. É uma cópia do color original
// 	println("New color: " + c.Color)
// }

// // Função
// func soma(x int, y int) int {
// 	return x + y
// }

func main() {
	//println("Hello, World!")

	// a := 10
	// b := a //copiou o valor de a, mas criou seu próprio espaço na memória
	// b = 20
	// println(&a)
	// println(a)
	// println(b)
	// println(a)

	// car := Car{ // Declarando e atribuindo a variável Car
	// 	Model: "Ferrari",
	// 	Color: "Red",
	// }

	// car.Start()
	// car.ChangeColor("Blue")
	// println(car.Color)

	// order, err := entity.NewOrder("1", 10, 1)
	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(order.ID)
	// }

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	//println(output)
	fmt.Println(output)
}
