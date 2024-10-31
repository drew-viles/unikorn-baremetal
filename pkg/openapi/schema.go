// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
