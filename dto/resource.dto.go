package dto

import "github.com/fdeddys/tes/model"

type RequestResource struct {
	MaxResults int16  `json:"maxResults"`
	NextToken  string `json:"nextToken"`
}

type AwsBodyRequest struct {
	MaxResults        int64    `json:"MaxResults"`
	NextToken         string   `json:"NextToken"`
	Principal         string   `json:"Principal"`
	ResourceArns      []string `json:"ResourceArns"`
	ResourceOwner     string   `json:"ResourceOwner"`
	ResourceShareArns []string `json:"ResourceShareArns"`
	ResourceType      string   `json:"ResourceType"`
}

type AwsBodyResponse struct {
	NextToken string         `json:"nextToken"`
	Resources model.Resource `json:"resources"`
}
