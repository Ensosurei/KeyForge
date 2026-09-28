package verifier

import "testing"

func TestVerifyPassword(t *testing.T) {
	password := "T!g3rL!ly#2024!RunS"
	res := VerifyPassword(password)

	if !res.IsSecure {
		t.Errorf("Se esperaba que la contraseña '%s' fuera marcada como segura", password)
	}
}
