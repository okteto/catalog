module github.com/okteto/catalog/api

go 1.16

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/justinas/alice v1.2.0
	github.com/okteto/catalog/health v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.22.0
)

replace github.com/okteto/catalog/health => ../health
