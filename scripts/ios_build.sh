gomobile bind -target=ios -ldflags="-X main.SwiftProtoFiles=$(ls *.pb.swift)" ./ChemistryBackbone

cd ios

tuist generate

cd ..
