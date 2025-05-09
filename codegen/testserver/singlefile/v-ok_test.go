package singlefile

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/require"
)

func TestOk(t *testing.T) {
	resolver := &Stub{}
	resolver.QueryResolver.VOkCaseValue = func(ctx context.Context) (*VOkCaseValue, error) {
		return &VOkCaseValue{}, nil
	}
	resolver.QueryResolver.VOkCaseNil = func(ctx context.Context) (*VOkCaseNil, error) {
		return &VOkCaseNil{}, nil
	}

	srv := handler.New(NewExecutableSchema(Config{Resolvers: resolver}))
	srv.AddTransport(transport.POST{})
	c := client.New(srv)

	t.Run("v ok case value", func(t *testing.T) {
		var resp struct {
			VOkCaseValue struct {
				Value string
			}
		}
		err := c.Post(`query { vOkCaseValue { value } }`, &resp)
		require.NoError(t, err)
		require.Equal(t, "hi", resp.VOkCaseValue.Value)
	})

	t.Run("v ok case nil", func(t *testing.T) {
		var resp struct {
			VOkCaseNil struct {
				Value *string
			}
		}
		err := c.Post(`query { vOkCaseNil { value } }`, &resp)
		require.NoError(t, err)
		require.Nil(t, resp.VOkCaseNil.Value)
	})
}
