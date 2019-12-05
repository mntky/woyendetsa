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
	http.HandleFunc("/api/lxc/read", lxc_read)
	http.HandleFunc("/api/lxc/update", lxc_update)
	http.HandleFunc("/api/lxc/delete", lxc_delete)

	http.ListenAndServe(url, nil)
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

	//送られてきたjsonのspecを構造体にはめる
	err := json.Unmarshal(bufbody.Bytes(), &cdata)
	if err != nil {
		fmt.Println(err)
	}
	//cdata.Nameはコンテナ名、bodyは送られてきたjson
	err = PutContainerSpec(cdata.Name, bufbody.String() )
	if err != nil {
		fmt.Println(err)
	}

	w.Write(resptxt)
}


//get container spec
func lxc_read(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	err := json.Unmarshal(bufbody.Bytes(), &ddata)
	if err != nil {
		fmt.Println(err)
	}

	err = ReadContainerSpec(ddata.Name)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(resptxt)
}

func lxc_update(w http.ResponseWriter, r *http.Request) {
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
