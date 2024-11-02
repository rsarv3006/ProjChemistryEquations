package ChemistryBackbone

import "fmt"

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type EquationFilter string

const (
	AdvancedEquations EquationFilter = "AdvancedEquations"
)

type EquationId string

const (
	Density1 EquationId = "Density1"
	Density2 EquationId = "Density2"
	Density3 EquationId = "Density3"
)

type EquationSection struct {
	Id        string
	Name      string
	Equations []Equation
}

type EquationCalculation func([]float64) float64

type Equation struct {
	Id          EquationId
	Title       string
	Description string
	Filters     []EquationFilter
	Calculation EquationCalculation
	FieldLabels []string
}

var DensityEquations = EquationSection{
	Id:   "Density",
	Name: "Density",
	Equations: []Equation{
		Equation{
			Id:          Density1,
			Title:       "ρ = m / V",
			Description: DENSITY_EQUATION_DESCRIPTION,
			Filters:     make([]EquationFilter, 0),
			Calculation: func(x []float64) float64 { return x[0] / x[1] },
			FieldLabels: []string{
				"Mass:",
				"Volume:",
				"Density:",
			},
		},
		Equation{
			Id:          Density2,
			Title:       "m = ρ * V",
			Description: DENSITY_EQUATION_DESCRIPTION,
			Filters:     []EquationFilter{AdvancedEquations},
			Calculation: func(x []float64) float64 { return x[0] * x[1] },
			FieldLabels: []string{
				"Density:",
				"Volume:",
				"Mass:",
			},
		},
		Equation{
			Id:          Density3,
			Title:       "V = ρ / m",
			Description: DENSITY_EQUATION_DESCRIPTION,
			Filters:     []EquationFilter{AdvancedEquations},
			Calculation: func(x []float64) float64 { return x[0] / x[1] },
			FieldLabels: []string{
				"Density:",
				"Mass:",
				"Volume:",
			},
		},
	},
}

const DENSITY_EQUATION_DESCRIPTION = "The density of a substance is the mass of a substance divided by the volume of the substance."
