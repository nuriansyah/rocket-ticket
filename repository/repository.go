package repository

import (
	"database/sql"

	"github.com/nuriansyah/rocket-ticket/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) CreateVenue(venue *model.Venue) error {
	_, err := r.db.Exec("INSERT INTO venues (VenueID, Name, Address, Capacity) VALUES (?, ?, ?, ?)", venue.VenueID, venue.Name, venue.Address, venue.Capacity)
	return err
}

func (r *Repository) ReadVenue(venueID string) (*model.Venue, error) {
	var venue model.Venue
	err := r.db.QueryRow("SELECT VenueID, Name, Address, Capacity FROM venues WHERE VenueID = ?", venueID).
		Scan(&venue.VenueID, &venue.Name, &venue.Address, &venue.Capacity)
	if err != nil {
		return nil, err
	}
	return &venue, nil
}

func (r *Repository) CreateOrganizer(organizer *model.Organizer) error {
	_, err := r.db.Exec("INSERT INTO organizers (OrganizerID, Name, Email) VALUES (?, ?, ?)", organizer.OrganizerID, organizer.Name, organizer.Email)
	return err
}

func (r *Repository) ReadOrganizer(organizerID string) (*model.Organizer, error) {
	var organizer model.Organizer
	err := r.db.QueryRow("SELECT OrganizerID, Name, Email FROM organizers WHERE OrganizerID = ?", organizerID).
		Scan(&organizer.OrganizerID, &organizer.Name, &organizer.Email)
	if err != nil {
		return nil, err
	}
	return &organizer, nil
}

func (r *Repository) CreateEvent(event *model.Event) error {
	_, err := r.db.Exec("INSERT INTO events (EventID, Name, Description, Date, VenueID, OrganizerID, Capacity, Price) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", event.EventID, event.Name, event.Description, event.Date, event.VenueID, event.OrganizerID, event.Capacity, event.Price)
	return err
}

func (r *Repository) ReadEvent(eventID string) (*model.Event, error) {
	var event model.Event
	err := r.db.QueryRow("SELECT EventID, Name, Description, Date, VenueID, OrganizerID, Capacity, Price FROM events WHERE EventID = ?", eventID).
		Scan(&event.EventID, &event.Name, &event.Description, &event.Date, &event.VenueID, &event.OrganizerID, &event.Capacity, &event.Price)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *Repository) CreateTicket(ticket *model.Ticket) error {
	_, err := r.db.Exec("INSERT INTO ticket (TicketID, EventID, Quantity, TotalPrice) VALUES (?, ?, ?, ?)", ticket.TicketID, ticket.EventID, ticket.Quantity, ticket.TotalPrice)
	return err
}

func (r *Repository) ReadAllTickets() ([]*model.Ticket, error) {
	rows, err := r.db.Query("SELECT TicketID, EventID, Quantity, TotalPrice FROM ticket")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := []*model.Ticket{}
	for rows.Next() {
		var ticket model.Ticket
		err := rows.Scan(&ticket.TicketID, &ticket.EventID, &ticket.Quantity, &ticket.TotalPrice)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, &ticket)
	}
	return tickets, nil
}
