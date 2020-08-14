package main

import (
	"context"
	"log"
	"time"

	"github.com/ZergsLaw/ctx"
)

const timeout = time.Second

func main() {

	c, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res := make(chan string, 1)

	err := ctx.Go(c, func(c context.Context) error {
		str, err := LongRunningProcess()
		if err != nil {
			return err
		}

		res <- str
		return nil
	})

	if err != nil {
		log.Fatal("very long process ", err)
	}

}

func LongRunningProcess() (string, error) {
	time.Sleep(5 * time.Second)
	return "Hello, bro!", nil
}
