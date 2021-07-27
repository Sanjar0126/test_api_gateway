package models

//ResponseSuccess ...
type ResponseSuccess struct {
	Metadata interface{}
	Data     interface{}
}

//ResponseError ...
type ResponseError struct {
	Error interface{}
}

//InternalServerError ...
type InternalServerError struct {
	Code    string
	Message string
}

//ValidationError ...
type ValidationError struct {
	Code        string
	Message     string
	UserMessage string
}

type ResponseOK struct {
	Message interface{}
}

type Response struct {
	ID interface{} `json:"id"`
}
