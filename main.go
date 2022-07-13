package main

import (
	"fmt"
	"github.com/mxmCherry/translit"
	"github.com/pemistahl/lingua-go"
	"golang.org/x/text/transform"
	"strings"
)

func CheckLanguage(text string) string {
	languages := []lingua.Language{
		lingua.English,
		lingua.Russian,
	}

	detector := lingua.NewLanguageDetectorBuilder().FromLanguages(languages...).WithMinimumRelativeDistance(0.10).Build()

	languageOf, _ := detector.DetectLanguageOf(text)

	return fmt.Sprintf("%v", languageOf)
}

func Transcript(text string) string {

	CyrillicToLatin := translit.Map(
		map[string]string{
			"А": "A",
			"а": "a",
			"Б": "B",
			"б": "b",
			"В": "V",
			"в": "v",
			"Г": "G",
			"г": "g",
			"Д": "D",
			"д": "d",
			"E": "E",
			"e": "e",
			"Ё": "Yo",
			"ё": "yo",
			"Ж": "Zh",
			"ж": "zh",
			"З": "Z",
			"з": "z",
			"И": "i",
			"и": "i",
			"Й": "Y",
			"й": "y",
			"К": "K",
			"к": "k",
			"Л": "L",
			"л": "l",
			"М": "M",
			"м": "m",
			"Н": "N",
			"н": "n",
			"О": "O",
			"о": "o",
			"П": "P",
			"п": "p",
			"Р": "R",
			"р": "r",
			"С": "S",
			"с": "s",
			"Т": "T",
			"т": "t",
			"У": "U",
			"у": "u",
			"Ф": "F",
			"ф": "f",
			"Х": "Kh",
			"х": "kh",
			"Ц": "Ts",
			"ц": "ts",
			"Ы": "I",
			"ы": "i",
			"Ь": "'",
			"ь": "'",
			"Ъ": "",
			"ъ": "",
			"Ч": "Ch",
			"ч": "ch",
			"Ш": "Sh",
			"ш": "sh",
			"Щ": "Sh",
			"щ": "sh",
			"Э": "e",
			"э": "e",
			"Ю": "Yu",
			"ю": "yu",
			"Я": "Ya",
			"я": "ya",
		})

	LatinToCyrillic := translit.Map(
		map[string]string{
			"A":  "А",
			"a":  "а",
			"B":  "Б",
			"b":  "б",
			"C":  "Ц",
			"c":  "ц",
			"D":  "Д",
			"d":  "д",
			"E":  "Е",
			"e":  "е",
			"F":  "Ф",
			"f":  "ф",
			"G":  "Г",
			"g":  "г",
			"H":  "Х",
			"h":  "х",
			"I":  "И",
			"i":  "и",
			"J":  "Дж",
			"j":  "дж",
			"K":  "К",
			"k":  "к",
			"L":  "Л",
			"l":  "л",
			"M":  "М",
			"m":  "м",
			"N":  "Н",
			"n":  "н",
			"O":  "О",
			"o":  "о",
			"P":  "П",
			"p":  "п",
			"Q":  "К",
			"q":  "к",
			"R":  "Р",
			"r":  "р",
			"S":  "С",
			"s":  "с",
			"T":  "Т",
			"t":  "т",
			"U":  "У",
			"u":  "у",
			"Y":  "Й",
			"y":  "й",
			"X":  "",
			"x":  "",
			"V":  "В",
			"v":  "в",
			"Z":  "З",
			"z":  "з",
			"W":  "В",
			"w":  "в",
			"YO": "Ё",
			"Yo": "Ё",
			"yo": "ё",
			"YA": "Я",
			"Ya": "Я",
			"ya": "я",
			"YU": "Ю",
			"Yu": "Ю",
			"yu": "ю",
			"SH": "Ш",
			"Sh": "Ш",
			"sh": "ш",
			"CH": "Ч",
			"Ch": "Ч",
			"ch": "ч",
			"ZH": "Ж",
			"Zh": "Ж",
			"zh": "ж",
		})

	if strings.Compare(CheckLanguage(text), "Russian") == 0 {
		tr := CyrillicToLatin.Transformer()
		s, _, _ := transform.String(tr, text)
		return s
	}

	if strings.Compare(CheckLanguage(text), "English") == 0 {
		tr := LatinToCyrillic.Transformer()
		s, _, _ := transform.String(tr, text)
		return s
	}

	//err := errors.New("sdfasasffsd")
	return "NOT CYRILLIC OR NOT LATIN"
}

func main() {

}
