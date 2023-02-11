package model_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestClinicName_Validate(t *testing.T) {
	type (
		give struct {
			name string
		}
		want struct {
			err error
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "no_error",
			give: give{
				name: "鈴木歯科医院",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error_name_is_empty",
			give: give{
				name: "",
			},
			want: want{
				err: errors.New("clinic: name is empty"),
			},
		},
		{
			name: "name_is_just_50",
			give: give{
				name: strings.Repeat("あ", 50),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error_name_is_over_50",
			give: give{
				name: strings.Repeat("あ", 51),
			},
			want: want{
				err: errors.New("clinic: name exceeds 50 characters"),
			},
		},
	}

	for i := range tests {
		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			name := model.ClinicName(test.give.name)
			err := name.Validate()
			if test.want.err != nil {
				assert.EqualError(t, test.want.err, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestClinicNameKana_Validate(t *testing.T) {
	type (
		give struct {
			name string
		}
		want struct {
			err error
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "no_error",
			give: give{
				name: "スズキシカイイン",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error_name_is_empty",
			give: give{
				name: "",
			},
			want: want{
				err: errors.New("clinic: name_kana is empty"),
			},
		},
		{
			name: "name_is_just_100",
			give: give{
				name: strings.Repeat("あ", 100),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error_name_is_over_100",
			give: give{
				name: strings.Repeat("あ", 101),
			},
			want: want{
				err: errors.New("clinic: name_kana exceeds 100 characters"),
			},
		},
	}

	for i := range tests {
		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			kana := model.ClinicNameKana(test.give.name)
			err := kana.Validate()
			if test.want.err != nil {
				assert.EqualError(t, test.want.err, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
