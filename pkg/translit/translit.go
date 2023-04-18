package translit

import "bytes"

var ruEn = map[string]string{
	"а": "a", "А": "A", "Б": "B", "б": "b", "В": "V", "в": "v", "Г": "G", "г": "g",
	"Д": "D", "д": "d", "Е": "Je", "е": "je", "Ё": "Jo", "ё": "Jo", "Ж": "Zj", "ж": "zj", "З": "Z",
	"з": "z", "И": "I", "и": "i", "К": "K", "к": "k", "Л": "L", "л": "l", "М": "M",
	"м": "m", "Н": "N", "н": "n", "О": "O", "о": "o", "П": "P", "п": "p", "Р": "R",
	"р": "r", "С": "S", "с": "s", "Т": "T", "т": "t", "У": "U", "у": "u", "Ф": "F",
	"ф": "f", "Х": "X", "х": "x", "Ц": "c", "ц": "c", "Ч": "Ch", "ч": "ch", "Ш": "Sh", "ш": "sh",
	"Щ": "Shh", "щ": "shh", "Ъ": "", "ъ": "", "Ы": "Y", "ы": "y`", "Ь": "", "ь": "",
	"Э": "E", "э": "e", "Ю": "Yu", "ю": "yu", "Я": "Ya", "я": "ya", " ": "_",
}

func RuEn(text string) string {
	var input = bytes.NewBufferString(text)
	var output = bytes.NewBuffer(nil)

	for {
		r, _, err := input.ReadRune()

		if err != nil {
			break
		}

		if ch, ok := ruEn[string(r)]; ok {
			output.WriteString(ch)
			continue
		}

		output.WriteRune(r)
	}

	return output.String()

}