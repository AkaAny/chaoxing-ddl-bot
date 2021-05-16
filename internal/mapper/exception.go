package mapper

import "ddl-bot/pkg/exception"

func WrapDBError(err error) exception.InternalException {
	return exception.NewInternalException("error when query in db", "db", err)
}
