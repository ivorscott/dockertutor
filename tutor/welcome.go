package tutor

func dockerIntro() string {
	return `
===============================================================================
                        Welcome to the Docker CLI Tutor
===============================================================================
Docker is an open platform for developing, shipping, and running applications.
This tutor aims to be an interactive tutorial for learning Docker.
Docker enables you to separate your applications from your infrastructure so
you can deliver software quickly.
===============================================================================
`
}

func composeIntro() string {
	return `
===============================================================================
                    Welcome to the Docker Compose Tutor
===============================================================================
Compose is a tool for defining and running multi-container Docker applications.
This tutor aims to be an interactive tutorial for learning Docker Compose. With
Compose, you use a YAML file to configure your application’s services. Then,
with a single command, you create and start all the services from your
configuration.
===============================================================================
`
}

func swarmIntro() string {
	return `
===============================================================================
                    Welcome to the Docker Swarm mode Tutor 
===============================================================================
The cluster management and orchestration features embedded in the Docker Engine
are built using swarmkit. Swarmkit is a separate project which implements
Docker’s orchestration layer and is used directly within Docker. This tutor
aims to be an interactive tutorial for learning Docker Swarm mode. When Docker
is running in swarm mode, you can still run standalone containers on any of the
Docker hosts participating in the swarm, as well as swarm services.
===============================================================================
`
}

func lbreak() string {
	return "==============================================================================="
}
