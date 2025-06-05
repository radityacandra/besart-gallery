package jwt

import (
	"context"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/radityacandra/besart-gallery/pkg/jwt/types"
)

const (
	ISSUER      = "http://localhost:8080/realms/myrealm"
	AUDIENCE    = "account"
	CONTEXT_KEY = "token_data"
)

func AuthorizeToken(ctx context.Context, authorizationStr string) (map[string]interface{}, error) {
	authPart := strings.Split(authorizationStr, " ")
	if len(authPart) != 2 && authPart[0] != "Bearer" {
		return nil, types.NewAuthorizationError("invalid authorization header")
	}
	tokenStr := authPart[1]

	url := "http://localhost:8080/realms/myrealm/protocol/openid-connect/certs"
	jwks, err := keyfunc.NewDefaultCtx(ctx, []string{url})
	if err != nil {
		return nil, types.NewAuthorizationError("failed to create JWK Set")
	}

	token, err := jwt.Parse(tokenStr, jwks.Keyfunc)
	if err != nil {
		return nil, types.NewAuthorizationError("failed to parse the JWT")
	}

	if !token.Valid {
		return nil, types.NewAuthorizationError("token is not valid")
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, types.NewAuthorizationError("invalid token")
	}

	claim := token.Claims.(jwt.MapClaims)

	// validate expired claim
	if exp, err := claim.GetExpirationTime(); err != nil || exp.Unix() < time.Now().Unix() {
		return nil, types.NewAuthorizationError("token has been expired")
	}

	// validate issuer
	if issuer, err := claim.GetIssuer(); err != nil || issuer != ISSUER {
		return nil, types.NewAuthorizationError("invalid token")
	}

	// validate audience
	if audience, err := claim.GetAudience(); err != nil || audience[0] != AUDIENCE {
		return nil, types.NewAuthorizationError("invalid token")
	}

	mapData := make(map[string]interface{})
	for key, value := range claim {
		mapData[key] = value
	}

	return mapData, nil
}
