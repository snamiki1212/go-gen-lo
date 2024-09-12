package lib

func ToUpperFirst(str string) string {
	if str == "" {
		return str
	}
	return string(str[0]-32) + str[1:]
}
