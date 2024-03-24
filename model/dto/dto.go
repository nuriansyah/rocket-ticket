package dto

// VenueDTO ...

type VenueRequest struct {
	VenueID  string `json:"venue_id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Capacity int    `json:"capacity"`
}
