package container

// ServiceContainer is a struct to handle all service used in project
type ServiceContainer struct {
}

// CreateServiceContainer is a constructor for creating all services available in the app
func CreateServiceContainer() *ServiceContainer {
	return &ServiceContainer{}
}
