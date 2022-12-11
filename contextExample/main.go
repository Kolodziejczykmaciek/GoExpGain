package main

import (
	"context"
	"fmt"
	"log"
	"time"

)

type Response struct {
	value int
	err error
}

func fetchData(ctx context.Context, ID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := fetchRemoteData()
		respch <- Response{ 
			value: val,
			err: err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data took too long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchRemoteData() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 420, nil
}

func main() {
	start := time.Now()
	ctx := context.Background()
	ID := 40
	val, err := fetchData(ctx, ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))

}