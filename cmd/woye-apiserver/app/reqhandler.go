package app

import (
	"fmt"
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


func startServer(url string) {
	http.HandleFunc("/api/lxc/create", lxc_create)
	http.HandleFunc("/api/lxc/get", lxc_get)
	http.HandleFunc("/api/lxc/update", lxc_update)
	http.HandleFunc("/api/lxc/delete", lxc_delete)

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

	fmt.Println("--Print SPEC--")
	fmt.Println(spec)
	fmt.Println(spec.Spec)
	fmt.Println(spec.Spec.Container)
	fmt.Println(spec.Spec.Container.Name)
	fmt.Println(spec.Spec.Container.Distri)

	fmt.Println("--etcd put--")
	err = PutContainerSpec(spec.Name, spec.Spec)
	if err != nil {
		fmt.Println(err)
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
