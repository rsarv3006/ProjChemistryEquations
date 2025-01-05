package ChemistryBackbone

var GibbsFreeEnergyEquations = EquationSection{
	Id:   "GibbsFreeEnergy",
	Name: "Gibbs Free Energy",
	Equations: []*Equation{
		{
			Id:          "GibbsFreeEnergy1",
			Title:       "ΔG = ΔH - TΔS",
			Description: GIBBS_FREE_ENERGY_EQUATION_DESCRIPTION,
			Filters:     make([]string, 0),
			FieldLabels: []string{
				"Enthalpy:",
				"Entropy:",
				"Temperature:",
				"ΔG:",
			},
		},
		{
			Id:          "GibbsFreeEnergy2",
			Title:       "ΔH = ΔG + TΔS",
			Description: GIBBS_FREE_ENERGY_EQUATION_DESCRIPTION,
			Filters:     []string{AdvancedEquations},
			FieldLabels: []string{
				"Temperature:",
				"Entropy:",
				"ΔG:",
				"ΔH:",
			},
		},
		{
			Id:          "GibbsFreeEnergy3",
			Title:       "T = (ΔH - ΔG) / ΔS",
			Description: GIBBS_FREE_ENERGY_EQUATION_DESCRIPTION,
			Filters:     []string{AdvancedEquations},
			FieldLabels: []string{
				"ΔH:",
				"ΔG:",
				"ΔS:",
				"Temperature:",
			},
		},
		{
			Id:          "GibbsFreeEnergy4",
			Title:       "ΔS = (ΔH - ΔG) / T",
			Description: GIBBS_FREE_ENERGY_EQUATION_DESCRIPTION,
			Filters:     []string{AdvancedEquations},
			FieldLabels: []string{
				"Temperature:",
				"ΔH:",
				"ΔG:",
				"ΔS:",
			},
		},
	},
}

func CalculateGibbsFreeEnergy1(x []float64) float64 {
	val := x[0] - (x[1] * x[2])
	return val
}

func CalculateGibbsFreeEnergy2(x []float64) float64 {
	val := x[0] + (x[1] * x[2])
	return val
}

func CalculateGibbsFreeEnergy3(x []float64) float64 {
	val := (x[0] - x[1]) / x[2]
	return val
}

func CalculateGibbsFreeEnergy4(x []float64) float64 {
	val := (x[0] - x[1]) / x[2]
	return val
}

const GIBBS_FREE_ENERGY_EQUATION_DESCRIPTION = "The Gibbs free energy is the difference between the enthalpy and the entropy of a system at a constant temperature."
