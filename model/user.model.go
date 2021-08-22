package model

type User struct {
	ID            int64  `json:"id" gorm:"column:id"`
	Username      string `json:"username" gorm:"column:user_name"`
	ApiToken      string `json:"apiToken" gorm:"column:api_token"`
	AwsCredential string `json:"awsCredential" gorm:"column:aws_credential"`
}

func (u *User) TableName() string {
	return "public.users"
}
