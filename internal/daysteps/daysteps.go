package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Error vars definition.
var (
	ErrWrongData        = errors.New("wrong data") // Wrong data format error
	ErrNegativeSteps    = errors.New("number of steps cannot be negative or zero")
	ErrNegativeDuration = errors.New("duration cannot be negative or zero")
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	//Embedded struct
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	// split string

	splittedData := strings.Split(datastring, ",")
	if len(splittedData) != 2 {
		return ErrWrongData
	}

	// convert data

	steps, err := strconv.Atoi(splittedData[0])

	if err != nil {
		return err
	}

	if steps <= 0 {
		return ErrNegativeSteps
	}

	duration, err := time.ParseDuration(splittedData[1])

	if err != nil {
		return err
	}

	if duration <= 0 {
		return ErrNegativeDuration
	}

	ds.Steps, ds.Duration = steps, duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	dist := spentenergy.Distance(ds.Steps, ds.Height)

	enrgy, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, enrgy), nil

}
