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
)

// WalkingSpentCalories calculates spent calories for walking activity based on steps count, user weight and height and activity duration.
// Returns 0 and "Incorrect input parameters" error if any parameter == 0.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrIncorrectParams
	}

	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories calculates spent calories for running activity based on steps count, user weight and height and activity duration.
// Returns 0 and "Incorrect input parameters" error if any parameter == 0.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrIncorrectParams
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
