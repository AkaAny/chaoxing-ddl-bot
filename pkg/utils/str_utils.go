package utils

func WrapWith(str string, wrap string) string {
	return wrap + str + wrap
}

func WrapWithQuote(str string) string {
	return WrapWith(str, "\"")
}
