package createclient

import (
	"time"

	"github.com/davifs92/fc-ms-wallet/internal/entity"
	"github.com/davifs92/fc-ms-wallet/internal/gateway"
)

type CreateClientInputDTO struct {
	Name string
	Email string
}

type CreateClientOutputDTO struct {
	Id string
	Name string
	Email string
	CreatedAt time.Time
	UpdateAt time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}

}

func (uc *CreateClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) { 
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = uc.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		Id: client.Id,
		Name: client.Name,
		Email: client.Email,
		CreatedAt: client.CreatedAt,
		UpdateAt: client.UpdatedAt,
	}, nil

}