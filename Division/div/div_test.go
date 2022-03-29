package div

import "testing"

func TestDivision(t *testing.T) {
	_, err := Division(8, 4)

	if err != nil {
		t.Error(err)
		return
	}
}
func TestDivisionErrorByZero(t *testing.T) {
	_, err := Division(8, 0)

	if err == nil {
		t.Error(err)
		return
	}
}

// go test ./... para correr los tests

// go test ./... -v para correr los tests con mas info
