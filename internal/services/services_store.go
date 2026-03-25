package services

type ServicesStore struct {
	ProjectName *ProjectNameService
}

func NewServicesStore() *ServicesStore {
	return &ServicesStore{
		ProjectName: NewProjectNameService(),
	}
}
