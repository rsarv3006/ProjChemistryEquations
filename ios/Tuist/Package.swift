// swift-tools-version: 5.9
import PackageDescription

#if TUIST
    import ProjectDescription

    let packageSettings = PackageSettings(
        // Customize the product types for specific package product
        // Default is .staticFramework
        // productTypes: ["Alamofire": .framework,]
        productTypes: [
            "SwiftProtobuf": .framework,
            "LaTeXSwiftUI": .framework,
        ]
    )
#endif

let package = Package(
    name: "ChemistryEquations",
    dependencies: [
        .package(url: "https://github.com/apple/swift-protobuf.git", from: "1.28.0"),
        .package(url: "https://github.com/colinc86/LaTeXSwiftUI.git", from: "1.3.2"),
    ]
)
