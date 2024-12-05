package ChemistryBackbone

var DensityEquations = EquationSection{
	Id:   "Density",
	Name: "Density",
	Equations: []*Equation{
		{
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
		{
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
		{
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

func CalculateDensity1(x []float64) float64 {
	val := x[0] / x[1]
	return val
}

func CalculateDensity2(x []float64) float64 {
	val := x[0] * x[1]
	return val
}

func CalculateDensity3(x []float64) float64 {
	val := x[0] / x[1]
	return val
}

const DENSITY_EQUATION_DESCRIPTION = "The density of a substance is the mass of a substance divided by the volume of the substance."
