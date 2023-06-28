package main

import (
	"fmt"

	"github.com/luanaltino/fc-example-utils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	err = rabbitmq.Publish(ch, "Hello World!", "amq.direct")
	if err != nil {
		fmt.Println(err.Error())
	}

}
