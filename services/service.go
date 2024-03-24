package services

import (
	"errors"

	"github.com/nuriansyah/rocket-ticket/model"
	"github.com/nuriansyah/rocket-ticket/repository"
)

type VenueService struct {
	repo *repository.Repository
}

func NewVenueService(repo *repository.Repository) *VenueService {
	return &VenueService{
		repo: repo,
	}
}

func (s *VenueService) CreateVenue(name, address string, capacity int) error {
	if name == "" {
		return errors.New("Venue name cannot be empty")
	}

	if address == "" {
		return errors.New("Venue address cannot be empty")
	}

	if capacity <= 0 {
		return errors.New("Venue capacity must be greater than zero")
	}
	venue := &model.Venue{
		Name:     name,
		Address:  address,
		Capacity: capacity,
	}
	return s.repo.CreateVenue(venue)
}

func (s *VenueService) ReadVenue(venueID string) (*model.Venue, error) {
	return s.repo.ReadVenue(venueID)
}
