package ports

import "PrimeiroTestGo/Agendamento-Consulta/internal/domain/entities"

type Appointment_repository interface {
	Save(appoint entities.Appointment)
	Length() int
}