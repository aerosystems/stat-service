package RPCServices

import (
	RPCClient "github.com/aerosystems/stat-service/pkg/rpc_client"
	"github.com/google/uuid"
)

type ProjectService interface {
	GetProjectList(userUuid uuid.UUID) ([]ProjectRPCPayload, error)
	GetProject(projectToken string) (*ProjectRPCPayload, error)
}

type ProjectRPC struct {
	rpcClient *RPCClient.ReconnectRPCClient
}

func NewProjectRPC(rpcClient *RPCClient.ReconnectRPCClient) *ProjectRPC {
	return &ProjectRPC{
		rpcClient: rpcClient,
	}
}

type ProjectRPCPayload struct {
	Id       int
	UserUuid uuid.UUID
	Name     string
	Token    string
}

func (p *ProjectRPC) GetProjectList(userUuid uuid.UUID) ([]ProjectRPCPayload, error) {
	var result []ProjectRPCPayload
	if err := p.rpcClient.Call(
		"ProjectServer.GetProjectList",
		userUuid,
		&result,
	); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProjectRPC) GetProject(projectToken string) (*ProjectRPCPayload, error) {
	var result ProjectRPCPayload
	if err := p.rpcClient.Call(
		"ProjectServer.GetProject",
		projectToken,
		&result,
	); err != nil {
		return nil, err
	}
	return &result, nil
}
