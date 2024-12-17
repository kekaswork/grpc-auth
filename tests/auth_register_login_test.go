package tests

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt"
	"github.com/kekaswork/grpc-auth/tests/suite"
	authv1 "github.com/kekaswork/protos/gen/go/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	emptyAppID = 0
	appID      = 1
	appSecret  = "secret-sample"

	passDefaultLen = 10
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	password := randomFakePassowrd()

	respReg, err := st.AuthClient.Register(ctx, &authv1.RegisterRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	loginRequest := authv1.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    appID,
	}
	respLogin, err := st.AuthClient.Login(ctx, &loginRequest)
	require.NoError(t, err)

	token := respLogin.GetToken()
	require.NotEmpty(t, token)

	loginTime := time.Now()

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	require.NoError(t, err)

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	assert.True(t, ok)

	assert.Equal(t, respReg.GetUserId(), int64(claims["uid"].(float64)))
	assert.Equal(t, email, claims["email"].(string))
	assert.Equal(t, appID, int(claims["app_id"].(float64)))
	const deltaSeconds = 3
	assert.InDelta(t, loginTime.Add(st.Cfg.TokenTTL).Unix(), claims["exp"].(float64), deltaSeconds)
}

func randomFakePassowrd() string {
	return gofakeit.Password(
		true,
		true,
		true,
		true,
		false,
		passDefaultLen,
	)
}
