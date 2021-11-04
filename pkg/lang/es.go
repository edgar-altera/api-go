package lang

func init() {

	es["StatusUnauthorizedMessage"] = "No estás autorizado"
	es["StatusInternalServerErrorMessage"] = "Error interno del servidor"
}

func GetInSpanish(key string) string {

	return es[key]
}