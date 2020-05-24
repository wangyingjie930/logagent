package main

import (
	"fmt"
	"logagent/kafka"
	"logagent/tailLog"
	"time"
)

func run() {
	for  {
		select {
			case line := <-tailLog.ReadChan():
				kafka.SendToKafka("web_log", line.Text)
			default:
				time.Sleep(time.Second)
		}
	}
}

func main() {
	err := kafka.Init ([]string {"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf ("Iinit Kafka failed, err: %v\n", err)
		return
	}

	err = tailLog.Init("./my.log")
	if err != nil {
		fmt.Printf ("Iinit tailLog failed, err: %v\n", err)
		return
	}
	run()
}
