package rpc

import (
	"github.com/aerosystems/stat-service/internal/models"
	RpcClient "github.com/aerosystems/stat-service/pkg/rpc_client"
	"github.com/google/uuid"
)

type ProjectRepo struct {
	rpcClient *RpcClient.ReconnectRpcClient
}

func NewProjectRepo(rpcClient *RpcClient.ReconnectRpcClient) *ProjectRepo {
	return &ProjectRepo{
		rpcClient: rpcClient,
	}
}

func (p *ProjectRepo) GetProjectList(userUuid uuid.UUID) ([]models.Project, error) {
	var result []models.Project
	if err := p.rpcClient.Call(
		"ProjectServer.GetProjectList",
		userUuid,
		&result,
	); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *ProjectRepo) GetProject(projectToken string) (*models.Project, error) {
	var result models.Project
	if err := p.rpcClient.Call(
		"ProjectServer.GetProject",
		projectToken,
		&result,
	); err != nil {
		return nil, err
	}
	return &result, nil
}
