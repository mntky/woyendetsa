package action

import (
)

var saddr = X.X.X.X
var sport = 8080

func send(args) (string, error) {
	fmt.Printf("send :&s", args)
	ln, err = net.Dial("tcp", saddr+sport)
	if err != nil {
		return "", err
	}

	fmt.Fprintf(ln, args)
	ln.Close()

	return "ok", nil
}
