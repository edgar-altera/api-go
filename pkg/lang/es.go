package lang

func init() {

	es["StatusUnauthorizedMessage"] = "No estás autorizado"
	es["StatusInternalServerErrorMessage"] = "Error interno del servidor"
	es["StatusUnprocessableEntity"] = "Unprocessable Entity"
	es["StatusNotFound"] = "Resource not found"
	es["StatusForbidden"] = "No tienes permisos"
}

func GetInSpanish(key string) string {

	return es[key]
}