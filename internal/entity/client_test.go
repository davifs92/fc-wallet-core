package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err:= NewClient("John Doe", "j@j.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)

}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.Nil(t, client)
	assert.NotNil(t, err)

}

func TestUpdateClient(t *testing.T){
	client, _ := NewClient("Jonh Doe", "j@j.com")
	err := client.Update("John Doe update", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe update", client.Name)
	assert.Equal(t, "j@j.com", client.Email)

}

func TestUpdateClientWithInvalidData(t *testing.T){
	client, _ := NewClient("Jonh Doe", "j@j.com")
	err := client.Update("", "j@j.com")
	assert.Error(t, err)

}

func TestAddAccountToClient(t *testing.T){
	client, _ := NewClient("Jonh Doe", "j@j.com")
	account := NewAccount(client)

	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}