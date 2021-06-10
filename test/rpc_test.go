package test

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestClient(t *testing.T) {
	var args = Args{A: 32, B: 14}
	var result = Result{}
	logrus.Info("test")

	var client, err = rpc.DialHTTP("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("connect rpc server failed, err:%v", err)
	}

	err = client.Call("MathService.Divide", args, &result)
	if err != nil {
		fmt.Printf("call math service failed, err:%v", err)
	}
	fmt.Printf("call RPC server success, result:%f", result.Value)
}

func TestServer(t *testing.T) {
	var ms = new(MathService)
	// 注册 RPC 服务
	err := rpc.Register(ms)
	if err != nil {
		fmt.Printf("rpc server register faild, err:%s", err)
	}
	// 将 RPC 服务绑定到 HTTP 服务中去
	rpc.HandleHTTP()

	fmt.Printf("server start ....")
	err = http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Printf("listen and server is failed, err:%v\n", err)
	}

	fmt.Printf("server stop ....")
}

type Args struct {
	A, B float32
}

type Result struct {
	Value float32
}

type MathService struct{}

func (s *MathService) Add(args *Args, result *Result) error {
	result.Value = args.A + args.B
	logrus.Info(result)
	return nil
}

func (s *MathService) Divide(args *Args, result *Result) error {
	if args.B == 0 {
		return errors.New("arge.B is 0")
	}

	result.Value = args.A / args.B
	logrus.Info(result)
	return nil
}
