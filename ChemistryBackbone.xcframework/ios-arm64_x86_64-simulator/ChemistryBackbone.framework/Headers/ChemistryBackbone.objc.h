// Objective-C API for talking to ChemistryBackbone/ChemistryBackbone Go package.
//   gobind -lang=objc ChemistryBackbone/ChemistryBackbone
//
// File is generated by gobind. Do not edit.

#ifndef __ChemistryBackbone_H__
#define __ChemistryBackbone_H__

@import Foundation;
#include "ref.h"
#include "Universe.objc.h"


@class ChemistryBackboneEquation;
@class ChemistryBackboneEquationCalculationRequest;
@class ChemistryBackboneEquationCalculationResponse;
@class ChemistryBackboneEquationSection;
@class ChemistryBackboneEquationSectionList;

@interface ChemistryBackboneEquation : NSObject <goSeqRefInterface> {
}
@property(strong, readonly) _Nonnull id _ref;

- (nonnull instancetype)initWithRef:(_Nonnull id)ref;
- (nonnull instancetype)init;
@property (nonatomic) NSString* _Nonnull id_;
@property (nonatomic) NSString* _Nonnull title;
@property (nonatomic) NSString* _Nonnull description;
// skipped field Equation.Filters with unsupported type: []string

// skipped field Equation.FieldLabels with unsupported type: []string

// skipped method Equation.Descriptor with unsupported parameter or return types

- (NSString* _Nonnull)getDescription;
// skipped method Equation.GetFieldLabels with unsupported parameter or return types

// skipped method Equation.GetFilters with unsupported parameter or return types

- (NSString* _Nonnull)getId;
- (NSString* _Nonnull)getTitle;
- (void)protoMessage;
// skipped method Equation.ProtoReflect with unsupported parameter or return types

- (void)reset;
- (NSString* _Nonnull)string;
@end

@interface ChemistryBackboneEquationCalculationRequest : NSObject <goSeqRefInterface> {
}
@property(strong, readonly) _Nonnull id _ref;

- (nonnull instancetype)initWithRef:(_Nonnull id)ref;
- (nonnull instancetype)init;
@property (nonatomic) NSString* _Nonnull id_;
// skipped field EquationCalculationRequest.Values with unsupported type: []float64

// skipped method EquationCalculationRequest.Descriptor with unsupported parameter or return types

- (NSString* _Nonnull)getId;
// skipped method EquationCalculationRequest.GetValues with unsupported parameter or return types

- (void)protoMessage;
// skipped method EquationCalculationRequest.ProtoReflect with unsupported parameter or return types

- (void)reset;
- (NSString* _Nonnull)string;
@end

@interface ChemistryBackboneEquationCalculationResponse : NSObject <goSeqRefInterface> {
}
@property(strong, readonly) _Nonnull id _ref;

- (nonnull instancetype)initWithRef:(_Nonnull id)ref;
- (nonnull instancetype)init;
@property (nonatomic) double value;
// skipped method EquationCalculationResponse.Descriptor with unsupported parameter or return types

- (double)getValue;
- (void)protoMessage;
// skipped method EquationCalculationResponse.ProtoReflect with unsupported parameter or return types

- (void)reset;
- (NSString* _Nonnull)string;
@end

@interface ChemistryBackboneEquationSection : NSObject <goSeqRefInterface> {
}
@property(strong, readonly) _Nonnull id _ref;

- (nonnull instancetype)initWithRef:(_Nonnull id)ref;
- (nonnull instancetype)init;
@property (nonatomic) NSString* _Nonnull id_;
@property (nonatomic) NSString* _Nonnull name;
// skipped field EquationSection.Equations with unsupported type: []*ChemistryBackbone/ChemistryBackbone.Equation

// skipped method EquationSection.Descriptor with unsupported parameter or return types

// skipped method EquationSection.GetEquations with unsupported parameter or return types

- (NSString* _Nonnull)getId;
- (NSString* _Nonnull)getName;
- (void)protoMessage;
// skipped method EquationSection.ProtoReflect with unsupported parameter or return types

- (void)reset;
- (NSString* _Nonnull)string;
@end

@interface ChemistryBackboneEquationSectionList : NSObject <goSeqRefInterface> {
}
@property(strong, readonly) _Nonnull id _ref;

- (nonnull instancetype)initWithRef:(_Nonnull id)ref;
- (nonnull instancetype)init;
// skipped field EquationSectionList.Sections with unsupported type: []*ChemistryBackbone/ChemistryBackbone.EquationSection

// skipped method EquationSectionList.Descriptor with unsupported parameter or return types

// skipped method EquationSectionList.GetSections with unsupported parameter or return types

- (void)protoMessage;
// skipped method EquationSectionList.ProtoReflect with unsupported parameter or return types

- (void)reset;
- (NSString* _Nonnull)string;
@end

FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneAdvancedEquations;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneDENSITY_EQUATION_DESCRIPTION;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneDensity1;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneDensity2;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneDensity3;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneGIBBS_FREE_ENERGY_EQUATION_DESCRIPTION;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneGibbsFreeEnergy1;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneGibbsFreeEnergy2;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneGibbsFreeEnergy3;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneGibbsFreeEnergy4;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneMOLARITY_EQUATION_DESCRIPTION;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneMolarity1;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneMolarity2;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneMolarity3;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneTEMPERATURE_EQUATION_DESCRIPTION;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneTemperature1;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneTemperature2;
FOUNDATION_EXPORT NSString* _Nonnull const ChemistryBackboneTemperature3;

@interface ChemistryBackbone : NSObject
// skipped variable DensityEquations with unsupported type: ChemistryBackbone/ChemistryBackbone.EquationSection

// skipped variable File_ChemistryBackbone_proto with unsupported type: google.golang.org/protobuf/reflect/protoreflect.FileDescriptor

// skipped variable GibbsFreeEnergyEquations with unsupported type: ChemistryBackbone/ChemistryBackbone.EquationSection

// skipped variable MolarityEquations with unsupported type: ChemistryBackbone/ChemistryBackbone.EquationSection

// skipped variable TemperatureEquations with unsupported type: ChemistryBackbone/ChemistryBackbone.EquationSection

@end

// skipped function CalculateDensity1 with unsupported parameter or return types


// skipped function CalculateDensity2 with unsupported parameter or return types


// skipped function CalculateDensity3 with unsupported parameter or return types


FOUNDATION_EXPORT NSData* _Nullable ChemistryBackboneCalculateEquation(NSData* _Nullable bytes, NSError* _Nullable* _Nullable error);

// skipped function CalculateGibbsFreeEnergy1 with unsupported parameter or return types


// skipped function CalculateGibbsFreeEnergy2 with unsupported parameter or return types


// skipped function CalculateGibbsFreeEnergy3 with unsupported parameter or return types


// skipped function CalculateGibbsFreeEnergy4 with unsupported parameter or return types


// skipped function CalculateMolarity1 with unsupported parameter or return types


// skipped function CalculateMolarity2 with unsupported parameter or return types


// skipped function CalculateMolarity3 with unsupported parameter or return types


// skipped function CalculateTemperature1 with unsupported parameter or return types


// skipped function CalculateTemperature2 with unsupported parameter or return types


// skipped function CalculateTemperature3 with unsupported parameter or return types


FOUNDATION_EXPORT NSData* _Nullable ChemistryBackboneGetEquations(NSError* _Nullable* _Nullable error);

#endif
