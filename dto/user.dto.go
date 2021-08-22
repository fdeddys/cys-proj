package dto

type RequestUser struct {
	Username      string `json:"username"`
	ApiToken      string `json:"apiToken"`
	AwsCredential string `json:"awsCredential"`
}

type RequestGather struct {
	Username string `json:"username"`
}
