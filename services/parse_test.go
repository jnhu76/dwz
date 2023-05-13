package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUrl(t *testing.T) {
	url1 := "https://www.baidu.com"
	url2 := "httt://www.baidu.com"
	url3 := "dafdasf"

	assert.True(t, ParseUrl(url1))
	assert.False(t, ParseUrl(url2))
	assert.False(t, ParseUrl(url3))
}
