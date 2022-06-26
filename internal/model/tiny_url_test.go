package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTinyUrl(t *testing.T) {
	if err := InitDB("xjk:An0thrS3crt@(localhost:3307)/service"); err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	if err := CreateTinyUrl("https://baidu.com/absfd1", "jhA32F"); err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestRetrieveByShotCode(t *testing.T) {
	if err := InitDB("xjk:An0thrS3crt@(localhost:3307)/service"); err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	res, err := RetrieveByShotCode("jh832F")
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, res, "retrieve nothing")
	assert.Equal(t, "https://baidu.com/absc", res.FullUrl, "retrieve url isn't original url")
}
