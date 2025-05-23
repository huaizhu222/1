package hanlder

import (
	"io"
	"os"

	"user_srv/proto"
)

type Server struct {
	proto.UnimplementedSystemServer
}

func (s *Server) SendFile(req *proto.FileRequest, stream proto.System_SendFileServer) error {
	file, err := os.Open(req.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stream.Send(&proto.FileResponse{Chunk: buf[:n]})
	}
	return nil
}
