package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_verify(t *testing.T) {
	_, err := New(DBName(""))
	assert.Equal(t, err, ErrDBNameEmpty, "db name is empty")
	_, err = New(DBName("test"))
	assert.Nil(t, err, "db name is empty")

}
