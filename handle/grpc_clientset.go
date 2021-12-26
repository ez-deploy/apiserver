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

	identityClient := identitypb.NewIdentityOpsClient(identityConn)

	clientset := &grpcClientSet{
		identityClient: identityClient,
	}

	return clientset, nil
}
