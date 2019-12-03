package app

import (
	//"os"
	"fmt"
	"time"
	"context"

	"go.etcd.io/etcd/clientv3"
)

func NewEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:	[]string{"http://localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func Putkvs(cli *clientv3.Client, containername, reqvalue string) error {
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Put(ctx, containername, reqvalue)
	cancel()
	if err != nil {
		return err
	}
	fmt.Println(resp)

	resp2, err := cli.Get(ctx, containername)
	if err != nil {
		return err
	}
	fmt.Printf("%q\n",resp2.Kvs)
	fmt.Printf("%T\n", cancel)
	return nil
}
