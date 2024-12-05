package ChemistryBackbone

import (
	"google.golang.org/protobuf/proto"
)

const AdvancedEquations = "AdvancedEquations"

const (
	Density1 string = "Density1"
	Density2 string = "Density2"
	Density3 string = "Density3"

	GibbsFreeEnergy1 string = "GibbsFreeEnergy1"
	GibbsFreeEnergy2 string = "GibbsFreeEnergy2"
	GibbsFreeEnergy3 string = "GibbsFreeEnergy3"
	GibbsFreeEnergy4 string = "GibbsFreeEnergy4"

	Temperature1 string = "Temperature1"
	Temperature2 string = "Temperature2"
	Temperature3 string = "Temperature3"

	Molarity1 string = "Molarity1"
	Molarity2 string = "Molarity2"
	Molarity3 string = "Molarity3"
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

	case Temperature1:
		calculatedValue = CalculateTemperature1(request.Values)
	case Temperature2:
		calculatedValue = CalculateTemperature2(request.Values)
	case Temperature3:
		calculatedValue = CalculateTemperature3(request.Values)

	case GibbsFreeEnergy1:
		calculatedValue = CalculateGibbsFreeEnergy1(request.Values)
	case GibbsFreeEnergy2:
		calculatedValue = CalculateGibbsFreeEnergy2(request.Values)
	case GibbsFreeEnergy3:
		calculatedValue = CalculateGibbsFreeEnergy3(request.Values)
	case GibbsFreeEnergy4:
		calculatedValue = CalculateGibbsFreeEnergy4(request.Values)

	case Molarity1:
		calculatedValue = CalculateMolarity1(request.Values)
	case Molarity2:
		calculatedValue = CalculateMolarity2(request.Values)
	case Molarity3:
		calculatedValue = CalculateMolarity3(request.Values)
	default:
		calculatedValue = 0.0
	}

	response := &EquationCalculationResponse{
		Value: calculatedValue,
	}

	return proto.Marshal(response)
}

func GetEquations() ([]byte, error) {
	equationSections := &EquationSectionList{
		Sections: []*EquationSection{
			&DensityEquations,
			&TemperatureEquations,
			&GibbsFreeEnergyEquations,
			&MolarityEquations,
		},
	}

	return proto.Marshal(equationSections)
}
