package gen

//go:generate oapi-codegen -generate "chi-server" resources/openapi/user.yaml > openapi/server.go
//go:generate oapi-codegen -generate "types" resources/openapi/user.yaml > openapi/type.go
//go:generate oapi-codegen -generate "spec" resources/openapi/user.yaml > openapi/spec.go
