package main

import (
	"fmt"
	"os"
	"time"
	"context"

	"go.etcd.io/etcd/client"
)

func main() {
	cfg := client.Config{
		Endpoints:	[]string{"http://localhost:2379"},
		Transport:	client.DefaultTransport,
		HeaderTimeoutPerRequest:	time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	api := client.NewKeysAPI(c)

	fmt.Println("---Set---")
	resp, err := api.Set(context.Background(), "test", "minato", nil)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("metadata : %q\n",resp)
	}

	fmt.Println("---Get---")
	resp, err = api.Get(context.Background(), "test", nil)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%q\n", resp)
		fmt.Printf("%q : %q\n", resp.Node.Key,resp.Node.Value)
	}
}
