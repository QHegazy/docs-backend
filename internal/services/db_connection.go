package services

import "docs/internal/database"

type Services struct {
	Conne *database.Connect
}

var Service = &Services{
	Conne: database.New(),
}
