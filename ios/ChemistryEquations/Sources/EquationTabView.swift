import SwiftUI

struct EquationTabView: View {
    @State private var selectedTab = "One"
    var equation: ChemistryBackbone_Equation

    var body: some View {
        TabView(selection: $selectedTab) {
            EquationScreen(equation: equation)
                .tag("One")
                .tabItem {
                    Label("Equation", systemImage: "function")
                }

            EquationDescriptionScreen(equation: equation)
                .tabItem {
                    Label("Description", systemImage: "doc.text")
                }
                .tag("Two")
        }
    }
}
