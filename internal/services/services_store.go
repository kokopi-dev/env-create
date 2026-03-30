package services

type ServicesStore struct {
	ProjectName *ProjectNameService
	Scp         *ScpService
	Config      *ConfigService
}

func NewServicesStore() *ServicesStore {
	cfg, _ := NewConfigService()
	return &ServicesStore{
		ProjectName: NewProjectNameService(),
		Scp:         NewScpService(),
		Config:      cfg,
	}
}
