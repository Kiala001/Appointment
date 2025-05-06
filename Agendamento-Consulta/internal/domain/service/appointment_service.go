package service

import (
	"PrimeiroTestGo/Agendamento-Consulta/internal/domain/entities"
	"PrimeiroTestGo/Agendamento-Consulta/internal/domain/ports"
)

type Appointment_service struct {
	Appoint_repo ports.Appointment_repository
}

func NewAppoimentService(repo ports.Appointment_repository) *Appointment_service {
	return &Appointment_service{Appoint_repo: repo}
}

func (ap *Appointment_service) Schedule(appoint entities.Appointment) {
	ap.Appoint_repo.Save(appoint)
}