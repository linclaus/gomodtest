package test

import (
	"errors"
	"net/http"
	"net/rpc"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRPC(t *testing.T) {
	go serverTest()
	clientTest()
}

func clientTest() {
	var args = Args{A: 32, B: 14}
	var result = Result{}
	logrus.Info("test")

	var client, err = rpc.DialHTTP("tcp", "127.0.0.1:9090")
	if err != nil {
		logrus.Infof("connect rpc server failed, err:%v", err)
	}

	err = client.Call("MathService.Divide", args, &result)
	if err != nil {
		logrus.Infof("call math service failed, err:%v", err)
	}
	logrus.Infof("call RPC server success, result:%f", result.Value)
}

func serverTest() {
	var ms = new(MathService)
	// 注册 RPC 服务
	err := rpc.Register(ms)
	if err != nil {
		logrus.Infof("rpc server register faild, err:%s", err)
	}
	// 将 RPC 服务绑定到 HTTP 服务中去
	rpc.HandleHTTP()

	logrus.Infof("server start ....")
	err = http.ListenAndServe(":9090", nil)

	if err != nil {
		logrus.Infof("listen and server is failed, err:%v\n", err)
	}

	logrus.Infof("server stop ....")
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
