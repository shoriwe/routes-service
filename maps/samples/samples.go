package samples

import (
	"encoding/json"

	_ "embed"

	"github.com/shoriwe/routes-service/maps"
)

//go:embed imaginary_city.json
var ImaginaryCityContents []byte

var ImaginaryCity maps.Map

func loadMap(contents []byte, target *maps.Map) {
	uErr := json.Unmarshal(contents, target)
	if uErr != nil {
		panic(uErr)
	}
}

func init() {
	loadMap(ImaginaryCityContents, &ImaginaryCity)
}
