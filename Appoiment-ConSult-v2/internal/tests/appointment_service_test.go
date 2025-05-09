package tests

import (
	inmemory "Enrollment/internal/adapters/in_memory"
	"Enrollment/internal/application/events"
	"Enrollment/internal/domain/entities"
	"Enrollment/internal/domain/service"
	"testing"

	"github.com/stretchr/testify/mock"
)


func TestDeveAdicionarConsultaAoRepositorio(t *testing.T){

	event := events.NewEventBus()
	event.On("Publish", mock.Anything, mock.Anything).Return()
	appointment_repository := inmemory.NewAppointmentRepositoryInMemeory()

	appointment := entities.Appointment{
		Id: "001",
		Patient_id: "PT-99484",
		Medic_id: "MD-23445",
		Creation_date: "12/02/2024",
		Creation_time: "12:45",
		Status: "Agendada",
	}

	appointment_service := service.NewAppoimentService(appointment_repository, event)
	appointment_service.Schedule(appointment)

	expected := 1
	length_repo := appointment_repository.Length()

	if length_repo != expected {
		t.Errorf("Esperava %d, mas obteve %d",length_repo, expected)
	}
}

func TestDeveDispararConsultaAgendada(t *testing.T){
	event := events.NewEventBus()

	event.On("Publish", mock.Anything, mock.Anything).Return()
	appointment_repository := inmemory.NewAppointmentRepositoryInMemeory()
	appointment := entities.Appointment{
		Id: "001",
		Patient_id: "PT-99484",
		Medic_id: "MD-23445",
		Creation_date: "12/02/2024",
		Creation_time: "12:45",
		Status: "Agendada",
	}

	appointment_service := service.NewAppoimentService(appointment_repository, event)
	appointment_service.Schedule(appointment)

	event.AssertCalled(t, "Publish", "ConsultaAgendada")
}