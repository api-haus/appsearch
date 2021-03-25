package appsearch

import (
	"context"
)

type mock struct {
	Engines []EngineDescription
	Schemas map[string]SchemaDefinition

	MockSearch func(ctx context.Context, engineName string, query interface{}, into interface{}) (err error)
}

func Mock(Engines []EngineDescription, Schemas map[string]SchemaDefinition) *mock {
	return &mock{Engines: Engines, Schemas: Schemas}
}

func (m *mock) ListAllEngines(ctx context.Context) (data []EngineDescription, err error) {
	return m.Engines, nil
}

func (m *mock) ListSchema(ctx context.Context, engineName string) (data SchemaDefinition, err error) {
	return m.Schemas[engineName], nil
}

func (m *mock) Search(ctx context.Context, engineName string, query interface{}, into interface{}) (err error) {
	return m.MockSearch(ctx, engineName, query, into)
}
