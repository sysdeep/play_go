package calendar

import (
	"fmt"
	"time"
)

var Months = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
}

// DaysInMonth возвращает количество дней в заданном году и месяце.
// month – число от 1 (январь) до 12 (декабрь). Если month находится вне диапазона,
// функция вернёт ошибку.
func DaysInMonth(year int, month time.Month) (int, error) {
	if month < time.January || month > time.December {
		return 0, fmt.Errorf("неверный месяц: %d", month)
	}

	// Создаём дату: первое число следующего месяца.
	// Затем вычитаем один день – получаем последний день текущего месяца.
	// Его поле Day() и есть количество дней в месяце.
	nextMonth := month + 1
	nextYear := year
	if nextMonth > time.December {
		nextMonth = time.January
		nextYear++
	}
	// 0‑й день в Go считается последним днём предыдущего месяца,
	// поэтому удобно создать дату 0‑го дня текущего месяца:
	//   time.Date(year, month+1, 0, 0,0,0,0, time.UTC)
	// Это эквивалентно: последний день текущего месяца.
	lastDay := time.Date(nextYear, nextMonth, 0, 0, 0, 0, 0, time.UTC)

	return lastDay.Day(), nil
}

func Weekday(year int, month time.Month, day int, loc *time.Location) (time.Weekday, error) {
	// time.Date корректно обрабатывает «переполняющиеся» значения,
	// поэтому простая проверка диапазона достаточно.
	if month < time.January || month > time.December {
		return 0, fmt.Errorf("недопустимый месяц: %d", month)
	}
	if day < 1 || day > 31 {
		return 0, fmt.Errorf("недопустимый день: %d", day)
	}
	t := time.Date(year, month, day, 0, 0, 0, 0, loc)
	return t.Weekday(), nil
}

// var ruWeekday = map[time.Weekday]string{
// 	time.Sunday:    "Воскресенье",
// 	time.Monday:    "Понедельник",
// 	time.Tuesday:   "Вторник",
// 	time.Wednesday: "Среда",
// 	time.Thursday:  "Четверг",
// 	time.Friday:    "Пятница",
// 	time.Saturday:  "Суббота",
// }

var Weekdays = []time.Weekday{
	time.Monday,
	time.Tuesday,
	time.Wednesday,
	time.Thursday,
	time.Friday,
	time.Saturday,
	time.Sunday,
}

var ruMonths = map[time.Month]string{
	time.January:   "Январь",
	time.February:  "Февраль",
	time.March:     "Март",
	time.April:     "Апрель",
	time.May:       "Май",
	time.June:      "Июнь",
	time.July:      "Июль",
	time.August:    "Август",
	time.September: "Сентябрь",
	time.October:   "Октябрь",
	time.November:  "Ноябрь",
	time.December:  "Декабрь",
}

func MonthName(m time.Month) string {
	return ruMonths[m]
}
