package main

import (
	"database/sql"
	"fmt"

	"github.com/CarlosVinicius3258/go-calculate-tax/internal/infra/database"
	"github.com/CarlosVinicius3258/go-calculate-tax/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main(){
	db, err := sql.Open("sqlite3","db.sqlite3")
	if err != nil {
		panic(err)
	} 
	order_repo := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(order_repo)
	input:= usecase.OrderInput{
		ID: "1235",
		Price: 10,
		Tax: 1,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Print(output)
}