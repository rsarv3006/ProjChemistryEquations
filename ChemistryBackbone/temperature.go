package ChemistryBackbone

var TemperatureEquations = EquationSection{
	Id:   "Temperature",
	Name: "Temperature",
	Equations: []*Equation{
		{
			Id:          "Temperature1",
			Title:       "C = (F - 32) x 5/9",
			Description: TEMPERATURE_EQUATION_DESCRIPTION,
			Filters:     make([]string, 0),
			FieldLabels: []string{
				"F:",
				"C:",
			},
		},
		{
			Id:          "Temperature2",
			Title:       "F = C x 9/5 + 32",
			Description: TEMPERATURE_EQUATION_DESCRIPTION,
			Filters:     make([]string, 0),
			FieldLabels: []string{
				"C:",
				"F:",
			},
		},
		{
			Id:          "Temperature3",
			Title:       "K = C + 273.15",
			Description: TEMPERATURE_EQUATION_DESCRIPTION,
			Filters:     make([]string, 0),
			FieldLabels: []string{
				"C:",
				"K:",
			},
		},
	},
}

func CalculateTemperature1(x []float64) float64 {
	val := (x[0] - 32) * 5 / 9
	return val
}

func CalculateTemperature2(x []float64) float64 {
	val := (x[0] * 9 / 5) + 32
	return val
}

func CalculateTemperature3(x []float64) float64 {
	val := x[0] + 273.15
	return val
}

const TEMPERATURE_EQUATION_DESCRIPTION = "The temperature of a system is the average kinetic energy of the particles in the system."
