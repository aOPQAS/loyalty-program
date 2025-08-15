package deps

import (
	"microservice/internal/pgsql"
	"microservice/pkg/telebon"
)

type Deps struct {
	PG      *pgsql.Client
	Telebon *telebon.TelebonClient
}
