package inmemory

import "Enrollment/internal/domain/entities"

type InMemoryAppointRepository struct {
	consults map[string]entities.Appointment
}

func NewAppointmentRepositoryInMemeory() *InMemoryAppointRepository {
	return &InMemoryAppointRepository{consults: make(map[string]entities.Appointment)}
}

func (im *InMemoryAppointRepository) Save(appoint entities.Appointment){
	im.consults[appoint.Id] = appoint
}

func (im *InMemoryAppointRepository) Length() int{
	return len(im.consults)
}