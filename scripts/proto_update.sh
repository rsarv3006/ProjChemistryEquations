protoc --go_out=./ChemistryBackbone --go_opt=paths=source_relative ChemistryBackbone.proto

protoc --swift_out=. ChemistryBackbone.proto
protoc --kotlin_out=. ChemistryBackbone.proto
