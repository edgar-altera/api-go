package lang

func init() {

	en["StatusUnauthorized"] = "Unauthorized"
	en["StatusInternalServerError"] = "Internal Server Error"
	en["StatusUnprocessableEntity"] = "Unprocessable Entity"
	en["StatusNotFound"] = "Resource not found"
}

func GetInEnglish(key string) string {

	return en[key]
}
