package services

type ServicesStore struct {
	ProjectName *ProjectNameService
	Scp         *ScpService
}

func NewServicesStore() *ServicesStore {
	return &ServicesStore{
		ProjectName: NewProjectNameService(),
		Scp:         NewScpService(),
	}
}
