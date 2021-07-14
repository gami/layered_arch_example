package domain

type Error struct {
	msg string
}

type errcode string

const (
	NotFound        errcode = "NotFound"
	InvalidArgument errcode = "InvalidArgument"
	Forbidden       errcode = "Forbidden"
	Internal        errcode = "Unknown"
)

func Invalid(msg string) Error {
	return Error{
		msg: msg,
	}
}

func (e Error) Error() string {
	return e.msg
}
