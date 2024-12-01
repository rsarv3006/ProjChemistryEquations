package ChemistryBackbone

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

const AdvancedEquations = "AdvancedEquations"

const (
	Density1 string = "Density1"
	Density2 string = "Density2"
	Density3 string = "Density3"
)

type EquationCalculation func([]float64) float64

func CalculateEquation(bytes []byte) ([]byte, error) {
	request := &EquationCalculationRequest{}
	err := proto.Unmarshal(bytes, request)

	if err != nil {
		return nil, err
	}

	calculatedValue := 0.0

	switch request.Id {
	case Density1:
		calculatedValue = CalculateDensity1(request.Values)
	case Density2:
		calculatedValue = CalculateDensity2(request.Values)
	case Density3:
		calculatedValue = CalculateDensity3(request.Values)
	default:
		return nil, errors.New("Equation not found")
	}

	response := &EquationCalculationResponse{
		Value: calculatedValue,
	}

	return proto.Marshal(response)
}

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
