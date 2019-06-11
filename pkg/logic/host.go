package logic

import "github.com/wexel-nath/meat-night/pkg/database"

func GetUpcomingHosts() []string {
	return database.SelectUpcomingHosts()
}
