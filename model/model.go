package model

type Venue struct {
	VenueID  string
	Name     string
	Address  string
	Capacity int
}
type Organizer struct {
	OrganizerID string
	Name        string
	Email       string
}
type Event struct {
	EventID     string
	Name        string
	Description string
	Date        string
	VenueID     string
	OrganizerID string
	Capacity    int
	Price       int
}
type Ticket struct {
	TicketID   string
	EventID    string
	Quantity   int
	TotalPrice int
}
