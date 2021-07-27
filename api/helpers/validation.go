package helpers

import (
	"errors"
	"regexp"
	"time"
)

func ValidatePhone(phone string) error {
	re := regexp.MustCompile(`^[+][9][9][8]\d{9}$`)
	if !re.MatchString(phone) {
		return errors.New("invalid phone number")
	}

	return nil
}

func ValidateDates(startDate, endDate string) error {
	var layout = "2006-01-02 15:04:05"

	from, err := time.Parse(layout, startDate)
	if err != nil {
		return errors.New("start_date is invalid")
	}

	to, err := time.Parse(layout, endDate)
	if err != nil {
		return errors.New("end_time is invalid")
	}

	if !from.Before(to) {
		return errors.New("start_time can not be greater than end_time")
	}

	return nil
}
