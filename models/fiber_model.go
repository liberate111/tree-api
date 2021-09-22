package models

type ResponseMessage struct {
	Status  int
	Message string
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
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

type ManageTree struct {
	Uuid       string `json:"uuid,omitempty"`
	TreeName   string `json:"treeName,omitempty"`
	Owner      string `json:"owner,omitempty"`
	Username   string `json:"username,omitempty"`
	TableName  string `json:"tableName,omitempty"`
	PrimaryKey string `json:"primaryKey,omitempty"`
}
