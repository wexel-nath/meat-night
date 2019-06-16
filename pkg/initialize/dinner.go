package initialize

import (
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func MaybeInsertDinners() {
	dinners, err := logic.GetAllDinners()
	if err != nil {
		logger.Error(err)
		return
	}
	if len(dinners) > 0 {
		return
	}

	logger.Info("Initializing Meat-Night from dinners.json")

	// read json and populate


}
