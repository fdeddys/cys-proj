package model

type Resource struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:user_name"`

	// The Amazon Resource Name
	Arn string `json:"arn" gorm:"column:arn"`
	// The time when the resource was associated with the resource share.
	CreationTime int64 `json:"creationTime" gorm:"column:creation_time"`
	// The time when the association was last updated.
	LastUpdatedTime int64 `json:"lastUpdatedTime" gorm:"column:last_updated_time"`
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceGroupArn string `json:"resourceGroupArn" gorm:"column:resource_group_arn"`
	// ResourceShareArn
	ResourceShareArn string `json:"resourceShareArn" gorm:"column:resource_share_arn"`
	//The status of the resource.
	// Valid Values: AVAILABLE | ZONAL_RESOURCE_INACCESSIBLE | LIMIT_EXCEEDED | UNAVAILABLE | PENDING
	Status string `json:"status" gorm:"column:status"`
	// A message about the status of the resource.
	StatusMessage string `json:"statusMessage" gorm:"column:status_message"`
	// The resource type.
	Type string `json:"type" gorm:"column:type"`
}

func (r *Resource) TableName() string {
	return "public.resources"
}
