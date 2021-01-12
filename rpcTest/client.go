package main

import (
	"net/rpc"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
}

type Params struct {
	Width, Height int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	ret := 0
	err = conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Rect Area is: %d", ret)

	err = conn.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Rect Perimeter is: %d", ret)
}
