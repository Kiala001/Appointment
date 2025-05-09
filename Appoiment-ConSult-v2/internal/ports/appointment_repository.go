package ports

import "Enrollment/internal/domain/entities"

type Appointment_repository interface {
	Save(appoint entities.Appointment)
	Length() int
}