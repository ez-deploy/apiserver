package handle

import (
	"fmt"

	authoritypb "github.com/ez-deploy/protobuf/authority"
	identitypb "github.com/ez-deploy/protobuf/identity"
	projectpb "github.com/ez-deploy/protobuf/project"
	"google.golang.org/grpc"
)

type grpcClientSet struct {
	identityClient  identitypb.IdentityOpsClient
	authorityClient authoritypb.AuthorityOpsClient
	projectClient   projectpb.ProjectOpsClient
}

func createGRPCClientsFromConfig(cfg *config) (*grpcClientSet, error) {
	identityAddr := fmt.Sprintf("%v:%v", cfg.GRPCIdentityHost, cfg.GRPCIdentityPort)
	identityConn, err := grpc.Dial(identityAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	projectAddr := fmt.Sprintf("%v:%v", cfg.GRPCProjectHost, cfg.GRPCProjectPort)
	projectConn, err := grpc.Dial(projectAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	authorityAddr := fmt.Sprintf("%v:%v", cfg.GRPCAuthorityHost, cfg.GRPCAuthorityPort)
	authorityConn, err := grpc.Dial(authorityAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	identityClient := identitypb.NewIdentityOpsClient(identityConn)
	projectClient := projectpb.NewProjectOpsClient(projectConn)
	authorityClient := authoritypb.NewAuthorityOpsClient(authorityConn)

	clientset := &grpcClientSet{
		identityClient:  identityClient,
		projectClient:   projectClient,
		authorityClient: authorityClient,
	}

	return clientset, nil
}
