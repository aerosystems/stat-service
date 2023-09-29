package helpers

import RPCClient "github.com/aerosystems/stat-service/internal/rpc_client"

func ContainsProjectToken(projectList []RPCClient.ProjectRPCPayload, projectToken string) bool {
	for _, project := range projectList {
		if project.Token == projectToken {
			return true
		}
	}
	return false
}
