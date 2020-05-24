package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"logagent/conf"
	"logagent/kafka"
	"logagent/tailLog"
	"time"
)

var (
	cfg = new(conf.AppConfig) //error: cfg *conf.AppConfig
)

func run() {
	for  {
		select {
			case line := <-tailLog.ReadChan():
				kafka.SendToKafka(cfg.Topic, line.Text)
			default:
				time.Sleep(time.Second)
		}
	}
}

func main() {
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf ("load ini failed, err: %v\n", err)
		return
	}

	err = kafka.Init ([]string {cfg.Address})
	if err != nil {
		fmt.Printf ("Iinit Kafka failed, err: %v\n", err)
		return
	}

	err = tailLog.Init(cfg.FileName)
	if err != nil {
		fmt.Printf ("Iinit tailLog failed, err: %v\n", err)
		return
	}
	run()
}
