package generator

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	opts := GeneratorOptions{
		Length:         16,
		IncludeUpper:   true,
		IncludeLower:   true,
		IncludeDigits:  true,
		IncludeSymbols: true,
	}

	pass := GeneratePassword(opts)

	if len(pass) != 16 {
		t.Errorf("Se esperaba longitud de 16, se obtuvo %d", len(pass))
	}
}

func TestGeneratePassphrase(t *testing.T) {
	phrase := GeneratePassphrase(4, "-")

	if phrase == "" {
		t.Errorf("La frase secreta generada no debe estar vacía")
	}
}
