package signature

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sign(t *testing.T) {
	assert := assert.New(t)
	signature, err := Sign(nil, nil)
	assert.Nil(signature)
	assert.Nil(err)
}

func Test_Verify(t *testing.T) {
	assert := assert.New(t)
	err := Verify(nil, nil, nil)
	assert.Nil(err)
}

func Test_VerifyMultiSignature(t *testing.T) {
	assert := assert.New(t)
	err := VerifyMultiSignature(nil, nil, 0, nil)
	assert.Nil(err)
}