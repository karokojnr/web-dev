package model

// import "go/types"

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
