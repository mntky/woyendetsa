package app

import (
	//"os"
	"fmt"
	"time"
	"context"

	"github.com/coreos/etcd/clientv3"
)

type EtcdElement struct {
	Cli			*clientv3.Client
	Kv			clientv3.KV
	Ctx			context.Context
	Cancel	context.CancelFunc
}

var (
	endpoints = []string{"localhost:2379", "localhost:22379", "localhost:32379"}
	timeout = 3 * time.Second
)

func newEtcdClient() (EtcdElement, error) {
	var etcd = EtcdElement{}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:		endpoints,
		DialTimeout:	timeout,
	})
	if err != nil {
		return etcd, err
	}
	kv := clientv3.NewKV(cli)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	//fmt.Printf("%T\n",ctx)
	//fmt.Printf("%T\n",cancel)

	etcd = EtcdElement{
		Cli:	cli,
		Kv:		kv,
		Ctx:	ctx,
		Cancel:	cancel,
	}
	return etcd, nil
}


func PutContainerSpec(specname, spec string) error {
	etcd, err := newEtcdClient()
	if err != nil {
		return err
	}
	defer etcd.Cli.Close()
	putresp, err := etcd.Kv.Put(etcd.Ctx, "/spec/"+specname, spec)
	etcd.Cancel()
	if err != nil {
		return err
	}

	fmt.Println(putresp.Header.Revision)
	return nil
}


//get container spec
func ReferContainerSpec(containername string) (string, error) {
	etcd, err := newEtcdClient()
	if err != nil {
		return "", err
	}
	defer etcd.Cli.Close()
	getresp, err := etcd.Kv.Get(etcd.Ctx, "/container/"+containername+"/spec")
	if err != nil {
		return "", err
	}

	fmt.Println(string(getresp.Kvs[0].Value))
	return string(getresp.Kvs[0].Value), nil
}


//delete container spec
func DeleteContainerSpec(containername string) error {
	etcd, err := newEtcdClient()
	if err != nil {
		return err
	}
	defer etcd.Cli.Close()

	etcd.Kv.Delete(etcd.Ctx, "/container/"+containername+"/spec", clientv3.WithPrefix())
	return nil
}

//delete all keys in etcd
func deleteAllKeys() error {
	etcd, err := newEtcdClient()
	if err != nil {
		return err
	}
	defer etcd.Cli.Close()

	etcd.Kv.Delete(etcd.Ctx, "", clientv3.WithPrefix())
	return nil
}
