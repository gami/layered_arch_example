package gen

//go:generate oapi-codegen -generate "chi-server" ../resource/openapi/user.yaml > openapi/server.go
//go:generate oapi-codegen -generate "types" ../resource/openapi/user.yaml > openapi/type.go
//go:generate oapi-codegen -generate "spec" ../resource/openapi/user.yaml > openapi/spec.go
