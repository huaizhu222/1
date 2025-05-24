package hanlder

import (
	"context"
	"io/ioutil"
	"net"
	"os"
	"testing"

	pb "user_srv/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

func TestSendFile(t *testing.T) {
	// 创建内存中的监听器
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterSystemServer(s, &Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()

	// 创建测试文件
	tmpContent := []byte("test file content")
	tmpFile, err := os.CreateTemp("E:/", "test.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(tmpContent); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	// 创建客户端连接
	conn, _ := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewSystemClient(conn)
	stream, _ := client.SendFile(context.Background(), &pb.FileRequest{FilePath: tmpFile.Name()})

	// 接收数据
	var received []byte
	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}
		received = append(received, res.Chunk...)
	}

	// 验证内容
	original, _ := ioutil.ReadFile(tmpFile.Name())
	if string(received) != string(original) {
		t.Errorf("Expected %q, got %q", original, received)
	}
}
