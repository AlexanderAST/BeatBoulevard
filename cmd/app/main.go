package main

import (
	"BeatBoulevard/internal/handlers"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env %v", err.Error())
	}
	fmt.Println("hello world")

	if err := handlers.Start(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
