package skelego

import "context"

//Service service
type Service interface {
	Configurifier(Config)
	Connect(context.Context, Config)
	Start(context.Context)
	Stop(context.Context)
}
