package service

import (
	"Enrollment/internal/domain/entities"
	"Enrollment/internal/ports"
)

type Appointment_service struct {
	event ports.Event
	Appoint_repo ports.Appointment_repository
}

func NewAppoimentService(repo ports.Appointment_repository, event ports.Event) *Appointment_service {
	return &Appointment_service{Appoint_repo: repo, event: event}
}

func (ap *Appointment_service) Schedule(appoint entities.Appointment) {
	ap.Appoint_repo.Save(appoint)
	ap.event.Publish("ConsultaAgendada", appoint)
}