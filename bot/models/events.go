package models

type InitializeEvent struct {
	Nickname string `json:"nickname"`
	Status   int8   `json:"status"`
}

type CreateUserEvent struct {
	Status int8 `json:"status"`
}

type DeleteUserEvent struct {
	Status int8 `json:"status"`
}
