package agent

import (
	"net"

	"github.com/yolo-sh/agent/constants"
)

type ClientBuilder interface {
	Build(config ClientConfig) ClientInterface
}

type DefaultClientBuilder struct{}

func NewDefaultClientBuilder() DefaultClientBuilder {
	return DefaultClientBuilder{}
}

func (DefaultClientBuilder) Build(config ClientConfig) ClientInterface {
	return NewClient(config)
}

func NewDefaultClientConfig(
	sshPrivateKeyBytes []byte,
	instancePublicIPAddress string,
) ClientConfig {

	return ClientConfig{
		ServerRootUser:           constants.YoloUserName,
		ServerSSHPrivateKeyBytes: sshPrivateKeyBytes,
		ServerAddr: net.JoinHostPort(
			instancePublicIPAddress,
			constants.SSHServerListenPort,
		),
		// Ends with ":" to let "net.listener"
		// choose a random available port for us
		LocalAddr:          "127.0.0.1:",
		RemoteAddrProtocol: constants.GRPCServerAddrProtocol,
		RemoteAddr:         constants.GRPCServerAddr,
	}
}
