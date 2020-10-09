package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

)

func Test_program(t *testing.T){

	assertions := assert.New(t)

	var ad (map[int]string)
	assertions.Nil(ad)
}
