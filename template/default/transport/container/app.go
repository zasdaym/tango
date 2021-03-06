package container

import "github.com/payfazz/tango/template/default/transport/http/container"

// AppContainer handle all requirement for app to run properly
type AppContainer struct {
	Clients        *ClientContainer
	HttpMiddleware *container.MiddlewareContainer
	Services       *ServiceContainer
}

// CreateServiceContainer construct all requirement for app
func CreateAppContainer() *AppContainer {
	repositories := CreateRepositoryContainer()
	clients := CreateClientContainer()

	return &AppContainer{
		Clients:        clients,
		HttpMiddleware: container.CreateMiddlewareContainer(),
		Services:       CreateServiceContainer(repositories, clients),
	}
}
