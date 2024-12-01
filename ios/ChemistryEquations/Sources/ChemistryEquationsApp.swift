import SwiftUI

@main
struct ChemistryEquationsApp: App {
    @StateObject private var storekitStore = StorekitStore()
    var body: some Scene {
        WindowGroup {
            ContentView()
                .environmentObject(storekitStore)
        }
    }
}
