package app

type Application struct {
	Server *ServerConfig
	Database *DatabaseConfig
}


func (app *Application) ApplicationSetup() {
	app.Server = &ServerConfig{}
	app.Database = &DatabaseConfig{}
	app.Server.Start()
}
