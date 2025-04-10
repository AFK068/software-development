package userinteraction

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
)

type UserInput interface {
	GetIntValue(message string, minValue, maxValue int) (int, error)
	GetFloatValue(message string, minValue, maxValue float64) (float64, error)
	GetStringValue(message string) (string, error)
	GetTimeValue(message string) (time.Time, error)
}

type DefaultUserInput struct{}

func NewUserInput() *DefaultUserInput {
	return &DefaultUserInput{}
}

func (d *DefaultUserInput) GetIntValue(message string, minValue, maxValue int) (int, error) {
	validate := func(input string) error {
		validate, err := strconv.Atoi(input)
		if err != nil {
			return err
		}

		if validate < minValue || validate > maxValue {
			return errors.New("value out of range")
		}

		return nil
	}

	promptValue := promptui.Prompt{
		Label:    message,
		Validate: validate,
	}

	valueStr, err := promptValue.Run()
	if err != nil {
		return 0, fmt.Errorf("prompting value: %w", err)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("converting value: %w", err)
	}

	return value, nil
}

func (d *DefaultUserInput) GetFloatValue(message string, minValue, maxValue float64) (float64, error) {
	validate := func(input string) error {
		validate, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return err
		}

		if validate < minValue || validate > maxValue {
			return errors.New("value out of range")
		}

		return nil
	}

	promptValue := promptui.Prompt{
		Label:    message,
		Validate: validate,
	}

	valueStr, err := promptValue.Run()
	if err != nil {
		return 0, fmt.Errorf("prompting value: %w", err)
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, fmt.Errorf("converting value: %w", err)
	}

	return value, nil
}

func (d *DefaultUserInput) GetStringValue(message string) (string, error) {
	promptValue := promptui.Prompt{
		Label: message,
	}

	value, err := promptValue.Run()
	if err != nil {
		return "", fmt.Errorf("prompting value: %w", err)
	}

	return value, nil
}

func (d *DefaultUserInput) GetTimeValue(message string) (time.Time, error) {
	validate := func(input string) error {
		_, err := time.Parse("2006-01-02 15:04:05", input)
		if err != nil {
			return errors.New("invalid time format, expected 'YYYY-MM-DD HH:MM:SS'")
		}

		return nil
	}

	promptValue := promptui.Prompt{
		Label:    message,
		Validate: validate,
	}

	valueStr, err := promptValue.Run()
	if err != nil {
		return time.Time{}, fmt.Errorf("prompting value: %w", err)
	}

	value, err := time.Parse("2006-01-02 15:04:05", valueStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("parsing time: %w", err)
	}

	return value, nil
}
