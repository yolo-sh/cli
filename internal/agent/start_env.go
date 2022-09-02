package agent

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/yolo-sh/agent/proto"
	"github.com/yolo-sh/cli/internal/constants"
)

type BuildAndStartEnvStream interface {
	Recv() (*proto.BuildAndStartEnvReply, error)
}

func (c Client) BuildAndStartEnv(
	startEnvRequest *proto.BuildAndStartEnvRequest,
	streamHandler func(stream BuildAndStartEnvStream) error,
) error {

	return c.Execute(func(agentGRPCClient proto.AgentClient) error {
		startEnvStream, err := agentGRPCClient.BuildAndStartEnv(
			context.TODO(),
			startEnvRequest,
		)

		if err != nil {
			return err
		}

		return streamHandler(startEnvStream)
	})
}

func BuildAndStartEnvDefaultStreamHandler(stream BuildAndStartEnvStream) error {
	for {
		startEnvReply, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if len(startEnvReply.LogLineHeader) > 0 {
			bold := constants.Bold
			fmt.Println(bold("[" + startEnvReply.LogLineHeader + "]\n"))
		}

		if len(startEnvReply.LogLine) > 0 {
			log.Println(startEnvReply.LogLine)
		}
	}

	return nil
}
