package agent

import (
	"context"
	"io"
	"log"

	"github.com/yolo-sh/agent/proto"
)

type InitEnvStream interface {
	Recv() (*proto.InitEnvReply, error)
}

func (c Client) InitEnv(
	initEnvRequest *proto.InitEnvRequest,
	streamHandler func(stream InitEnvStream) error,
) error {

	return c.Execute(func(agentGRPCClient proto.AgentClient) error {
		initEnvStream, err := agentGRPCClient.InitEnv(
			context.TODO(),
			initEnvRequest,
		)

		if err != nil {
			return err
		}

		return streamHandler(initEnvStream)
	})
}

func InitEnvDefaultStreamHandler(stream InitEnvStream) error {
	for {
		initEnvReply, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		log.Println(initEnvReply.LogLine)
	}

	return nil
}
