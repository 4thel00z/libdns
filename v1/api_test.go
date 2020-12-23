package v1

import (
	"fmt"
	"github.com/4thel00z/libdns/v1/pkg/libdns/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleQueryOnce(t *testing.T) {
	response, err := SimpleQueryOnce(utils.CloudflarePrimary, "google.de", utils.A, utils.InternetClass, 10)
	assert.Nil(t, err)
	assert.Greater(t, response.ANCount, uint16(0))
	assert.Greater(t, response.QDCount, uint16(0))
	fmt.Printf("%v", response)
}
