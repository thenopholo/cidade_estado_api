package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/thenopholo/cidade_estado_api.git/internal/domain/entities"
	"github.com/thenopholo/cidade_estado_api.git/internal/infrastructure/repositories/location/dto"
	"net/http"
	"time"
)

type LocationRepository struct {
}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{}
}

func (l *LocationRepository) GetStates() ([]entities.StateEntity, error) {
	httpClient := http.Client{
		Timeout: 20 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/ibge/uf/v1", nil)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao buscar estados: ", err)
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Erro ao fechar o Body de busca de estados %v", err)
		}
	}()

	var statesResponse []dto.BrazilApiResponse

	err = json.NewDecoder(resp.Body).Decode(&statesResponse)
	if err != nil {
		fmt.Println("Erro ao decodificar estados: ", err)
		return nil, err
	}

	var states []entities.StateEntity

	for _, s := range statesResponse {
		states = append(states, entities.StateEntity{
			Acronym: s.Sigla,
			Name:    s.Name,
		})
	}

	return states, nil

}
