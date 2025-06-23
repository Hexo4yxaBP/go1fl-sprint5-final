package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	ErrUnknownTrainType = errors.New("неизвестный тип тренировки") // Unknown Training Type error
	ErrIncorrectParams  = errors.New("incorrect input parameters") // Incorrect input parameters error
	ErrWrongData        = errors.New("wrong data format")          // Wrong data format error
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration

	//Embedded struct
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	splittedData := strings.Split(datastring, ",")

	if len(splittedData) != 3 {
		return ErrWrongData
	}

	steps, err := strconv.Atoi(splittedData[0])

	if err != nil || steps <= 0 {
		return errors.Join(ErrWrongData, err)
	}

	activity := splittedData[1]

	if len(activity) == 0 {
		return ErrUnknownTrainType
	}

	duration, err := time.ParseDuration(splittedData[2])

	if err != nil || duration <= 0 {
		return errors.Join(ErrWrongData, err)
	}

	t.Steps, t.TrainingType, t.Duration = steps, activity, duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(t.Steps, t.Height)
	spd := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var enrgy float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		enrgy, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		enrgy, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", ErrUnknownTrainType
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, spd, enrgy), nil
}
