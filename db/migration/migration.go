package migration

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Migrate(db *sql.DB) {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS venues (
            VenueID VARCHAR(36) PRIMARY KEY,
            Name VARCHAR(255) NOT NULL,
            Address VARCHAR(255) NOT NULL,
            Capacity INT
        );
        CREATE TABLE IF NOT EXISTS organizers (
            OrganizerID VARCHAR(36) PRIMARY KEY,
            Name VARCHAR(255) NOT NULL,
            Email VARCHAR(255) NOT NULL
        );
        CREATE TABLE IF NOT EXISTS events (
            EventID VARCHAR(36),
            Name VARCHAR(255) NOT NULL,
            Description TEXT,
            Date DATETIME  NOT NULL,
            VenueID VARCHAR(36) NOT NULL,
            OrganizerID VARCHAR(36) NOT NULL,
            Capacity INT,
            Price INT NOT NULL,
            FOREIGN KEY (VenueID) REFERENCES Venues(VenueID),
            FOREIGN KEY (OrganizerID) REFERENCES Organizers(OrganizerID)
        );
        CREATE TABLE IF NOT EXISTS ticket (
            TicketID VARCHAR(36) PRIMARY KEY,
            EventID VARCHAR(36) NOT NULL,
            Quantity INT NOT NULL,
            TotalPrice INT NOT NULL,
            FOREIGN KEY (EventID) REFERENCES Events(EventID)
        );
    `)

	if err != nil {
		panic(err)
	}
}
