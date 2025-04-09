package bootstrap

type Application struct {
	Env  *Environment
	Psql *psql.Client
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.Psql = NewPsqlDatabase(app.Env)
	return *app
}

func (appPtr *Application) CloseDBConnection() {
	ClosePsqlDBConnection(appPtr.Psql)
}
