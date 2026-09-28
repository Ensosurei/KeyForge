package generator

import (
	"crypto/rand"
	"math/big"
	"strings"
	"unicode"
)

type GeneratorOptions struct {
	Length         int  `json:"length"`
	IncludeUpper   bool `json:"includeUpper"`
	IncludeLower   bool `json:"includeLower"`
	IncludeDigits  bool `json:"includeDigits"`
	IncludeSymbols bool `json:"includeSymbols"`
}

type PassphraseOptions struct {
	WordCount   int    `json:"wordCount"`
	Separator   string `json:"separator"`
	Capitalize  bool   `json:"capitalize"`
	EnableLeet  bool   `json:"enableLeet"`
}

const (
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars  = "0123456789"
	symbolChars = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// Banco ampliado de palabras en español
var passphraseWords = []string{
	"fuego", "codigo", "llave", "fuerza", "sombra", "trueno", "dragon",
	"viento", "tormenta", "cristal", "guardian", "secreto", "vortex",
	"galaxia", "orbita", "prisma", "escudo", "faro", "sendero", "abismo",
	"horizonte", "enigma", "cometa", "relampago", "espejo", "legado", "umbra",
	"nebulosa", "santuario", "matriz", "vector", "nucleo", "centinela",
}

func GeneratePassword(opts GeneratorOptions) string {
	if opts.Length < 8 {
		opts.Length = 8
	}

	var charset string
	var required []byte

	if opts.IncludeLower {
		charset += lowerChars
		required = append(required, getRandomChar(lowerChars))
	}
	if opts.IncludeUpper {
		charset += upperChars
		required = append(required, getRandomChar(upperChars))
	}
	if opts.IncludeDigits {
		charset += digitChars
		required = append(required, getRandomChar(digitChars))
	}
	if opts.IncludeSymbols {
		charset += symbolChars
		required = append(required, getRandomChar(symbolChars))
	}

	if charset == "" {
		charset = lowerChars + digitChars
	}

	result := make([]byte, opts.Length)
	copy(result, required)

	for i := len(required); i < opts.Length; i++ {
		result[i] = getRandomChar(charset)
	}

	for i := range result {
		j := getRandomInt(len(result))
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// GeneratePassphrase crea frases secretas avanzadas con leet speak y sin palabras repetidas
func GeneratePassphrase(opts PassphraseOptions) string {
	if opts.WordCount < 3 {
		opts.WordCount = 3
	}
	if opts.Separator == "" {
		opts.Separator = "-"
	}

	selected := make([]string, 0, opts.WordCount)
	lastWord := ""

	for len(selected) < opts.WordCount {
		idx := getRandomInt(len(passphraseWords))
		word := passphraseWords[idx]

		// Evitar palabras repetidas consecutivas
		if word == lastWord {
			continue
		}

		lastWord = word

		// Aplicar mayúscula inicial si está activado
		if opts.Capitalize {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			word = string(runes)
		}

		// Aplicar Leet Speak (Sustitución de caracteres) si está activado
		if opts.EnableLeet {
			word = applyLeetSpeak(word)
		}

		selected = append(selected, word)
	}

	return strings.Join(selected, opts.Separator)
}

func applyLeetSpeak(word string) string {
	replacer := strings.NewReplacer(
		"a", "4", "A", "4",
		"e", "3", "E", "3",
		"i", "1", "I", "1",
		"o", "0", "O", "0",
		"t", "7", "T", "7",
		"s", "5", "S", "5",
	)
	return replacer.Replace(word)
}

func getRandomChar(charset string) byte {
	idx := getRandomInt(len(charset))
	return charset[idx]
}

func getRandomInt(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0
	}
	return int(nBig.Int64())
}
