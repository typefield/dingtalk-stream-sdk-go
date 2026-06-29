package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * @Author linya.jj
 * @Date 2023/3/22 14:50
 */

func TestAppCredentialConfig_Valid(t *testing.T) {
	conf := NewAppCredentialConfig("clientId", "clientSecret")
	assert.Nil(t, conf.Valid())

	conf.ClientId = ""
	assert.NotNil(t, conf.Valid())

	conf = nil
	assert.NotNil(t, conf.Valid())
}

func TestDingtalkGoSDKUserAgent_Valid(t *testing.T) {
	conf := NewDingtalkGoSDKUserAgent()
	assert.Nil(t, conf.Valid())

	conf.UserAgent = ""
	assert.NotNil(t, conf.Valid())

	conf = nil
	assert.NotNil(t, conf.Valid())
}

func TestUserConnectionConfig_Valid(t *testing.T) {
	conf := NewUserConnectionConfig("open_source", "123456", "987654")
	assert.Nil(t, conf.Valid())

	conf.ChannelType = ""
	assert.NotNil(t, conf.Valid())

	conf = nil
	assert.Nil(t, conf.Valid())
}
