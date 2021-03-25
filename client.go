package appsearch

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/go-resty/resty/v2"
)

type client struct {
	*resty.Client
}

// Open APIClient with endpoint and key
func Open(endpoint, key string) (*client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	u = u.ResolveReference(&url.URL{Path: "/api/as/v1/"})

	return &client{
		resty.New().
			SetHostURL(u.String()).
			SetAuthToken(key).
			SetAuthScheme("Bearer"),
	}, nil
}

// List all available engines
func (c *client) ListAllEngines(ctx context.Context) (engines []EngineDescription, err error) {
	page := 0
	totalPages := 1

	for page < totalPages {
		page += 1

		res, err := c.ListEngines(ctx, Page{page, 25})
		if err != nil {
			return nil, err
		}

		totalPages = res.Meta.Page.TotalPages
		engines = append(engines, res.Results...)
	}

	return engines, err
}

// List engines with pagination
func (c *client) ListEngines(ctx context.Context, page Page) (data EngineResponse, err error) {
	err = c.Call(ctx, page, &data, http.MethodGet, "engines")

	return data, err
}

// List a schema by engineName
func (c *client) ListSchema(ctx context.Context, engineName string) (data SchemaDefinition, err error) {
	err = c.Call(ctx, nil, &data, http.MethodGet, "engines/%s/schema", engineName)

	if data != nil {
		data["id"] = "text"
	}

	return data, err
}

// Patch a list of documents. Every document must contain "id".
// Every document is patched separately.
// Documents without ID will be rejected.
// Non-existing documents will be rejected.
func (c *client) PatchDocuments(ctx context.Context, engineName string, documents interface{}) (res []UpdateResponse, err error) {
	err = c.Call(ctx, documents, &res, http.MethodPatch, "engines/%s/documents", engineName)

	return res, err
}

// Update (replace) a list of documents
// Every document is created (or replaced) separately.
// Documents without ID will have auto-generated ID's.
// Non-existing documents will be automatically created.
func (c *client) UpdateDocuments(ctx context.Context, engineName string, documents interface{}) (res []UpdateResponse, err error) {
	err = c.Call(ctx, documents, &res, http.MethodPost, "engines/%s/documents", engineName)

	return res, err
}

// Remove a list of documents specified as string ID's or documents with "id" field
// Every document is deleted separately.
func (c *client) RemoveDocuments(ctx context.Context, engineName string, documents interface{}) (res []DeleteResponse, err error) {
	err = c.Call(ctx, documents, &res, http.MethodDelete, "engines/%s/documents", engineName)

	return res, err
}

// Search documents by query
// TODO: query struct. Refer to https://www.elastic.co/guide/en/app-search/current/search.html#search-api-request-body
func (c *client) SearchDocuments(ctx context.Context, engineName string, query interface{}, into interface{}) (err error) {
	err = c.Call(ctx, query, into, http.MethodPost, "engines/%s/search", engineName)

	return err
}

func (c *client) Call(ctx context.Context, requestBody, resultPtr interface{}, method, urlFormat string, args ...interface{}) error {
	r, err := c.R().
		SetBody(requestBody).
		SetError(&Error{}).
		SetResult(resultPtr).
		SetContext(ctx).
		Execute(method, fmt.Sprintf(urlFormat, args...))
	if err != nil {
		return err
	}

	if r.IsError() {
		err = r.Error().(*Error)
		err.(*Error).StatusCode = r.StatusCode()
		return err
	}

	if resultPtr != nil {
		outElem := reflect.ValueOf(resultPtr).Elem()
		resultElem := reflect.ValueOf(r.Result()).Elem()

		if outElem.Type() != resultElem.Type() {
			return fmt.Errorf("cannot assign result: different types: %s != %s",
				outElem.Type().String(), resultElem.Type().String())
		}

		outElem.Set(resultElem)
	}

	return nil
}
