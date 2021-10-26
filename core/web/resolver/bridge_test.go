package resolver

import (
	"database/sql"
	"net/url"
	"testing"

	"github.com/graph-gophers/graphql-go/gqltesting"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/bridges"
	bridgeORMMocks "github.com/smartcontractkit/chainlink/core/bridges/mocks"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

func Test_Bridges(t *testing.T) {
	t.Parallel()

	var (
		f         = setupFramework(t)
		bridgeORM = &bridgeORMMocks.ORM{}
	)

	bridgeURL, err := url.Parse("https://external.adapter")
	require.NoError(t, err)

	t.Cleanup(func() {
		mock.AssertExpectationsForObjects(t,
			bridgeORM,
		)
	})

	f.App.On("BridgeORM").Return(bridgeORM)
	bridgeORM.On("BridgeTypes", PageDefaultOffset, PageDefaultLimit).Return([]bridges.BridgeType{
		{
			Name:                   "bridge1",
			URL:                    models.WebURL(*bridgeURL),
			Confirmations:          uint32(1),
			OutgoingToken:          "outgoingToken",
			MinimumContractPayment: assets.NewLinkFromJuels(1),
			CreatedAt:              f.Timestamp(),
		},
	}, 1, nil)

	gqltesting.RunTest(t, &gqltesting.Test{
		Context: f.Ctx,
		Schema:  f.RootSchema,
		Query: `
			{
				bridges {
					results {
						name
						url
						confirmations
						outgoingToken
						minimumContractPayment
						createdAt
					}
					metadata {
						total
					}
				}
			}
		`,
		ExpectedResult: `
			{
				"bridges": {
					"results": [{
						"name": "bridge1",
						"url": "https://external.adapter",
						"confirmations": 1,
						"outgoingToken": "outgoingToken",
						"minimumContractPayment": "1",
						"createdAt": "2021-01-01T00:00:00Z"
					}],
					"metadata": {
						"total": 1
					}
				}
			}
		`,
	})
}

func Test_Bridge(t *testing.T) {
	var (
		query = `{
			bridge(name: "bridge1") {
				... on Bridge {
					name
					url
					confirmations
					outgoingToken
					minimumContractPayment
					createdAt
				}
				... on NotFoundError {
					message
					code
				}
			}
		}`

		name = bridges.TaskType("bridge1")
	)
	bridgeURL, err := url.Parse("https://external.adapter")
	require.NoError(t, err)

	testCases := []struct {
		name   string
		before func(*gqlTestFramework, *bridgeORMMocks.ORM)
		query  string
		result string
	}{
		{
			name: "success",
			before: func(f *gqlTestFramework, bridgeORM *bridgeORMMocks.ORM) {
				f.App.On("BridgeORM").Return(bridgeORM)
				bridgeORM.On("FindBridge", name).Return(bridges.BridgeType{
					Name:                   name,
					URL:                    models.WebURL(*bridgeURL),
					Confirmations:          uint32(1),
					OutgoingToken:          "outgoingToken",
					MinimumContractPayment: assets.NewLinkFromJuels(1),
					CreatedAt:              f.Timestamp(),
				}, nil)
			},
			query: query,
			result: `{
				"bridge": {
					"name": "bridge1",
					"url": "https://external.adapter",
					"confirmations": 1,
					"outgoingToken": "outgoingToken",
					"minimumContractPayment": "1",
					"createdAt": "2021-01-01T00:00:00Z"
				}
			}`,
		},
		{
			name: "not found",
			before: func(f *gqlTestFramework, bridgeORM *bridgeORMMocks.ORM) {
				f.App.On("BridgeORM").Return(bridgeORM)
				bridgeORM.On("FindBridge", name).Return(bridges.BridgeType{}, sql.ErrNoRows)
			},
			query: query,
			result: `{
				"bridge": {
					"message": "bridge not found",
					"code": "NOT_FOUND"
				}
			}`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var (
				f         = setupFramework(t)
				bridgeORM = &bridgeORMMocks.ORM{}
			)

			t.Cleanup(func() {
				mock.AssertExpectationsForObjects(t,
					bridgeORM,
				)
			})

			if tc.before != nil {
				tc.before(f, bridgeORM)
			}

			gqltesting.RunTest(t, &gqltesting.Test{
				Context:        f.Ctx,
				Schema:         f.RootSchema,
				Query:          tc.query,
				ExpectedResult: tc.result,
			})
		})
	}
}

