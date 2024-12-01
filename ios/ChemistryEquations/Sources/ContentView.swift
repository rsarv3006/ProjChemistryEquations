import SwiftUI
import ChemistryBackbone

func getEquationsFromProto(data: Data) -> ChemistryBackbone_EquationSectionList? {
    return try? ChemistryBackbone_EquationSectionList(serializedBytes: data)
}

func getEquationSections() -> [ChemistryBackbone_EquationSection]? {
    let errorPointer = NSErrorPointer(nilLiteral: ())
    let protoData = ChemistryBackboneGetEquations(errorPointer)
    
    guard let protoData else {
        return nil
    }
    
    let thing = getEquationsFromProto(data: protoData)
    
    guard let thing else {
        return nil
    }
    
    return thing.sections
    
}

public struct ContentView: View {
    public var body: some View {
        NavigationStack {
            HomeScreen()
        }
    }
}
