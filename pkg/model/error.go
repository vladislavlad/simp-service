package model

type Error struct {
	Message string `json:"message"`
}

func MakeError(msg string) Error {
	return Error{
		Message: msg,
	}
}
