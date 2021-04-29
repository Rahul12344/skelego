package skelego

import (
	"context"
)

//Runner Take in services and configures, connects, starts, and stops them
type Runner struct {
	configFile string
}

//NewRunner New instance of a runner with configfile
func NewRunner(configFile string) *Runner {
	return &Runner{
		configFile: configFile,
	}
}

//Configure sets up configurations
func (r *Runner) Configure(conf Config) {

}

//Connect connects to services
func (r *Runner) Connect(ctx context.Context, config Config, services ...Service) {
	for _, service := range services {
		service.Connect(ctx, config)
	}
}

//Start starts up services that have been added
func (r *Runner) Start(ctx context.Context, config Config, services ...Service) {
	for _, service := range services {
		service.Start(ctx)
	}
}

//Stop shuts down services
func (r *Runner) Stop(ctx context.Context, config Config, services ...Service) {
	for _, service := range services {
		service.Stop(ctx)
	}
}
