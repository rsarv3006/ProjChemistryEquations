package ChemistryBackbone

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func SimpleReturn(x int) int { return x }

const AdvancedEquations = "AdvancedEquations"

type EquationId string

const (
	Density1 EquationId = "Density1"
	Density2 EquationId = "Density2"
	Density3 EquationId = "Density3"
)

// type EquationSection struct {
// 	Id        string
// 	Name      string
// 	Equations []Equation
// }

type EquationCalculation func([]float64) float64

// type Equation struct {
// 	Id          string
// 	Title       string
// 	Description string
// 	Filters     []EquationFilter
// 	Calculation EquationCalculation
// 	FieldLabels []string
// }

func GetEquations() ([]byte, error) {
	equationSections := &EquationSectionList{
		Sections: []*EquationSection{&DensityEquations},
	}

	return proto.Marshal(equationSections)
}

var DensityEquations = EquationSection{
	Id:   "Density",
	Name: "Density",
	Equations: []*Equation{
		&Equation{
			Id:          "Density1",
			Title:       "ρ = m / V",
			Description: DENSITY_EQUATION_DESCRIPTION,
			Filters:     make([]string, 0),
			FieldLabels: []string{
				"Mass:",
				"Volume:",
				"Density:",
			},
		},
		&Equation{
			Id:          "Density2",
			Title:       "m = ρ * V",
			Description: DENSITY_EQUATION_DESCRIPTION,
			Filters:     []string{AdvancedEquations},
			FieldLabels: []string{
				"Density:",
				"Volume:",
				"Mass:",
			},
		},
		&Equation{
			Id:          "Density3",
			Title:       "V = ρ / m",
			Description: DENSITY_EQUATION_DESCRIPTION,
			Filters:     []string{AdvancedEquations},
			FieldLabels: []string{
				"Density:",
				"Mass:",
				"Volume:",
			},
		},
	},
}

// Calculation: func(x []float64) float64 { return x[0] / x[1] },
// Calculation: func(x []float64) float64 { return x[0] * x[1] },
// Calculation: func(x []float64) float64 { return x[0] / x[1] },
const DENSITY_EQUATION_DESCRIPTION = "The density of a substance is the mass of a substance divided by the volume of the substance."
