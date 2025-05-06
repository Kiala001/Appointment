package tests

import (
	"PrimeiroTestGo/Agendamento-Consulta/internal/domain/entities"
	"PrimeiroTestGo/Agendamento-Consulta/internal/domain/service"
	inmemory "PrimeiroTestGo/Agendamento-Consulta/internal/infra/adapters/in_memory"
	"testing"
)

func TestDeveAdicionarConsultaAoRepositorio(t *testing.T){

	appointment_repository := inmemory.NewAppointmentRepositoryInMemeory()
	appointment := entities.Appointment{
		Id: "001",
		Patient_id: "PT-99484",
		Medic_id: "MD-23445",
		Creation_date: "12/02/2024",
		Creation_time: "12:45",
		Status: "Agendada",
	}

	appointment_service := service.NewAppoimentService(appointment_repository)
	appointment_service.Schedule(appointment)

	expected := 1
	length_repo := appointment_repository.Length()

	if length_repo != expected {
		t.Errorf("Esperava %d, mas obteve %d",length_repo, expected)
	}
}