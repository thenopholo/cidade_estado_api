package locations

import (
	"github.com/gin-gonic/gin"
	"github.com/thenopholo/cidade_estado_api.git/internal/app/handlers/locations/dto"
	repositories "github.com/thenopholo/cidade_estado_api.git/internal/infrastructure/repositories/location"
)

type Locationhanedler struct {
	locationRepository *repositories.LocationRepository
}

func NewLocationhanedler(locationRepository *repositories.LocationRepository) *Locationhanedler {
	return &Locationhanedler{locationRepository: locationRepository}
}

func (l *Locationhanedler) GetAllStates(c *gin.Context) {
	states, err := l.locationRepository.GetStates()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var statesResponse []dto.StateResponse
	for _, s := range states {
		statesResponse = append(statesResponse, dto.StateResponse{
			Acronym: s.Acronym,
			Nome:    s.Nome,
		})
	}
	c.JSON(200, statesResponse)
}
