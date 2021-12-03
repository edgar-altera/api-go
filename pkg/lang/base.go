package lang

var en map[string] string
var es map[string] string

func init() {

	en = make(map[string] string) 
	es = make(map[string] string)

}

func Get(key, lang string ) string {

	var message string

	switch lang {
	
		case "en":

			message = GetInEnglish(key)

		case "es":
				
			message = GetInSpanish(key)

		default:

			message = GetInEnglish(key)
	}

	return message
}
