package app

import (
	"fmt"
	//"net"
	"bytes"
	//"io/ioutil"
	yaml "gopkg.in/yaml.v2"
	"encoding/json"
	"net/http"
	//"github.com/lxc/lxd/shared/api"
)

var (
	resptxt = "ok"
	spec specstruct

	//test
	ddata specstruct
)
type specstruct struct {
	Name	string `yaml:"name"`
	Spec	Specmeta `yaml:"spec"`
}
type Specmeta struct {
	Container containerstruct `yaml:"container"`
	Test	test	`yaml:"test"`
}
type containerstruct struct {
	Name		string `yaml:"name"`
	Distri	string `yaml:"distri"`
	Release	string `yaml:"release"`
	Arch		string `yaml:"arch"`
}
type test struct {
	Name	string `yaml:"name"`
}

type NodeMetadata struct {
	Name	string `yaml:"name"`
	Addr	string `yaml:"addr"`
}


func startServer(url string) {
	http.HandleFunc("/api/lxc/create", lxc_create)
	http.HandleFunc("/api/lxc/get", lxc_get)
	http.HandleFunc("/api/lxc/update", lxc_update)
	http.HandleFunc("/api/lxc/delete", lxc_delete)
	http.HandleFunc("/api/node/add", node_add)
	http.HandleFunc("/api/refer/spec", refer_spec)
//	http.HandleFunc("/api/node/delete", node_delete)

	err := http.ListenAndServe(url, nil)
	fmt.Println(err)
}


//woctl create command.
//register containername:spec to etcd
func lxc_create(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	fmt.Println("--create--")
	//送られてきたjsonのspecを構造体にはめる
	err := yaml.Unmarshal(bufbody.Bytes(), &spec)
	if err != nil {
		fmt.Println(err)
		resptxt = err.Error()
	}

	fmt.Println("--etcd put--")
	err = PutContainerSpec(spec.Name, spec.Spec)
	if err != nil {
		fmt.Println(err)
	}else {
		notification("/spec/"+spec.Name)
	}

	w.Write([]byte(resptxt))
}


//refer container spec
func lxc_get(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	err := json.Unmarshal(bufbody.Bytes(), &spec)
	if err != nil {
		fmt.Println(err)
		resptxt = err.Error()
	}

	resptxt, err := ReferContainerSpec(spec.Name)
	if err != nil {
		fmt.Println(err)
		resptxt = err.Error()
	}

	w.Write([]byte(resptxt))
}

func lxc_update(w http.ResponseWriter, r *http.Request) {
	lxc_create(w, r)
}

func lxc_delete(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)

	err := yaml.Unmarshal(bufbody.Bytes(), &spec)
	if err != nil {
		fmt.Println(err)
		resptxt = err.Error()
	}

	err = DeleteContainerSpec(spec.Name)
	if err != nil {
		fmt.Println(err)
		resptxt = err.Error()
	}
	w.Write([]byte(resptxt))
}

func node_add(w http.ResponseWriter, r *http.Request) {
	nodemeta := &NodeMetadata{}
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	err := json.Unmarshal(bufbody.Bytes(), &nodemeta)
	if err != nil {
		fmt.Println(err)
	}
	err = PutNodeMeta(*nodemeta)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	fmt.Println("request data")
	fmt.Println(string(nodemeta.Name))
	fmt.Println(string(nodemeta.Addr))
	w.Write([]byte("ok"))
}

func refer_spec(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api/refer debug")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	fmt.Println(buf.String())

	spec, err := ReferContainerSpec(buf.String())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(spec)
	w.Write([]byte(spec))
}

//func node_delete(w http.ResponseWriter, r *http.Request) {
//	nodemeta := &nodemetadata{}
//	srcip := r.Header.Get("X-Real-IP")
//	fmt.Println(srcip)
//	nodemeta.Name = "test"
//
//	w.Write([]byte("ok"))
//}
