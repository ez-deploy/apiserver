package identity

import (
	"encoding/base64"
	"encoding/json"

	"github.com/ez-deploy/apiserver/models"
)

func StringifyToken(token *models.ModelToken) (string, error) {
	rawToken, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(rawToken), nil
}

func NewTokenFromString(base64RawToken string) (*models.ModelToken, error) {
	rawToken, err := base64.StdEncoding.DecodeString(base64RawToken)
	if err != nil {
		return nil, err
	}

	token := &models.ModelToken{}
	if err := json.Unmarshal(rawToken, token); err != nil {
		return nil, err
	}

	if token.Type == nil {
		token.Type = models.NewModelTokenType(models.ModelTokenTypeSession)
	}

	return token, nil
}
