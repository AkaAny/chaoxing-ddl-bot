package response

type DBException struct {
	msg string
	err error
}

func NewDBException(msg string, err error) DBException {
	return DBException{msg: msg, err: err}
}

func (e DBException) Error() string {
	return e.msg + "(" + e.err.Error() + ")"
}
