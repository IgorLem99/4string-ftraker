package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	dist := distance(action)
	return dist / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	var dist, speed, calories float64
	switch trainingType {
	case "Бег":
		dist = distance(action)
		speed = meanSpeed(action, duration)
		calories = RunningSpentCalories(action, weight, duration)
	case "Ходьба":
		dist = distance(action)
		speed = meanSpeed(action, duration)
		calories = WalkingSpentCalories(action, duration, weight, height)
	case "Плавание":
		dist = float64(lengthPool*countPool) / mInKm
		speed = swimmingMeanSpeed(lengthPool, countPool, duration)
		calories = SwimmingSpentCalories(lengthPool, countPool, duration, weight)
	default:
		return "неизвестный тип тренировки"
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 0.063 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 3.5   // сдвиг калорий для среднего бега.
)

// RunningSpentCalories возвращает количество потраченных калорий при беге.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	speed := meanSpeed(action, duration)
	return (runningCaloriesMeanSpeedMultiplier * speed * weight * duration * minInH)
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	speed := meanSpeed(action, duration)
	return (walkingCaloriesWeightMultiplier*weight + (math.Pow(speed*kmhInMsec, 2)/(height/cmInM))*walkingSpeedHeightMultiplier*weight) * duration * minInH
}

// Константы для расчета калорий, расходуемых при плавании.
const (
	swimmingCaloriesMeanSpeedShift   = 1.1 // сдвиг калорий для среднего плавания.
	swimmingCaloriesWeightMultiplier = 2.0 // множитель веса при плавании.
)

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool*countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}

func main() {
	action := 5000
	trainingType := "Бег"
	duration := 1.5
	weight := 70.0
	height := 175.0
	lengthPool := 25
	countPool := 40

	info := ShowTrainingInfo(action, trainingType, duration, weight, height, lengthPool, countPool)
	fmt.Println(info)
}
