package dtos

type RegisterDto struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type PromoteDto struct {
	Username string `json:"username" bson:"username"`
}
