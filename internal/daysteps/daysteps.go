package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	dataSplit := strings.Split(data, ",")
	if len(dataSplit) != 2 {
		return 0, 0, fmt.Errorf("Неверный формат")
	}
	stepsInt, err := strconv.Atoi(dataSplit[0])
	if err != nil {
		return 0, 0, err
	}
	if stepsInt <= 0 {
		return 0, 0, fmt.Errorf("Количество шагов указано неверно")
	}

	timeDur, err := time.ParseDuration(dataSplit[1])
	if err != nil {
		return 0, 0, err
	}
	if timeDur <= 0 {
		return 0, 0, fmt.Errorf("Time must be positive")
	}
	return stepsInt, timeDur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}

	dist := (float64(steps) * stepLength) / float64(mInKm)
	calories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	report := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, dist, calories)
	return report
}
