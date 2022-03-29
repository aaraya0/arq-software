package div //si algo empieza con minuscula es privado al package
import "errors"

func Division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir")
	}
	return a / b, nil
}
