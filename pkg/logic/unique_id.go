package logic

import (
	"time"

	"github.com/speps/go-hashids"
	"github.com/wexel-nath/meat-night/pkg/logger"
)

var (
	hashID *hashids.HashID
)

func Configure() {
	var err error
	hashID, err = hashids.New()
	if err != nil {
		logger.Error(err)
	}
}

func generateUniqueID(mateoID int64) (string, error) {
	n := []int64{
		mateoID,
		time.Now().UnixNano(),
	}
	return hashID.EncodeInt64(n)
}
