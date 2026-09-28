package verifier

import (
	"regexp"
	"strings"
	"unicode"
)

// CheckResult representa el estado individual de cada validación
type CheckResult struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Passed      bool   `json:"passed"`
	Description string `json:"description"`
}

// VerificationResponse es el objeto final que el API retornará al frontend
type VerificationResponse struct {
	Score    int           `json:"score"`    // Porcentaje de 0 a 100
	Strength string        `json:"strength"` // Débil, Media, Fuerte, Muy Fuerte
	IsSecure bool          `json:"isSecure"` // Cumple con todos los criterios mínimos
	Checks   []CheckResult `json:"checks"`   // Lista de checks con sus estados
	Feedback []string      `json:"feedback"` // Sugerencias o advertencias
}

// Palabras comunes a rechazar ("Evita lo obvio")
var commonWords = []string{
	"password", "123456", "qwerty", "admin", "contraseña", "secreta", "welcome",
}

// VerifyPassword analiza una contraseña y retorna los checks y su nivel de fortaleza
func VerifyPassword(password string) VerificationResponse {
	checks := make([]CheckResult, 0)
	feedback := make([]string, 0)

	// 1. Check de Longitud (Mínimo 12 caracteres)
	hasLength := len([]rune(password)) >= 12
	checks = append(checks, CheckResult{
		ID:          "length",
		Label:       "Longitud suficiente",
		Passed:      hasLength,
		Description: "La contraseña debe contener al menos 12 caracteres.",
	})
	if !hasLength {
		feedback = append(feedback, "Incrementa la longitud a mínimo 12 caracteres.")
	}

	// 2. Check de Variedad (Mayúsculas, Minúsculas, Números y Símbolos)
	var hasUpper, hasLower, hasDigit, hasSymbol bool
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSymbol = true
		}
	}

	hasVariety := hasUpper && hasLower && hasDigit && hasSymbol
	checks = append(checks, CheckResult{
		ID:          "variety",
		Label:       "Variedad de caracteres",
		Passed:      hasVariety,
		Description: "Debe combinar letras mayúsculas, minúsculas, números y símbolos.",
	})
	if !hasVariety {
		feedback = append(feedback, "Añade una mezcla de mayúsculas, minúsculas, números y símbolos.")
	}

	// 3. Check de "Evita lo Obvio"
	isObvious := checkIsObvious(password)
	passedObvious := !isObvious
	checks = append(checks, CheckResult{
		ID:          "avoid_obvious",
		Label:       "Evita patrones obvios",
		Passed:      passedObvious,
		Description: "No debe contener palabras comunes ni secuencias sencillas.",
	})
	if isObvious {
		feedback = append(feedback, "Evita palabras comunes o secuencias repetitivas.")
	}

	// Cálculo de Fortaleza y Puntaje
	passedCount := 0
	for _, c := range checks {
		if c.Passed {
			passedCount++
		}
	}

	score := 0
	if len(password) > 0 {
		score = (passedCount * 25)
		if len(password) >= 16 && hasVariety {
			score += 25
		}
		if score > 100 {
			score = 100
		}
	}

	strength := "Débil"
	switch {
	case score >= 90:
		strength = "Muy Fuerte"
	case score >= 75:
		strength = "Fuerte"
	case score >= 50:
		strength = "Media"
	}

	isSecure := hasLength && hasVariety && passedObvious

	return VerificationResponse{
		Score:    score,
		Strength: strength,
		IsSecure: isSecure,
		Checks:   checks,
		Feedback: feedback,
	}
}

func checkIsObvious(password string) bool {
	lowerPass := strings.ToLower(password)

	for _, word := range commonWords {
		if strings.Contains(lowerPass, word) {
			return true
		}
	}

	matchSeq, _ := regexp.MatchString(`(.)\1{3,}`, lowerPass)
	return matchSeq
}
