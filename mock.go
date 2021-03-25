package appsearch

import (
	"context"
)

type mock struct {
	Engines []EngineDescription
	Schemas map[string]SchemaDefinition

	MockSearch func(ctx context.Context, engineName string, query interface{}, into interface{}) (err error)
}

// Create mock APIClient
func Mock(Engines []EngineDescription, Schemas map[string]SchemaDefinition) *mock {
	return &mock{Engines: Engines, Schemas: Schemas}
}

func (m *mock) ListAllEngines(ctx context.Context) (data []EngineDescription, err error) {
	return m.Engines, nil
}

func (m *mock) ListEngines(ctx context.Context, page Page) (data EngineResponse, err error) {
	return EngineResponse{
		Results: m.Engines,
	}, nil
}

func (m *mock) ListSchema(ctx context.Context, engineName string) (data SchemaDefinition, err error) {
	return m.Schemas[engineName], nil
}

func (m *mock) SearchDocument(ctx context.Context, engineName string, query interface{}, into interface{}) (err error) {
	return m.MockSearch(ctx, engineName, query, into)
}

func (m *mock) PatchDocuments(ctx context.Context, engineName string, documents interface{}) (res []UpdateResponse, err error) {
	panic("implement me")
}

func (m *mock) UpdateDocuments(ctx context.Context, engineName string, documents interface{}) (res []UpdateResponse, err error) {
	panic("implement me")
}

func (m *mock) RemoveDocuments(ctx context.Context, engineName string, documents interface{}) (res []DeleteResponse, err error) {
	panic("implement me")
}

func (m *mock) SearchDocuments(ctx context.Context, engineName string, query interface{}, into interface{}) (err error) {
	panic("implement me")
}
