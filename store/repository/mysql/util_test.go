package mysql

import (
	"api-gmr/store/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

type filterTest struct {
	year   int
	month  int
	userID int
	status string
}

func (f filterTest) GetYear() int {
	return f.year
}

func (f filterTest) GetMonth() int {
	return f.month
}

func (f filterTest) GetUserID() int {
	return f.userID
}

func (f filterTest) GetStatus() string {
	return f.status
}

func Test_buildBillingFilter(t *testing.T) {
	type args struct {
		filter repository.BillingFilter
	}

	type want struct {
		param []string
		value []interface{}
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "test filter #1",
			args: args{filterTest{2021, 1, 1, "4"}},
			want: want{
				param: []string{"year = ?", "month = ?", "user_id = ?", "status = ?"},
				value: []interface{}{"2021", "1", "1", "4"},
			},
		}, {
			name: "test filter #2",
			args: args{filterTest{0, 0, 1, "4"}},
			want: want{
				param: []string{"user_id = ?", "status = ?"},
				value: []interface{}{"1", "4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, args := buildBillingFilter(tt.args.filter)
			assert.Equal(t, tt.want.param, v)
			assert.Equal(t, tt.want.value, args)
		})
	}
}
