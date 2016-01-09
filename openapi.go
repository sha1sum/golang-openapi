/*
	Package openapi provides type definitions for (un)marshaling OpenAPI (formerly Swagger) specifications.
*/
package openapi

type (
	// OpenAPI is the root swagger document itself
	OpenAPI struct {
		// SwaggerVersion, as of this writing, is usually "2.0"
		SwaggerVersion string `json:"swagger"`
		// Info is a section defining basic information about the API
		Info Info `json:"info"`
		// Host is the FQDN of the server for API requests
		Host string `json:"host"`
		// Schemes is a list of accepted schemes for API requests (http, https)
		Schemes []string `json:"schemes"`
		// BasePath is a path to append to the Host in all API requests, if any
		BasePath string `json:"basePath"`
		// Produces is a list of generated mime types for responses
		Produces []string `json:"produces"`
		// Paths is a map of string paths with Path specifications for API requests
		Paths map[string]map[string]Request `json:"paths"`
		// Definitions define type definitions for use in requests and responses
		Definitions map[string]Definition `json:"definitions"`
	}

	// Info is an OpenAPI section for basic information about the API
	Info struct {
		// Title is the short name of the API
		Title string `json:"title"`
		// Description is a description of the API
		Description string `json:"description"`
		// Version is the API's current version
		Version string `json:"version"`
	}

	// Request is a specification of request parameters and documentation for an HTTP verb and its request
	Request struct {
		// Summary is a title for the HTTP request
		Summary string `json:"summary"`
		// Description is a description of the HTTP request
		Description string `json:"description"`
		// Parameters is a list of parameters for the HTTP request
		Parameters []Parameter `json:"parameters"`
		// Tags is an array of tags for the request
		Tags []string `json:"tags"`
		// Responses houses possible responses to a requestj
		Responses Responses `json:"responses"`
	}

	// Parameter is an HTTP request parameter
	Parameter struct {
		// Name is a name for the parameter
		Name string `json:"name"`
		// In defines where the parameter is specified. One of "query", "header", "path", "formData", or "body"
		In string `json:"in"`
		// Description is a text description of the parameter
		Description string `json:"description"`
		// Required indicates whether the parameter is required or not
		Required bool `json:"required"`
		// Type defines the type of data being inputted and is not needed if Query=="body". The value MUST
		// be one of "string", "number", "integer", "boolean", "array" or "file". If type is "file", the consumes
		// MUST be either "multipart/form-data", " application/x-www-form-urlencoded" or both and the parameter
		// MUST be in "formData".
		Type string `json:"type"`
		// Format further specifies the format of the data in the parameter and is not needed if Query=="body".
		// See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#dataTypeFormat
		Format string `json:"format"`
	}

	// Responses is a list of possible responses to an HTTP request
	Responses struct {
		// OK is a 200 HTTP code response
		OK Response `json:"200"`
	}

	// Response is an individual response for an HTTP request
	Response struct {
		// Description is a short description of the response
		Description string `json:"description"`
		// Schema is the Schema for the returned response
		Schema Schema `json:"schema"`
	}

	// Schema is a response schema definition
	Schema struct {
		// Type is the type of response returned
		Type string `json:"type"`
		// Items is an item type reference for the response
		Items ItemRef `json:"items"`
	}

	// ItemRef is a reference to the type of item used in a request or a response
	ItemRef struct {
		Ref string `json:"$ref"`
	}

	// Definition is a type definition for an HTTP request or response
	Definition struct {
		// Type is the type of data being used (usually "object", "string", "integer", "array")
		Type string `json:"type"`
		// Properties is a list of properties of an object
		Properties map[string]Property `json:"properties"`
	}

	// Property is a definition of a property of a Definition
	Property struct {
		// Type is a data type. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types
		Type string `json:"type"`
		// Format is the data type format for an Property
		Format string `json:"format"`
		// Description is a description of the property
		Description string `json:"description"`
		// Items is an option ItemRef for the items in the Property if the Property references another item type
		Items ItemRef `json:"items"`
	}
)
