package model

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"
)

const (
	maxClinicNameChars     int = 50
	maxClinicNameKanaChars int = 100
)

type (
	Clinic struct {
		ID       ID
		Name     ClinicName
		NameKana ClinicNameKana
		Contact  PhoneNumber
	}
	ClinicName     string
	ClinicNameKana string
)

func RegisterClinic(
	name ClinicName,
	nameKana ClinicNameKana,
) (Clinic, error) {
	id, err := NewID("clinic")
	if err != nil {
		return Clinic{}, err
	}

	c := Clinic{
		ID:       id,
		Name:     name,
		NameKana: nameKana,
	}

	if err := c.Validate(); err != nil {
		return Clinic{}, err
	}

	return c, nil
}

func GetClinic(
	id ID,
	name ClinicName,
	nameKana ClinicNameKana,
) (Clinic, error) {
	c := Clinic{
		ID:       id,
		Name:     name,
		NameKana: nameKana,
	}

	if err := c.Validate(); err != nil {
		return Clinic{}, err
	}

	return c, nil
}

func (c Clinic) Validate() error {
	if err := c.Name.Validate(); err != nil {
		return err
	}

	if err := c.NameKana.Validate(); err != nil {
		return err
	}
	return nil
}

func (c ClinicName) Validate() error {
	if c == "" {
		return apperror.New(errors.New("clinic: name is empty"))
	}

	if utf8.RuneCountInString(string(c)) > maxClinicNameChars {
		err := fmt.Errorf("clinic: name exceeds %d characters", maxClinicNameChars)
		return apperror.New(err)
	}

	return nil
}

func (c ClinicNameKana) Validate() error {
	if c == "" {
		return apperror.New(errors.New("clinic: name_kana is empty"))
	}

	if utf8.RuneCountInString(string(c)) > maxClinicNameKanaChars {
		err := fmt.Errorf("clinic: name_kana exceeds %d characters", maxClinicNameKanaChars)
		return apperror.New(err)
	}
	return nil
}
