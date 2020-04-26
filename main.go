package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sergrom/cached/config"
	"log"
)

// protoc -I=./pb --go_out=plugins=grpc:./pb ./pb/cached.proto

func main()  {
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "The config yaml file")
	flag.Parse()

	conf := &config.Config{}
	err := conf.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Error occurred while loading config file \"%s\" %s", configFile, err.Error())
	}

	cachedServer = &Cached{}
	cachedServer.Start(conf)
}