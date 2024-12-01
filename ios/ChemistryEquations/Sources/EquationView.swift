import SwiftUI

struct EquationView: View {
    let equation: ChemistryBackbone_Equation
    
    var body: some View {
        Text(equation.title)
            .font(.title)
            .padding()
        
        
    }
}
