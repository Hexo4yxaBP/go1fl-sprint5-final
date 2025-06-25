package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Error vars definition.

var (
	ErrUnknownTrainType = errors.New("неизвестный тип тренировки") // Unknown Training Type error
	ErrIncorrectParams  = errors.New("incorrect input parameters") // Incorrect input parameters error
	ErrWrongData        = errors.New("wrong data format")          // Wrong data format error
	ErrNegativeSteps    = errors.New("number of steps cannot be negative or zero")
	ErrNegativeDuration = errors.New("duration cannot be negative or zero")
	ErrNegativeWeight   = errors.New("weight cannot be negative or zero")
	ErrNegativeHeight   = errors.New("height cannot be negative or zero")
)

// WalkingSpentCalories calculates spent calories for walking activity based on steps count, user weight and height and activity duration.
// Returns 0 and "Incorrect input parameters" error if any parameter == 0.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	switch {
	case steps <= 0:
		return 0, ErrNegativeSteps
	case weight <= 0:
		return 0, ErrNegativeWeight
	case height <= 0:
		return 0, ErrNegativeHeight
	case duration <= 0:
		return 0, ErrNegativeDuration
	}

	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories calculates spent calories for running activity based on steps count, user weight and height and activity duration.
// Returns 0 and "Incorrect input parameters" error if any parameter == 0.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	switch {
	case steps <= 0:
		return 0, ErrNegativeSteps
	case weight <= 0:
		return 0, ErrNegativeWeight
	case height <= 0:
		return 0, ErrNegativeHeight
	case duration <= 0:
		return 0, ErrNegativeDuration
	}

	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

// MeanSpeed calculates the meanSpeed based on step count, user height and activity duration.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if duration <= 0 || steps < 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// Distance calculates the distance based on step count and user height.
func Distance(steps int, height float64) float64 {

	return height * stepLengthCoefficient * float64(steps) / mInKm
}
