package models

type ResponseMessage struct {
	Status  int
	Message string
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Uuid struct {
	Uuid string `json:"uuid"`
}

type Item struct {
	Tree []string `json:"tree"`
}

type ManageUser struct {
	Uuid       string `json:"uuid,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	TableName  string `json:"tableName,omitempty"`
	PrimaryKey string `json:"primaryKey,omitempty"`
}
