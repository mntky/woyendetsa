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
type deletedata struct {
	Name	string
}

func startServer(url string) {
	http.HandleFunc("/api/lxc/create", lxc_create)
	http.HandleFunc("/api/lxc/get", lxc_get)
	http.HandleFunc("/api/lxc/update", lxc_update)
	http.HandleFunc("/api/lxc/delete", lxc_delete)

	err := http.ListenAndServe(url, nil)
	fmt.Println(err)
}

var (
	resptxt = []byte("ok")
	cdata createdata
	ddata deletedata
)

//woctl create command.
//register containername:spec to etcd
func lxc_create(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	fmt.Println("--create--")
	//送られてきたjsonのspecを構造体にはめる
	err := json.Unmarshal(bufbody.Bytes(), &cdata)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--etcd put--")
	//cdata.Nameはコンテナ名、bodyは送られてきたjson
	err = PutContainerSpec(cdata.Name, bufbody.String() )
	if err != nil {
		fmt.Println(err)
	}

	w.Write(resptxt)
}


//get container spec
func lxc_get(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	err := json.Unmarshal(bufbody.Bytes(), &ddata)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := ReferContainerSpec(ddata.Name)
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(resp))
}

func lxc_update(w http.ResponseWriter, r *http.Request) {
	lxc_create(w, r)
}

func lxc_delete(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	err := json.Unmarshal(bufbody.Bytes(), &ddata)
	if err != nil {
		fmt.Println(err)
	}

	err = DeleteContainerSpec(ddata.Name)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(resptxt)
}
