package lang

func init() {

	es["StatusUnauthorizedMessage"] = "No estás autorizado"
	es["StatusInternalServerErrorMessage"] = "Error interno del servidor"
	en["StatusUnprocessableEntity"] = "Unprocessable Entity"
	en["StatusNotFound"] = "Resource not found"
}

func GetInSpanish(key string) string {

	return es[key]
}