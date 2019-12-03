package main

import (
	"os"
	"fmt"
	"time"
	"context"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:	[]string{"http://localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Get(ctx, "/pod")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%q\n",resp.Kvs)
	fmt.Printf("%T\n", cancel)
}

