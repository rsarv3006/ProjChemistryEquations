import ProjectDescription

let project = Project(
    name: "ChemistryEquations",
    targets: [
        .target(
            name: "ChemistryEquations",
            destinations: .iOS,
            product: .app,
            bundleId: "io.tuist.ChemistryEquations",
            infoPlist: .extendingDefault(
                with: [
                    "UILaunchStoryboardName": "LaunchScreen.storyboard",
                ]
            ),
            sources: ["ChemistryEquations/Sources/**"],
            resources: ["ChemistryEquations/Resources/**"],
            dependencies: []
        ),
        .target(
            name: "ChemistryEquationsTests",
            destinations: .iOS,
            product: .unitTests,
            bundleId: "io.tuist.ChemistryEquationsTests",
            infoPlist: .default,
            sources: ["ChemistryEquations/Tests/**"],
            resources: [],
            dependencies: [.target(name: "ChemistryEquations")]
        ),
    ]
)
