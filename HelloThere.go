package hellothere

const (
	CommaString       = ","
	MissingNameString = ""
	SpaceString       = " "
)

type Language struct {
	Greeting    string
	NiceDay     string
	DefaultName string
}

var English = Language{
	Greeting:    "Hello",
	NiceDay:     "have a nice day!",
	DefaultName: "User",
}

var Hinglish = Language{
	Greeting:    "Namaste",
	NiceDay:     "aapka din shubh ho!",
	DefaultName: "Meri Jaan",
}

var Spanish = Language{
	Greeting:    "Hola",
	NiceDay:     "¡que tengas un gran día!",
	DefaultName: "Usuario",
}

var SupportedLanguages = map[string]Language{
	"English":  English,
	"Hinglish": Hinglish,
	"Spanish":  Spanish,
}

func SayHello(name string, lang string) string {

	var greetingString, nameString, niceDayString string

	currentLanguage, ok := SupportedLanguages[lang]
	if !ok {
		currentLanguage = English
	}

	greetingString = currentLanguage.Greeting
	nameString = name
	if nameString == MissingNameString {
		nameString = currentLanguage.DefaultName
	}
	niceDayString = currentLanguage.NiceDay

	return greetingString + SpaceString + nameString + CommaString + SpaceString + niceDayString
}
