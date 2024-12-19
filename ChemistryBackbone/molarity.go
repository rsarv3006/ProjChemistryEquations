package ChemistryBackbone

var MolarityEquations = EquationSection{
	Id:   "Molarity",
	Name: "Molarity",
	Equations: []*Equation{
		{
			Id:          "Molarity1",
			Title:       "M = moles of solute / liters of solution",
			Description: MOLARITY_EQUATION_DESCRIPTION,
			Filters:     make([]string, 0),
			FieldLabels: []string{
				"Moles of Solute:",
				"Liters of Solution:",
				"Molarity:",
			},
		},
		{
			Id:          "Molarity2",
			Title:       "moles of solute = M * liters of solution",
			Description: MOLARITY_EQUATION_DESCRIPTION,
			Filters:     []string{},
			FieldLabels: []string{
				"Moles:",
				"liters of Solution:",
				"moles of Solute:",
			},
		},
		{
			Id:          "Molarity3",
			Title:       "liters of solution = M / moles of solute",
			Description: MOLARITY_EQUATION_DESCRIPTION,
			Filters:     []string{},
			FieldLabels: []string{
				"Moles of Solute:",
				"Moles:",
				"Liters of Solution:",
			},
		},
	},
}

func CalculateMolarity1(x []float64) float64 {
	val := x[0] / x[1]
	return val
}

func CalculateMolarity2(x []float64) float64 {
	val := x[0] * x[1]
	return val
}

func CalculateMolarity3(x []float64) float64 {
	val := x[0] / x[1]
	return val
}

const MOLARITY_EQUATION_DESCRIPTION = "The molarity of a substance is the number of moles of a substance divided by the mass of the substance."
