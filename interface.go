package appsearch

import (
	"context"
)

type ListAllEngines interface {
	ListAllEngines(ctx context.Context) (data []EngineDescription, err error)
}

type ListEngines interface {
	ListEngines(ctx context.Context, page Page) (data EngineResponse, err error)
}

type ListSchema interface {
	ListSchema(ctx context.Context, engineName string) (data SchemaDefinition, err error)
}

type SearchDocuments interface {
	SearchDocuments(ctx context.Context, engineName string, query interface{}, into interface{}) (err error)
}

type PatchDocuments interface {
	PatchDocuments(ctx context.Context, engineName string, documents interface{}) (res []UpdateResponse, err error)
}

type UpdateDocuments interface {
	UpdateDocuments(ctx context.Context, engineName string, documents interface{}) (res []UpdateResponse, err error)
}

type RemoveDocuments interface {
	RemoveDocuments(ctx context.Context, engineName string, documents interface{}) (res []DeleteResponse, err error)
}

// APIClient interface
type APIClient interface {
	ListSchema
	ListEngines
	ListAllEngines
	PatchDocuments
	UpdateDocuments
	RemoveDocuments
	SearchDocuments
}
