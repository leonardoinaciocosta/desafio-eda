package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	client, err := NewClient("Leonardo Inacio", "abc@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Leonardo Inácio", "123@uol.com")
	err := client.Update("Leonardo Inácio", "456@uol.com")
	assert.Nil(t, err)
	assert.Equal(t, "Leonardo Inácio", client.Name)
	assert.Equal(t, "123@uol.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Leonardo Inácio", "123@uol.com")
	err := client.Update("", "456@uol.com")
	assert.Error(t, err, "Name is required")
	assert.Equal(t, "Leonardo Inácio", client.Name)
	assert.Equal(t, "123@uol.com", client.Email)
}
