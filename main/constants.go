package main

//environment variables
const (
	ENV_PORT = "PORT"
	ENV_DIR_ENTRY = "DIR_ENTRY"
	ENV_DOCKER = "DOCKER"
	ENV_CANNONICAL_DOMAIN = "CANNONICAL_DOMAIN"
	ENV_STAGE = "STAGE"
)

//directories
const (
	DIR_TEMPLATES = "/templates/"
	DIR_STATIC_ASSETS_FOLDER = "static"
)

//headers
const (
	HEADER_LOAD_BALANCER_SCHEME = "X-Forwarded-Proto"
)

//values
const (
	VALUE_PRODUCTION = "production"
)
