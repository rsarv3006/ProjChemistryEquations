import SwiftUI
import ChemistryBackbone

func getEquationsFromProto(data: Data) -> ChemistryBackbone_EquationSectionList? {
    return try? ChemistryBackbone_EquationSectionList(serializedData: data)
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
    var sections: [ChemistryBackbone_EquationSection] = []
    public init() {
        sections = getEquationSections() ?? []
    }

    public var body: some View {
        ForEach(sections, id: \.id) { section in
            ForEach(section.equations, id: \.id) { equation in
                Text(equation.title)
            }
        }
    }
}


struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
