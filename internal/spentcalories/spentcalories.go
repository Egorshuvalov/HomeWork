package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	splitData := strings.Split(data, ",")

	if len(splitData) != 3 {
		return 0, "", 0, fmt.Errorf("Неверный ввод")
	}
	steps, err := strconv.Atoi(splitData[0])

	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("Wrong steps type")
	}
	if splitData[1] == "" {
		return 0, "", 0, fmt.Errorf("Wrong training type")
	}
	duration, err := time.ParseDuration(splitData[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("Wrong time type")
	}
	return steps, splitData[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	oneStep := height * stepLengthCoefficient
	dist := (oneStep * float64(steps)) / mInKm
	return dist

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	speedHours := duration.Hours()
	mSpeed := dist / speedHours
	return mSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activ, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if steps <= 0 {
		log.Println("steps must be positive")
		return "", fmt.Errorf("steps must be positive")

	}
	if duration <= 0 {
		log.Println("training time must be positive")
		return "", fmt.Errorf("training time must be positive")
	}

	resDuration := duration.Hours()
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	switch activ {

	case "Бег":
		calories, _ := RunningSpentCalories(steps, weight, height, duration)
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activ, resDuration, dist, speed, calories)
		return result, nil

	case "Ходьба":
		calories, _ := WalkingSpentCalories(steps, weight, height, duration)
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activ, resDuration, dist, speed, calories)
		return result, nil

	default:

		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Invalid input")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	spentCalories := (weight * speed * minutes) / float64(minInH)
	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Invalid input")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	spentCalories := (weight * speed * minutes * walkingCaloriesCoefficient) / float64(minInH)
	return spentCalories, nil
}
