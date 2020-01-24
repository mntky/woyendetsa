package app

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

type Notice struct {
	Name	string
}

func notification(specname string) {
	fmt.Println("debug" + specname)
	//node取得
	keys, err := GetKey("/node")
	if err != nil {
		fmt.Println(err)
	}
	for _, d := range keys {
		fmt.Printf(d + ":")
		value, err := ReferContainerSpec(d)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = noticesend(value, specname)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func noticesend(addr, specname string) error {
	noticestruct := &Notice{Name: specname}
	noticejson, _ := json.Marshal(noticestruct)
	req, err := http.NewRequest(
		"POST",
		"http://"+addr+"/api/notice",
		bytes.NewBuffer([]byte(noticejson)),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}
