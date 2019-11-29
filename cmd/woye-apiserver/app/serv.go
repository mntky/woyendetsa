package app

import (
	"fmt"
	"net/http"
	"github.com/lxc/lxd/shared/api"
)


func startServer(url string) {
	http.HandleFunc("/api/lxc", lxc)

	http.ListenAndServe(url, nil)
}

func lxc(w http.ResponseWriter, r *http.Request) {
}
