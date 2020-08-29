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

	res := ""

	err := ctxutils.Go(ctx, func() error {
		var err error
		res, err = LongRunningProcess()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal("very long process ", err)
	}

	log.Println(res)
}

func LongRunningProcess() (string, error) {
	time.Sleep(5 * time.Second)
	return "Hello, bro!", nil
}
