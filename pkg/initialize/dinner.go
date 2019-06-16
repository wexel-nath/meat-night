package initialize

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func MaybeInsertDinners() {
	existingDinners, err := logic.GetAllDinners()
	if err != nil {
		logger.Error(err)
		return
	}
	if len(existingDinners) > 0 {
		return
	}

	logger.Info("Initializing Meat-Night from dinners.json")
	jsonFile, err := os.Open("db/dinners.json")
	if err != nil {
		logger.Error(err)
		return
	}
	defer jsonFile.Close()

	dinnersJson, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Error(err)
		return
	}

	var dinners []model.Dinner
	err = json.Unmarshal(dinnersJson, &dinners)
	if err != nil {
		logger.Error(err)
		return
	}

	for _, dinner := range dinners {
		_, err = logic.CreateDinner(dinner)
		if err != nil {
			logger.Error(err)
			return
		}
	}
}
