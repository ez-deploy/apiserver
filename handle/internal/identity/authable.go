package identity

import (
	"context"
	"fmt"

	"github.com/ez-deploy/apiserver/models"
	pb "github.com/ez-deploy/protobuf/identity"
	"github.com/ez-deploy/protobuf/model"
)

func NewAuthableImpl(identityClient pb.IdentityOpsClient) *AuthableImpl {
	return &AuthableImpl{cli: identityClient}
}

type AuthableImpl struct {
	cli pb.IdentityOpsClient
}

func (h *AuthableImpl) KeyAuth(token string) (*models.IdentityVerifyResp, error) {
	fullToken, err := NewTokenFromString(token)
	if err != nil {
		return nil, err
	}

	fmt.Println("=========================\n full Token: ", fullToken)

	pbreq := &pb.VerifyReq{
		Token: &model.Token{
			Type:  model.TokenType(model.TokenType_value[string(*fullToken.Type)]),
			Token: fullToken.Token,
		},
	}

	pbresp, err := h.cli.Verify(context.Background(), pbreq)
	if err != nil {
		return nil, err
	}

	resp := &models.IdentityVerifyResp{
		Identity: &models.ModelIdentity{
			Email: pbresp.Identity.Email,
			Name:  pbreq.Token.Token,
		},
		TokenType: models.NewModelTokenType(models.ModelTokenType(pbresp.TokenType.String())),
	}

	return resp, nil
}
