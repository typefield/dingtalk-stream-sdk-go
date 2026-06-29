package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/handler"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/payload"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
 * @Author linya.jj
 * @Date 2023/3/22 14:23
 */

func TestNewDingtalkOpenStreamClient(t *testing.T) {
}

func TestDingtalkOpenStreamClient_Start(t *testing.T) {
}

func TestDingtalkOpenStreamClient_processDataFrame(t *testing.T) {
}

func TestDingtalkOpenStreamClient_Close(t *testing.T) {

}

func TestDingtalkOpenStreamClient_reconnect(t *testing.T) {
}

func TestDingtalkOpenStreamClient_GetHandler(t *testing.T) {
}

func TestDingtalkOpenStreamClient_CheckConfigValid(t *testing.T) {
}

func TestDingtalkOpenStreamClient_GetConnectionEndpoint(t *testing.T) {
	var requestBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, utils.GetConnectionEndpointAPIUrl, r.URL.Path)
		require.NoError(t, json.NewDecoder(r.Body).Decode(&requestBody))
		_, _ = w.Write([]byte(`{"endpoint":"wss://example.com/connect","ticket":"ticket_xxx"}`))
	}))
	defer server.Close()

	cli := NewStreamClient(
		WithAppCredential(NewAppCredentialConfig("clientId", "clientSecret")),
		WithOpenApiHost(server.URL),
		WithUserConnection(NewUserConnectionConfig("open_source", "123456", "987654")),
	)
	cli.RegisterAllEventRouter(func(ctx context.Context, df *payload.DataFrame) (*payload.DataFrameResponse, error) {
		return payload.NewSuccessDataFrameResponse(), nil
	})

	endpoint, err := cli.GetConnectionEndpoint(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "wss://example.com/connect", endpoint.Endpoint)
	assert.Equal(t, "ticket_xxx", endpoint.Ticket)
	assert.Equal(t, "clientId", requestBody["clientId"])
	assert.Equal(t, "clientSecret", requestBody["clientSecret"])
	assert.Equal(t, "open_source", requestBody["channelType"])
	assert.Equal(t, "123456", requestBody["orgId"])
	assert.Equal(t, "987654", requestBody["uid"])
	_, ok := requestBody["subscriptions"]
	assert.False(t, ok)
}

func TestDingtalkOpenStreamClient_GetConnectionEndpointWithSubscriptions(t *testing.T) {
	var requestBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, json.NewDecoder(r.Body).Decode(&requestBody))
		_, _ = w.Write([]byte(`{"endpoint":"wss://example.com/connect","ticket":"ticket_xxx"}`))
	}))
	defer server.Close()

	cli := NewStreamClient(
		WithAppCredential(NewAppCredentialConfig("clientId", "clientSecret")),
		WithOpenApiHost(server.URL),
		WithSubscription(utils.SubscriptionTypeKEvent, "topic_xxx", handler.IFrameHandler(func(ctx context.Context, df *payload.DataFrame) (*payload.DataFrameResponse, error) {
			return payload.NewSuccessDataFrameResponse(), nil
		})),
	)

	_, err := cli.GetConnectionEndpoint(context.Background())
	require.NoError(t, err)
	subscriptions, ok := requestBody["subscriptions"].([]interface{})
	require.True(t, ok)
	require.NotEmpty(t, subscriptions)
	assert.Nil(t, requestBody["channelType"])
	assert.Nil(t, requestBody["orgId"])
	assert.Nil(t, requestBody["uid"])
}

func TestDingtalkOpenStreamClient_OnDisconnect(t *testing.T) {

}

func TestDingtalkOpenStreamClient_OnPing(t *testing.T) {
}

func TestDingtalkOpenStreamClient_SendDataFrameResponse(t *testing.T) {
}

func TestDingtalkOpenStreamClient_SendErrorResponse(t *testing.T) {
}
