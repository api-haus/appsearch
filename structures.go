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
type UpdateResponse struct {
	ID     string `json:"id"`
	Errors string `json:"errors"`
}

type SchemaDefinition map[string]string

type EngineResponse struct {
	Meta struct {
		Page struct {
			TotalPages int `json:"total_pages"`
		} `json:"page"`
	} `json:"meta"`
	Results []EngineDescription `json:"results"`
}

type EngineDescription struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Language      *string `json:"language"`
	DocumentCount int     `json:"document_count"`
}

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
