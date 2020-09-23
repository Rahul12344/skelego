package skelego

import "context"

//Service service
type Service interface {
	Configurifier(Config)
	Connect(context.Context, Config, Logging)
	Start(context.Context, Logging)
	Stop(context.Context, Logging)
}
