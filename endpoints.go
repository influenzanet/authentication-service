package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	influenzanet "github.com/Influenzanet/api/dist/go"
	auth_api "github.com/Influenzanet/api/dist/go/auth-service"
	user_api "github.com/Influenzanet/api/dist/go/user-management"
)

func checkTokenAgeMaturity(issuedAt int64) bool {
	return time.Now().Unix() < time.Unix(issuedAt, 0).Add(minTokenAge).Unix()
}

func (s *authServiceServer) Status(ctx context.Context, _ *empty.Empty) (*influenzanet.Status, error) {
	return nil, errors.New("not implemented")
}

func (s *authServiceServer) LoginWithEmail(ctx context.Context, req *influenzanet.UserCredentials) (*auth_api.EncodedToken, error) {
	if req == nil {
		return nil, errors.New("invalid username and/or password")
	}
	resp, err := userManagementClient.LoginWithEmail(context.Background(), req)
	if err != nil {
		log.Printf("error during login with email: %s", err.Error())
		return nil, errors.New("invalid username and/or password")
	}

	token, err := generateNewToken(resp.UserId, resp.Roles, resp.InstanceId)
	if err != nil {
		return nil, err
	}

	return &auth_api.EncodedToken{
		Token: token,
	}, nil
}

func (s *authServiceServer) SignupWithEmail(ctx context.Context, req *user_api.NewUser) (*auth_api.EncodedToken, error) {
	if req == nil || req.Auth == nil || req.Profile == nil {
		return nil, errors.New("missing arguments")
	}
	resp, err := userManagementClient.SignupWithEmail(context.Background(), req)
	if err != nil {
		log.Printf("error during signup with email: %s", err.Error())
		return nil, err
	}

	token, err := generateNewToken(resp.UserId, resp.Roles, resp.InstanceId)
	if err != nil {
		return nil, err
	}

	return &auth_api.EncodedToken{
		Token: token,
	}, nil
}

func (s *authServiceServer) ValidateJWT(ctx context.Context, req *auth_api.EncodedToken) (*influenzanet.ParsedToken, error) {
	return nil, errors.New("not implemented")
}

/*
func ValidateJWT(context *gin.Context) {
	token := context.MustGet("encodedToken").(string)

	// Parse and validate token
	parsedToken, ok, oldKey, err := validateToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid", "reason": err.Error()})
		return
	}
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid", "reason": "wrong signiture"})
		return
	}

	// TODO: what to do when token is generated by old key
	if oldKey {
		log.Println(oldKey)
	}

	context.JSON(http.StatusOK, gin.H{"token": parsedToken})
}
*/

func (s *authServiceServer) RenewJWT(ctx context.Context, req *auth_api.EncodedToken) (*auth_api.EncodedToken, error) {
	return nil, errors.New("not implemented")
}

/*
func renewTokenHandl(context *gin.Context) {
	token := context.MustGet("encodedToken").(string)

	// Parse and validate token
	parsedToken, ok, _, err := validateToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid", "reason": err.Error()})
		return
	}
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid", "reason": "wrong signiture"})
		return
	}

	// Check for too frequent requests:
	if checkTokenAgeMaturity(parsedToken.StandardClaims.IssuedAt) {
		context.JSON(http.StatusTeapot, gin.H{"error": "can't renew token so often"})
		return
	}

	// Generate new token:
	newToken, err := generateNewToken(parsedToken.UserID, parsedToken.Roles, parsedToken.AuthenticatedRole)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send response
	tokenResp := tokenMessage{
		Token: newToken,
	}
	context.JSON(http.StatusOK, tokenResp)
}
*/
