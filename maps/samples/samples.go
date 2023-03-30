package samples

import (
	"encoding/json"
	"os"

	"github.com/shoriwe/routes-service/maps"
)

var ImaginaryCity maps.Map

func loadMap(fileName string, target *maps.Map) {
	file, oErr := os.Open("imaginary_city.json")
	if oErr != nil {
		panic(oErr)
	}
	uErr := json.NewDecoder(file).Decode(target)
	if uErr != nil {
		panic(uErr)
	}
}

func init() {
	loadMap("imaginary_city.json", &ImaginaryCity)
}
