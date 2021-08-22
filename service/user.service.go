package service

import (
	"bytes"
	"encoding/json"

	"github.com/fdeddys/tes/database"
	"github.com/fdeddys/tes/dto"
	"github.com/fdeddys/tes/model"
	"github.com/siddontang/go/log"

	"net/http"
)

var AWS_URL string

func init() {
	AWS_URL = "localhost:1234/api/listresources"
}

// UserService ...
type UserService struct {
}

//SaveUser ..
func (u UserService) SaveUser(reqUser *dto.RequestUser) error {

	var user model.User
	user.ApiToken = reqUser.ApiToken
	user.AwsCredential = reqUser.AwsCredential
	user.Username = reqUser.Username

	return database.SaveUser(&user)

}

//SaveUser ..
func (u UserService) ListResourceByUsername(username, token string) (int, string, []model.Resource) {

	code, desc, _ := validateInput(username, token)
	if code != 200 {
		return code, desc, nil
	}
	resources, err := database.FindResourceByUsername(username)
	if err != nil {
		return 500, err.Error(), nil
	}
	return 200, "OK", resources
}

//SaveUser ..
func (u UserService) GatherService(username, token string) (int, string) {

	code, desc, curUser := validateInput(username, token)
	if code != 200 {
		return code, desc
	}

	log.Info("Start Gathering data ...")
	tokenAWS := curUser.AwsCredential
	resCode := 200
	resDesc := "OK"
	for {
		awsResponse, errCode, errDesc := sendHttp(tokenAWS)
		if errCode != 200 {
			resCode = errCode
			resDesc = errDesc
			break
		}
		// save to DATABASE
		var resource model.Resource
		resource.Username = username
		resource.Arn = awsResponse.Resources.Arn
		resource.CreationTime = awsResponse.Resources.CreationTime
		resource.LastUpdatedTime = awsResponse.Resources.LastUpdatedTime
		resource.ResourceGroupArn = awsResponse.Resources.ResourceGroupArn
		resource.ResourceShareArn = awsResponse.Resources.ResourceShareArn
		resource.Status = awsResponse.Resources.Status
		resource.StatusMessage = awsResponse.Resources.StatusMessage
		resource.Type = awsResponse.Resources.Type

		statusCode, statusDesc := database.SaveResource(&resource)
		if statusCode != 200 {
			resCode = statusCode
			resDesc = statusDesc
			break
		}

		if awsResponse.NextToken == "" {
			break
		}
		tokenAWS = awsResponse.NextToken
	}

	return resCode, resDesc

}

func sendHttp(awsCredential string) (dto.AwsBodyResponse, int, string) {

	// gahtering process

	var awsBodyRequest dto.AwsBodyRequest
	var awsResponse dto.AwsBodyResponse

	awsBodyRequest.MaxResults = 10
	awsBodyRequest.NextToken = awsCredential
	awsBodyRequest.Principal = ""
	awsBodyRequest.ResourceArns = []string{}
	awsBodyRequest.ResourceOwner = ""
	awsBodyRequest.ResourceShareArns = []string{}
	awsBodyRequest.ResourceType = ""

	requestByte, _ := json.Marshal(awsBodyRequest)
	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest("POST", AWS_URL, requestReader)
	if err != nil {
		return awsResponse, 500, "cannot create request !"
	}

	var client = &http.Client{}
	resp, errPost := client.Do(req)

	if errPost != nil {
		return awsResponse, 500, "Cannot post to Server AWS"
	}
	defer resp.Body.Close()

	errDecoder := json.NewDecoder(resp.Body).Decode(&awsResponse)
	if errDecoder != nil {
		return awsResponse, 500, "cannot decode data from Server AWS"
	}

	return awsResponse, 200, "OK"
}

func validateInput(username, token string) (int, string, model.User) {

	curUser, err := database.FindUserByUsername(username)
	if err != nil {
		return 500, err.Error(), curUser
	}
	if curUser.ApiToken != token {
		return 401, "Invalid token", curUser
	}

	return 200, "OK", curUser
}
