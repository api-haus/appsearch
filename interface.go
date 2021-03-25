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

type Search interface {
	Search(ctx context.Context, engineName string, query interface{}, into interface{}) (err error)
}

type APIClient interface {
	Search
	ListSchema
	ListEngines
	ListAllEngines
}
