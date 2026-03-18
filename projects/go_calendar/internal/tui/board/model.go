package board

import (
	"tcalendar/internal/calendar"
	"time"
)

type DayCell struct {
	Day     int
	Weekday time.Weekday
	Month   time.Month
	Year    int
}

type Model struct {
	currentDate time.Time

	Year  int
	Month time.Month
	Days  []DayCell
}

func newModel() *Model {

	now := time.Now()

	year := now.Year()
	month := now.Month()
	days := makeDays(year, month)

	return &Model{
		Year:        year,
		Month:       month,
		Days:        days,
		currentDate: now,
	}
}

func (m *Model) NextMonth() {
	nextMonth := m.Month + 1
	if nextMonth > time.December {
		nextMonth = time.January
		m.Year += 1
	}
	m.Month = nextMonth
	m.updateDays()
}

func (m *Model) PrevMonth() {
	prevMonth := m.Month - 1
	if prevMonth < time.January {
		prevMonth = time.December
		m.Year -= 1
	}
	m.Month = prevMonth
	m.updateDays()
}

func (m *Model) NextYear() {
	m.Year += 1
	m.updateDays()
}

func (m *Model) PrevYear() {
	m.Year -= 1
	m.updateDays()
}

func (m *Model) IsCurrent(day DayCell) bool {
	return day.Year == m.currentDate.Year() && day.Month == m.currentDate.Month() && day.Day == m.currentDate.Day()
}

func (m *Model) updateDays() {
	m.Days = makeDays(m.Year, m.Month)
	m.currentDate = time.Now()
}

func makeDays(year int, month time.Month) []DayCell {
	loc, _ := time.LoadLocation("Europe/Moscow")
	result := []DayCell{}

	days, _ := calendar.DaysInMonth(year, month)

	for i := range days {
		weekday, _ := calendar.Weekday(year, month, i+1, loc)
		result = append(result, DayCell{Day: i + 1, Weekday: weekday, Year: year, Month: month})
	}

	return result

}
