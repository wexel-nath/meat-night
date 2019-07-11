package logic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wexel-nath/meat-night/pkg/config"
)

func init() {
	config.Configure()
}

func TestGetNextDinnerTime(t *testing.T) {
	tests := map[string]struct {
		now  time.Time
		want time.Time
	}{
		"Tuesday": {
			now:  time.Date(2019, 7, 9, 11, 9, 29, 8, time.Local),
			want: time.Date(2019, 7, 10, 19, 0, 0, 0, time.Local),
		},
		"Wednesday": {
			now:  time.Date(2019, 7, 10, 20, 9, 29, 8, time.Local),
			want: time.Date(2019, 7, 17, 19, 0, 0, 0, time.Local),
		},
		"Thursday": {
			now:  time.Date(2019, 7, 11, 11, 9, 29, 8, time.Local),
			want: time.Date(2019, 7, 17, 19, 0, 0, 0, time.Local),
		},
	}

	for name, test := range tests {
		t.Run(name, func(st *testing.T) {
			got := getNextDinnerTime(test.now)
			assert.Equal(st, test.want, got)
		})
	}
}
