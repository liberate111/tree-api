package models

type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    *Data  `json:"data,omitempty"`
}

type Data struct {
	Uuid      string   `json:"uuid,omitempty"`
	Tree      []string `json:"tree,omitempty"`
	Users     []User   `json:"users,omitempty"`
	Trees     []Tree   `json:"trees,omitempty"`
	TreeName  string   `json:"treeName,omitempty"`
	Owner     string   `json:"owner,omitempty"`
	Level     *int     `json:"level,omitempty"`
	State     string   `json:"state,omitempty"`
	StartTime *int     `json:"startTime,omitempty"`
	StopTime  *int     `json:"stopTime,omitempty"`
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
	Level      int    `json:"level,omitempty"`
	State      string `json:"state,omitempty"`
	StartTime  int    `json:"startTime,omitempty"`
	StopTime   int    `json:"stopTime,omitempty"`
}
