package app

type Application struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

var Core *Application

func (app *Application) ApplicationSetup() {
	app.Server = &ServerConfig{}
	app.Database = &DatabaseConfig{}
	app.Database.Connect()
	app.Database.Migration()
}

func Setup() {
	appInit := Application{}
	appInit.ApplicationSetup()
	Core = &appInit
}
