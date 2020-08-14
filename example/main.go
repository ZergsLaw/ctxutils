package main

import (
	"context"
	"log"
	"time"

	"github.com/ZergsLaw/ctxutils"
)

const timeout = time.Second

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res := make(chan string, 1)

	err := ctxutils.Go(ctx, func(ctx context.Context) error {
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
