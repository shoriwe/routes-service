package random

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/shoriwe/routes-service/models"
)

func Plate() string {
	i, err := rand.Int(rand.Reader, big.NewInt(99999))
	if err != nil {
		panic(err)
	}
	s := i.String()
	switch len(s) {
	case 1:
		return fmt.Sprintf("AAA00%s", s)
	case 2:
		return fmt.Sprintf("AAA0%s", s)
	case 3:
		return fmt.Sprintf("AAA%s", s)
	case 4:
		return fmt.Sprintf("AA%s", s)
	case 5:
		return fmt.Sprintf("A%s", s)
	}
	panic("Invalid")
}

func Vehicle() *models.Vehicle {
	return &models.Vehicle{
		Name:  String(),
		Plate: Plate(),
	}
}
