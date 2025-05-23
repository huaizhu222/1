package hanlder

import (
	"io"
	"os"
	"user_srv/proto"
)

type User_File struct {
	proto.UnimplementedSystemServer
}

func (s *User_File) SendFile(req *proto.SendFileRequest, stream proto.System_SendFileServer) error {
	file, _ := os.Open(req.FilePath)
	defer file.Close()

	buffer := make([]byte, 1024*1024) // 1MB 分块
	chunkNumber := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			return nil
		}
		// 发送分块
		stream.Send(&proto.SendFileResponse{
			ChunkData:   buffer[:n],
			ChunkNumber: int32(chunkNumber),
		})
		chunkNumber++
	}
}
