package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	gqlclient "github.com/hasura/go-graphql-client"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/api"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/schema"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/testhelpers"
)

func TestMain(m *testing.M) {
	testhelpers.InitMySQL()
	m.Run()
}

type roundTripperWithAuth struct {
	handler http.Handler
}

func (r roundTripperWithAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	r.handler.ServeHTTP(w, req)
	return w.Result(), nil
}

func NewGraphQLTestClient(
	handler http.Handler,
) *gqlclient.Client {
	return gqlclient.NewClient(
		"/",
		&http.Client{
			Transport: roundTripperWithAuth{
				handler: handler,
			},
		},
	)
}

func newClient(t *testing.T, name string, seeds []interface{}) (*gqlclient.Client, func(), error) {
	rgst, clnup := testhelpers.NewRegistry(t, name, seeds)

	s, err := schema.String()
	if err != nil {
		return nil, clnup, err
	}

	cont := &api.Controller{
		Registry: rgst,
	}
	schema := graphql.MustParseSchema(s, cont)
	h := &relay.Handler{Schema: schema}
	client := gqlclient.NewClient(
		"/",
		&http.Client{
			Transport: roundTripperWithAuth{
				handler: h,
			},
		},
	)

	return client, clnup, nil
}
