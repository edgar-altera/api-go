package lang

func init() {

	en["StatusUnauthorizedMessage"] = "Unauthorized"
	en["StatusInternalServerErrorMessage"] = "Internal Server Error"
}

func GetInEnglish(key string) string {

	return en[key]
}
