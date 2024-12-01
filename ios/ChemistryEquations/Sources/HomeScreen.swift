import LaTeXSwiftUI
import SwiftUI

struct HomeScreen: View {
    var sections: [ChemistryBackbone_EquationSection] = []
    public init() {
        sections = getEquationSections() ?? []
    }
    
    @EnvironmentObject var storekitStore: StorekitStore
    @State private var selectedEquation: ChemistryBackbone_Equation? = nil
    @State private var isLoadingStorekit = true
    
    var body: some View {
        if !isLoadingStorekit {
            VStack {
                ZStack {
                    HStack {
                        Spacer()
                        
                        NavigationLink(destination: SettingsScreen()) {
                            Image(systemName: "gear.circle")
                                .font(.title2)
                                .foregroundStyle(Color.accentColor)
                        }
                        .padding([.trailing])
                    }
                    
                    Text("Chemistry Equations")
                        .font(.title2)
                        .foregroundStyle(Color.accentColor)
                }
                
                Rectangle()
                    .fill(Color.accentColor)
                    .frame(height: 1)
                
                ScrollView {
                    ForEach(sections, id: \.name) { section in
                        Section(header:
                                    HStack {
                            Text(section.name)
                                .font(.title3)
                                .foregroundStyle(Color.accentColor)
                                .padding(.horizontal)
                            Spacer()
                        }) {
                            ForEach(section.equations, id: \.id) { equation in
                                HomeScreenEquationButton(equation: equation, selectedEquation: $selectedEquation)
                            }
                            .padding(.bottom)
                        }
                    }
                }
                .navigationDestination(isPresented: Binding<Bool>(
                    get: { selectedEquation != nil },
                    set: { _ in selectedEquation = nil }
                )) {
                    if let equation = selectedEquation {
                        EquationTabView(equation: equation)
                    }
                }
            }
        } else {
            ProgressView()
                .onAppear {
                    Task {
                        await storekitStore.updateCustomerProductStatus()
                        isLoadingStorekit = false
                    }
                }
        }
    }
}
