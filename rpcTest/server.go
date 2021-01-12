package main

import (
	"net/http"
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

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	log.Infof("Rect Area: %d", *ret)
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	log.Infof("Rect Perimeter: %d", *ret)
	return nil
}

func main() {
	log.Println("Start Server")
	rect := new(Rect)
	rpc.Register(rect)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
}
