package command

import "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/database"

func RunDatabaseMigrations() error {
	return database.GetDatabase().RunMigrations()
}
