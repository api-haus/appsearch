package appsearch

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	c, err := Open(os.Getenv("APPSEARCH_ENDPOINT"), os.Getenv("APPSEARCH_KEY"))
	require.NoError(t, err)

	t.Run("ListAllEngines", func(t *testing.T) {
		engines, err := c.ListAllEngines(context.TODO())
		require.NoError(t, err)
		require.NotEmpty(t, engines)

		fmt.Printf("%v", engines)
	})

	t.Run("ListSchema", func(t *testing.T) {
		engines, err := c.ListAllEngines(context.TODO())
		require.NoError(t, err)

		if len(engines) == 0 {
			return
		}

		schema, err := c.ListSchema(context.TODO(), engines[0].Name)
		require.NoError(t, err)
		require.NotEmpty(t, schema)

		fmt.Printf("%v", schema)
	})

	t.Run("Must implement APIClient", func(t *testing.T) {
		var c APIClient
		c = &mock{}
		c = &client{}
		_ = c
	})
}
