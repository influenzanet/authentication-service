package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type userCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserModel holds information relevant for authentication
type UserModel struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	ID       uint     `json:"user_id"`
	Roles    []string `json:"roles"` // PARTICIPANT RESEARCHER ADMIN
}

type errorResponse struct {
	Error string `json:"error"`
}

type tokenMessage struct {
	Token string `json:"token"`
}

func checkTokenAgeMaturity(issuedAt int64) bool {
	return time.Now().Unix() < time.Unix(issuedAt, 0).Add(minTokenAge).Unix()
}

func loginParticipantHandl(context *gin.Context) {
	var requiredAuthLevel = "PARTICIPANT"
	var creds userCredentials
	if err := context.ShouldBindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginPayload, err := json.Marshal(creds)

	// reach out to user-management service to check credentials
	resp, err := http.Post(userManagementServer+"/login", "application/json", bytes.NewBuffer(loginPayload))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		currentError := errorResponse{}
		jsonErr := json.Unmarshal(respBody, &currentError)
		if jsonErr != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(resp.StatusCode, currentError)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	currentUser := UserModel{}
	jsonErr := json.Unmarshal(respBody, &currentUser)
	if jsonErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// generate token
	if len(currentUser.Roles) < 1 || currentUser.Roles[0] != requiredAuthLevel { // TODO implement contains method to check for the required role instead of hardcoding
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "required account authorization level not met"})
		return
	}
	token, err := generateNewToken(currentUser.ID, []string{"PARTICIPANT"})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send response
	tokenResp := tokenMessage{
		Token: token,
	}
	context.JSON(http.StatusOK, tokenResp)
}

func loginResearcherHandl(context *gin.Context) {
	var requiredAuthLevel = "RESEARCHER"
	var creds userCredentials
	if err := context.ShouldBindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginPayload, err := json.Marshal(creds)

	// reach out to user-management service to check credentials
	resp, err := http.Post(userManagementServer+"/login", "application/json", bytes.NewBuffer(loginPayload))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		currentError := errorResponse{}
		jsonErr := json.Unmarshal(respBody, &currentError)
		if jsonErr != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(resp.StatusCode, currentError)
		return

	}

	respBody, err := ioutil.ReadAll(resp.Body)
	currentUser := UserModel{}
	jsonErr := json.Unmarshal(respBody, &currentUser)
	if jsonErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// generate token
	if len(currentUser.Roles) < 2 || currentUser.Roles[1] != requiredAuthLevel { // TODO implement contains method to check for the required role instead of hardcoding
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "required account authorization level not met"})
		return
	}
	token, err := generateNewToken(currentUser.ID, []string{"PARTICIPANT", "RESEARCHER"})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send response
	tokenResp := tokenMessage{
		Token: token,
	}
	context.JSON(http.StatusOK, tokenResp)
}

func loginAdminHandl(context *gin.Context) {
	var requiredAuthLevel = "ADMIN"
	var creds userCredentials
	if err := context.ShouldBindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginPayload, err := json.Marshal(creds)

	// reach out to user-management service to check credentials
	resp, err := http.Post(userManagementServer+"/login", "application/json", bytes.NewBuffer(loginPayload))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		currentError := errorResponse{}
		jsonErr := json.Unmarshal(respBody, &currentError)
		if jsonErr != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(resp.StatusCode, currentError)
		return

	}

	respBody, err := ioutil.ReadAll(resp.Body)
	currentUser := UserModel{}
	jsonErr := json.Unmarshal(respBody, &currentUser)
	if jsonErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// generate tokens
	if len(currentUser.Roles) < 3 || currentUser.Roles[2] != requiredAuthLevel { // TODO implement contains method to check for the required role instead of hardcoding
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "required account authorization level not met"})
		return
	}
	token, err := generateNewToken(currentUser.ID, []string{"PARTICIPANT", "RESEARCHER", "ADMIN"})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send response
	tokenResp := tokenMessage{
		Token: token,
	}
	context.JSON(http.StatusOK, tokenResp)
}

func signupParticipantHandl(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func signupResearcherHandl(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func signupAdminHandl(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func validateTokenHandl(context *gin.Context) {
	req := context.Request

	// Get token string from url or header field
	var token string
	tokens, ok := req.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
	} else if len(req.FormValue("token")) > 0 {
		token = req.FormValue("token")
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": "no Authorization token found"})
		return
	}

	// Parse and validate token
	parsedToken, ok, oldKey, err := validateToken(token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
		return
	}
	// TODO: what to do when token is generated by old key
	log.Println(oldKey)

	context.JSON(http.StatusOK, parsedToken)
}

func renewTokenHandl(context *gin.Context) {
	req := context.Request

	// Get token string from url or header field
	var token string
	tokens, ok := req.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
		if len(token) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "no Authorization token found"})
			return
		}
	} else if len(req.FormValue("token")) > 0 {
		token = req.FormValue("token")
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": "no Authorization token found"})
		return
	}

	// Parse and validate token
	parsedToken, ok, _, err := validateToken(token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
		return
	}

	// Check for too frequent requests:
	if checkTokenAgeMaturity(parsedToken.StandardClaims.IssuedAt) {
		context.JSON(http.StatusTeapot, gin.H{"error": "can't renew token so often"})
		return
	}

	// Generate new token:
	newToken, err := generateNewToken(parsedToken.UserID, parsedToken.Roles)
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
