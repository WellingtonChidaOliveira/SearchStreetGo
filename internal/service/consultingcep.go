package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/wellingtonchida/searchstreet/types"
)

type ConsultingCep interface {
	GetCep(ctx context.Context, zipCode string) (*types.Address, error)
}

func (c *CepFetcher) GetCep(ctx context.Context, zipCode string) (*types.Address, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("cep not found: " + zipCode)
	}
	address, err := jsonEncoder(resp.Body)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func jsonEncoder(address io.Reader) (*types.Address, error) {
	var cep types.Address
	err := json.NewDecoder(address).Decode(&cep)
	if err != nil {
		return nil, err
	}

	return &cep, nil
}
