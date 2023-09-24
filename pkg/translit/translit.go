package translit

import (
	"github.com/gosimple/slug"
)

func init() {
	slug.CustomRuneSub = make(map[rune]string)
	slug.CustomRuneSub['#'] = "sharp"
}

func Make(text string) string {
	return slug.Make(text)
}
