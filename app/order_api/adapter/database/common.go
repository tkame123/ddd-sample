package database

import "github.com/google/uuid"

func fromModelTicketID(ticketID uuid.NullUUID) *uuid.UUID {
	if ticketID.Valid {
		return &ticketID.UUID
	}
	return nil
}

func toModelTicketID(ticketID *uuid.UUID) uuid.NullUUID {
	if ticketID == nil {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{
		UUID:  *ticketID,
		Valid: true,
	}
}
