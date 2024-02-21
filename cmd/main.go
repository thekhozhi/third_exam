package main

import (
	"context"
	"develop/api"
	"develop/config"
	"develop/service"
	"develop/storage/postgres"
	"fmt"
	"log"
)

func main(){
	cfg := config.Load()

	pgStore, err := postgres.New(context.Background(),cfg)
	if err != nil{
		log.Fatalln("Error while connecting to db!", err.Error())
		return
	}

	defer pgStore.Close()

	services := service.New(pgStore)
	server := api.New(services)

	err = server.Run("localhost:8080")
	if err != nil{
		fmt.Println("Server is not running!",err.Error())
		return
	}
}