func Test_CreateBridge(t *testing.T) {
	t.Parallel()

	var (
		f         = setupFramework(t)
		bridgeORM = &bridgeORMMocks.ORM{}

		name = bridges.TaskType("bridge1")
	)

	bridgeURL, err := url.Parse("https://external.adapter")
	require.NoError(t, err)

	t.Cleanup(func() {
		mock.AssertExpectationsForObjects(t,
			bridgeORM,
		)
	})

	f.App.On("BridgeORM").Return(bridgeORM)
	bridgeORM.On("FindBridge", name).Return(bridges.BridgeType{}, sql.ErrNoRows)
	bridgeORM.On("CreateBridgeType", mock.IsType(&bridges.BridgeType{})).
		Run(func(args mock.Arguments) {
			arg := args.Get(0).(*bridges.BridgeType)
			*arg = bridges.BridgeType{
				Name:                   name,
				URL:                    models.WebURL(*bridgeURL),
				Confirmations:          uint32(1),
				OutgoingToken:          "outgoingToken",
				MinimumContractPayment: assets.NewLinkFromJuels(1),
				CreatedAt:              f.Timestamp(),
			}
		}).
		Return(nil)

	gqltesting.RunTest(t, &gqltesting.Test{
		Context: f.Ctx,
		Schema:  f.RootSchema,
		Query: `
			mutation createBridge($input: CreateBridgeInput!) {
				createBridge(input: $input) {
					... on CreateBridgeSuccess {
						bridge {
							name
							url
							confirmations
							outgoingToken
							minimumContractPayment
							createdAt
						}
					}
				}
			}
		`,
		Variables: map[string]interface{}{
			"input": map[string]interface{}{
				"name":                   "bridge1",
				"url":                    "https://external.adapter",
				"confirmations":          1,
				"minimumContractPayment": "1",
			},
		},
		// We need to test equality for the generated token but since it is
		// generated by a non mockable object, we can't do this right now.
		ExpectedResult: `
			{
				"createBridge": {
					"bridge": {
						"name": "bridge1",
						"url": "https://external.adapter",
						"confirmations": 1,
						"outgoingToken": "outgoingToken",
						"minimumContractPayment": "1",
						"createdAt": "2021-01-01T00:00:00Z"
					}
				}
			}
		`,
	})
}

func Test_UpdateBridge(t *testing.T) {
	var (
		name  = bridges.TaskType("bridge1")
		query = `
		mutation updateBridge($input: UpdateBridgeInput!) {
			updateBridge(name: "bridge1", input: $input) {
				... on UpdateBridgeSuccess {
					bridge {
						name
						url
						confirmations
						outgoingToken
						minimumContractPayment
						createdAt
					}
				}
				... on NotFoundError {
					message
					code
				}
			}
		}`
		variables = map[string]interface{}{
			"input": map[string]interface{}{
				"name":                   "bridge-updated",
				"url":                    "https://external.adapter.new",
				"confirmations":          2,
				"minimumContractPayment": "2",
			},
		}
	)
	bridgeURL, err := url.Parse("https://external.adapter")
	require.NoError(t, err)

	newBridgeURL, err := url.Parse("https://external.adapter.new")
	require.NoError(t, err)

	testCases := []struct {
		name      string
		before    func(*gqlTestFramework, *bridgeORMMocks.ORM)
		query     string
		variables map[string]interface{}
		result    string
	}{
		{
			name: "success",
			before: func(f *gqlTestFramework, bridgeORM *bridgeORMMocks.ORM) {
				// Initialize the existing bridge
				bridge := bridges.BridgeType{
					Name:                   name,
					URL:                    models.WebURL(*bridgeURL),
					Confirmations:          uint32(1),
					OutgoingToken:          "outgoingToken",
					MinimumContractPayment: assets.NewLinkFromJuels(1),
					CreatedAt:              f.Timestamp(),
				}

				f.App.On("BridgeORM").Return(bridgeORM)
				bridgeORM.On("FindBridge", name).Return(bridge, nil)

				btr := &bridges.BridgeTypeRequest{
					Name:                   bridges.TaskType("bridge-updated"),
					URL:                    models.WebURL(*newBridgeURL),
					Confirmations:          2,
					MinimumContractPayment: assets.NewLinkFromJuels(2),
				}

				bridgeORM.On("UpdateBridgeType", mock.IsType(&bridges.BridgeType{}), btr).
					Run(func(args mock.Arguments) {
						arg := args.Get(0).(*bridges.BridgeType)
						*arg = bridges.BridgeType{
							Name:                   "bridge-updated",
							URL:                    models.WebURL(*newBridgeURL),
							Confirmations:          2,
							OutgoingToken:          "outgoingToken",
							MinimumContractPayment: assets.NewLinkFromJuels(2),
							CreatedAt:              f.Timestamp(),
						}
					}).
					Return(nil)
			},
			query:     query,
			variables: variables,
			result: `{
				"updateBridge": {
					"bridge": {
						"name": "bridge-updated",
						"url": "https://external.adapter.new",
						"confirmations": 2,
						"outgoingToken": "outgoingToken",
						"minimumContractPayment": "2",
						"createdAt": "2021-01-01T00:00:00Z"
					}
				}
			}`,
		},
		{
			name: "not found",
			before: func(f *gqlTestFramework, bridgeORM *bridgeORMMocks.ORM) {
				f.App.On("BridgeORM").Return(bridgeORM)
				bridgeORM.On("FindBridge", name).Return(bridges.BridgeType{}, sql.ErrNoRows)
			},
			query:     query,
			variables: variables,
			result: `{
				"updateBridge": {
					"message": "bridge not found",
					"code": "NOT_FOUND"
				}
			}`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var (
				f         = setupFramework(t)
				bridgeORM = &bridgeORMMocks.ORM{}
			)

			t.Cleanup(func() {
				mock.AssertExpectationsForObjects(t,
					bridgeORM,
				)
			})

			if tc.before != nil {
				tc.before(f, bridgeORM)
			}

			gqltesting.RunTest(t, &gqltesting.Test{
				Context:        f.Ctx,
				Schema:         f.RootSchema,
				Query:          tc.query,
				Variables:      tc.variables,
				ExpectedResult: tc.result,
			})
		})
	}
}
