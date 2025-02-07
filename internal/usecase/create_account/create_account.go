package createaccount

import (
	"github.com/davifs92/fc-ms-wallet/internal/entity"
	"github.com/davifs92/fc-ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientId string

}

type CreateAccountOutputDTO struct {
	Id string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway: c,
	}

}


func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error){
	client, err  := uc.ClientGateway.Get(input.ClientId)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		Id: account.Id,
	}, nil
}