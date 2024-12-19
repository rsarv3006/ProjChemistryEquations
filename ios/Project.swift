import ProjectDescription

extension SettingsDictionary {
    func setProjectVersions() -> SettingsDictionary {
        return appleGenericVersioningSystem().merging([
            "ASSETCATALOG_COMPILER_GLOBAL_ACCENT_COLOR_NAME": "AccentColor"
        ])
    }
}

let project = Project(
    name: "ChemistryEquations",
    settings: Settings.settings(
        base: SettingsDictionary().setProjectVersions(),
        configurations: [
            Configuration.debug(
                name: "Debug",
                settings: SettingsDictionary().automaticCodeSigning(devTeam: "4QGR522B9M")
            ),
            Configuration.release(
                name: "Release",
                settings: SettingsDictionary().automaticCodeSigning(devTeam: "4QGR522B9M")
            ),
        ], defaultSettings: .recommended()
    ),
    targets: [
        .target(
            name: "ChemistryEquations",
            destinations: .iOS,
            product: .app,
            bundleId: "us.rjs-app-dev.ChemistryEquations",
            deploymentTargets: DeploymentTargets.iOS("16.0"),
            infoPlist: .extendingDefault(
                with: [
                    "UILaunchStoryboardName": "LaunchScreen"
                ]
            ),
            sources: ["ChemistryEquations/Sources/**"],
            resources: ["ChemistryEquations/Resources/**"],
            dependencies: [
                .xcframework(
                    path: "../ChemistryBackbone.xcframework", status: .required, condition: .none
                ),
                .external(name: "SwiftProtobuf"),
                .external(name: "LaTeXSwiftUI"),
            ]
        ),
        .target(
            name: "ChemistryEquationsTests",
            destinations: .iOS,
            product: .unitTests,
            bundleId: "us.rjs-app-dev.ChemistryEquationsTests",
            infoPlist: .default,
            sources: ["ChemistryEquations/Tests/**"],
            resources: [],
            dependencies: [.target(name: "ChemistryEquations")]
        ),
    ]
)
