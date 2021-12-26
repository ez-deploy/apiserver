package handle

import (
	"log"

	"github.com/ez-deploy/apiserver/handle/internal/authority"
	"github.com/ez-deploy/apiserver/handle/internal/configurable"
	"github.com/ez-deploy/apiserver/handle/internal/identity"
	"github.com/ez-deploy/apiserver/handle/internal/project"
)

const configFilename = "apiserver.conf.yaml"

type handlerImpl struct {
	*configurable.ConfigurableImpl
	*authority.AuthorityOpsImpl
	*identity.AuthableImpl
	*identity.IdentityOpsImpl
	*project.ProjectOpsImpl
}

func New() *handlerImpl {
	cfg, err := getConfigFromFile(configFilename)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := createGRPCClientsFromConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &handlerImpl{
		ConfigurableImpl: &configurable.ConfigurableImpl{},
		AuthableImpl:     identity.NewAuthableImpl(clientset.identityClient),
		IdentityOpsImpl:  identity.NewIdentityOpsImpl(clientset.identityClient),
		AuthorityOpsImpl: &authority.AuthorityOpsImpl{},
		ProjectOpsImpl:   &project.ProjectOpsImpl{},
	}
}
