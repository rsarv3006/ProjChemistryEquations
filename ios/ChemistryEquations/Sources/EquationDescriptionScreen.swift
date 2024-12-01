import LaTeXSwiftUI
import SwiftUI

struct EquationDescriptionScreen: View {
    var equation: ChemistryBackbone_Equation

    var body: some View {
        ScrollView {
            LaTeX(equation.title)
                .font(.title2)
                .foregroundColor(Color.accentColor)
                .padding()

            LaTeX(equation.description_p)
                .foregroundColor(Color.accentColor)
                .padding()
        }
    }
}
