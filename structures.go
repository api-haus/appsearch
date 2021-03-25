package appsearch

import (
	"fmt"
	"strings"
)

type m = map[string]interface{}

type Page struct {
	Page int `json:"current"`
	Size int `json:"size"`
}

// Response for Patch or Update operations
type UpdateResponse struct {
	// Updated document ID
	ID string `json:"id"`
	// List of errorss
	Errors string `json:"errors"`
}

// Response for Patch or Update operations
type DeleteResponse struct {
	// Deleted document ID
	ID string `json:"id"`
	// Was document deleted successfully
	Deleted bool `json:"deleted"`
	// List of errors
	Errors string `json:"errors"`
}

// Schema type defines 4 types of value: text (""), date (time.RFC3339), number (0) and geolocation ("0.0,0.0")
type SchemaType = string

const (
	SchemaTypeText        = "text"
	SchemaTypeDate        = "date"
	SchemaTypeNumber      = "number"
	SchemaTypeGeolocation = "geolocation"
)

// Schema definition as map[string]SchemaType
// "id" field of "text" type is added to schema automatically (non-standard behaviour).
type SchemaDefinition map[string]SchemaType

// ListEngines response
type EngineResponse struct {
	Meta struct {
		Page struct {
			TotalPages int `json:"total_pages"`
		} `json:"page"`
	} `json:"meta"`
	Results []EngineDescription `json:"results"`
}

// Engine description
type EngineDescription struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Language      *string `json:"language"`
	DocumentCount int     `json:"document_count"`
}

// API Error
type Error struct {
	Message    string   `json:"error"`
	Messages   []string `json:"errors"`
	StatusCode int
}

func (e *Error) Error() string {
	if len(e.Messages) > 0 {
		return strings.Join(e.Messages, ", ")
	}

	if e.Message != "" {
		return e.Message
	}

	return fmt.Sprintf("HTTP [%d]", e.StatusCode)
}
