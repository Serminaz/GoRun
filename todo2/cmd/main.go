package main

import (
	"github.com/Serminaz/GoRun/todo2"
	"github.com/Serminaz/GoRun/todo2/pkg/handler"
	"github.com/Serminaz/GoRun/todo2/pkg/repository"
	"github.com/Serminaz/GoRun/todo2/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		//если функция возвр ошибку, то прерываем приложение
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)    //сервис зависит от репозитория
	handlers := handler.NewHandler(services) // хандлер от сервиса

	srv := new(todo.Server) // экземпляр сервера
	//в качестве значение порта можно передать значение из вайпера по ключу
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error { //для иницил  конфиг файлов
	viper.AddConfigPath("configs") // имч директории
	viper.SetConfigName("config")  // имя файла
	return viper.ReadInConfig()
	//считывает значение конфигов и записывает их во внутр объект вайпера
	// возвр ошибку
}
