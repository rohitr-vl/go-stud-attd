package bootstrap

import "clean-architecture/psql"

type Application struct {
	Env  *Environment
	Psql *psql.PsqlRepository
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.Psql = NewPsqlDatabase(app.Env)
	return app
}
