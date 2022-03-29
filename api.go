package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"os"
)

type Categories []Category
type Category struct {
	ID   string `json:"id"`
	Name string `json: "name"`
}

func main() {
	cats, err := GetCategories("MLA")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Las categorias son ", cats)
}

func GetCategories(siteID string) (Categories, error) {
	response, err := http.Get("https://api.mercadolibre.com/sites")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	var cats Categories
	json.Unmarshal(bytes, &cats)
	return cats, nil
}

// Preguntar como visualizar como texto separado en vez de strings
