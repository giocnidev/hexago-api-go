package model

type Error struct {
	Code    string
	Message string
}

func (err *Error) Error() string {
	return err.Message
}
