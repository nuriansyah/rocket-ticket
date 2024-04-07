package services

import (
	"errors"
	"testing"
)

type MockVenueRepository struct{}

func (m MockVenueRepository) CreateVenue(name, address string, capacity int) error {
	if name == "" {
		return errors.New("Venue name cannot be empty")
	}
	if address == "" {
		return errors.New("Venue address cannot be empty")
	}
	if capacity <= 0 {
		return errors.New("Venue capacity must be greater than zero")
	}
	return nil
}

type MockVenueService struct {
	repo *MockVenueRepository
}

func NewMockVenueService(repo *MockVenueRepository) MockVenueService {
	return MockVenueService{repo: repo}
}

func (s MockVenueService) CreateVenue(name, address string, capacity int) error {
	if name == "" {
		return errors.New("Venue name cannot be empty")
	}
	if address == "" {
		return errors.New("Venue address cannot be empty")
	}
	if capacity <= 0 {
		return errors.New("Venue capacity must be greater than zero")
	}
	return s.repo.CreateVenue(name, address, capacity)
}

func TestCreateVenue(t *testing.T) {
	repo := MockVenueRepository{}
	service := NewMockVenueService(&repo)

	tests := []struct {
		name         string
		address      string
		capacity     int
		expectedErr  error
		expectedRepo bool
	}{
		{
			name:         "ValidVenue",
			address:      "123 Main St",
			capacity:     100,
			expectedErr:  nil,
			expectedRepo: true,
		},
		{
			name:         "",
			address:      "456 Elm St",
			capacity:     50,
			expectedErr:  errors.New("Venue name cannot be empty"),
			expectedRepo: false,
		},
		{
			name:         "EmptyAddress",
			address:      "",
			capacity:     50,
			expectedErr:  errors.New("Venue address cannot be empty"),
			expectedRepo: false,
		},
		{
			name:         "ZeroCapacity",
			address:      "789 Oak St",
			capacity:     0,
			expectedErr:  errors.New("Venue capacity must be greater than zero"),
			expectedRepo: false,
		},
		{
			name:         "NegativeCapacity",
			address:      "456 Pine St",
			capacity:     -10,
			expectedErr:  errors.New("Venue capacity must be greater than zero"),
			expectedRepo: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := service.CreateVenue(test.name, test.address, test.capacity)
			if (err != nil && test.expectedErr == nil) || (err == nil && test.expectedErr != nil) || (err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error()) {
				t.Errorf("Got error '%v', but expected '%v'", err, test.expectedErr)
			}
		})
	}
}
