package app

import (
	"fmt"
	"bytes"
	//"io/ioutil"
	"encoding/json"
	"net/http"
	//"github.com/lxc/lxd/shared/api"
)

type createdata struct {
	Name		string
	Replica int
	Test		string
}

func startServer(url string) {
	http.HandleFunc("/api/lxc/create", lxc_create)
	http.HandleFunc("/api/lxc/read", lxc_read)
	http.HandleFunc("/api/lxc/update", lxc_update)
	http.HandleFunc("/api/lxc/delete", lxc_delete)

	http.ListenAndServe(url, nil)
}

func lxc_create(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method %s \n",r.Method)

	fmt.Println("--serv bufbody---")
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()
	fmt.Println(body)

	var rdata createdata
	err := json.Unmarshal(bufbody.Bytes(), &rdata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rdata)

	cli, err := NewEtcdClient()
	if err != nil {
		fmt.Println(err)
	}
	err = Putkvs(cli, rdata.Name, body)
	if err != nil {
		fmt.Println(err)
	}


	w.Write([]byte("OK"))
}

func lxc_read(w http.ResponseWriter, r *http.Request) {
}

func lxc_update(w http.ResponseWriter, r *http.Request) {
}

func lxc_delete(w http.ResponseWriter, r *http.Request) {
}
