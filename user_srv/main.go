package main

import (
	"fmt"
	"log"
	"net"
	"user_srv/hanlder"
	"user_srv/initalize"
	"user_srv/proto"

	"net/http"
	_ "net/http/pprof"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//初始化
	initalize.InitLogger()
	initalize.InitConfig()
	initalize.InitDB()
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "120.55.0.93:6831",
		},
		ServiceName: "user-srv",
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))

	if err != nil {
		log.Fatalf("Failed to initialize Jaeger Tracer: %v", err)
	}
	//设置tracer
	defer closer.Close() //关闭jaeger
	opentracing.SetGlobalTracer(tracer)

	go func() {
		server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))
		proto.RegisterUserServer(server, &hanlder.User_Server{})
		proto.RegisterSystemServer(server, &hanlder.Server{})
		lis, err := net.Listen("tcp", "127.0.0.1:50051")
		if err != nil {
			zap.S().Errorw("连接失败")
			return
		}
		err = server.Serve(lis) //启动端口
		if err != nil {
			fmt.Println("错误:", err)
		}

	}()

	// 性能分析
	http.ListenAndServe(":50051", nil)

}
