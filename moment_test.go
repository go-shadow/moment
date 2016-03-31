package moment

import (
	"testing"
	"reflect"
	"time"
)

var (
	oneDay = time.Duration(time.Hour * 24)
)

// Test some times at the start of the day with some daylight savings
func TestStartOfDay(t *testing.T) {

	lon, _ := time.LoadLocation("Europe/London")

	testCases := []struct {
		before *Moment
		after  *Moment
	}{
		{
			// Clocks +1 hour
			before: NewMoment(time.Date(2015, 3, 29, 0, 0, 0, 0, lon)),
			after:  NewMoment(time.Date(2015, 3, 29, 0, 0, 0, 0, lon)),
		},
		{
			// Clocks -1 hour
			before: NewMoment(time.Date(2015, 10, 25, 2, 0, 0, 0, lon)),
			after:  NewMoment(time.Date(2015, 10, 25, 0, 0, 0, 0, lon)),
		},
		{
			// Clocks -1 hour (same as above)
			before: NewMoment(time.Date(2015, 10, 25, 12, 0, 0, 0, lon)),
			after:  NewMoment(time.Date(2015, 10, 25, 0, 0, 0, 0, lon)),
		},
		{
			// Clocks no change
			before: NewMoment(time.Date(2016, 1, 01, 10, 0, 0, 0, lon)),
			after:  NewMoment(time.Date(2016, 1, 01, 0, 0, 0, 0, lon)),
		},
	}

	for _, test := range testCases {

		res := test.before.StartOfDay()

		if !reflect.DeepEqual(res, test.after) {
			t.Errorf("Moment %s does not match expected %s", res.time.String(), test.after.time.String())
		}
	}
}

func TestYesterdayTodayTomorrow(t *testing.T) {
	now := time.Now()

	today := New().Strtotime("today").GetTime()
	if now.Sub(today) > time.Duration(1*time.Second) {
		t.Fatalf("Expected %s, got %s", now.Format(time.RFC3339), today.Format(time.RFC3339))
	}

	tomorrow := New().Strtotime("tomorrow").GetTime()
	if today.After(tomorrow) {
		t.Fatalf("Expected %s to be before %s", today.Format(time.RFC3339), tomorrow.Format(time.RFC3339))
	}

	tomorrow = New().Strtotime("+1day").GetTime()
	if today.After(tomorrow) {
		t.Fatalf("Expected %s to be before %s", today.Format(time.RFC3339), tomorrow.Format(time.RFC3339))
	}

	yesterday := New().Strtotime("yesterday").GetTime()
	if yesterday.After(today) {
		t.Fatalf("Expected %s to be before %s", yesterday.Format(time.RFC3339), today.Format(time.RFC3339))
	}

	yesterday = New().Strtotime("-1day").GetTime()
	if yesterday.After(today) {
		t.Fatalf("Expected %s to be before %s", yesterday.Format(time.RFC3339), today.Format(time.RFC3339))
	}

	today = today.Truncate(oneDay)
	tomorrow = tomorrow.Truncate(oneDay)
	yesterday = yesterday.Truncate(oneDay)

	diff := tomorrow.Sub(today)
	if diff != oneDay {
		t.Fatalf("Expected different between %s and %s to be one day - got %s", today.Format(time.RFC3339), tomorrow.Format(time.RFC3339), diff)
	}

	diff = today.Sub(yesterday)
	if diff != time.Duration(time.Hour*24) {
		t.Fatalf("Expected different between %s and %s to be one day - got %s", yesterday.Format(time.RFC3339), today.Format(time.RFC3339), diff)
	}
}

func checkTimeFormat(now time.Time, layout string, s string) (time.Time, bool) {
	timeToday := NewMoment(now).Strtotime(s).GetTime()
	if timeToday.Format(layout) != s {
		return timeToday, false
	}

	return timeToday, true
}

func TestHour(t *testing.T) {
	type TimeFormatCheck struct {
		Format string
		Value  string
	}

	pairs := []TimeFormatCheck{
		TimeFormatCheck{
			Format: "15:04",
			Value:  "12:45",
		},
		TimeFormatCheck{
			Format: "15:04:05.999999999",
			Value:  "12:45:01.112233445",
		},
	}

	now := time.Now().Truncate(time.Second).Add(time.Nanosecond * 112233445)

	for _, p := range pairs {
		if timeToCheck, ok := checkTimeFormat(now, p.Format, p.Value); !ok {
			t.Fatalf("Expected %s, got %s", p.Value, timeToCheck.Format(p.Format))
		}

	}
}

func TestNextLast(t *testing.T) {
	var n, d time.Time
	// next monday
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	nextMonday := time.Date(2016, 2, 1, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("next monday").GetTime()
	if nextMonday != n {
		t.Fatalf("Expected next monday to be %v, got %v instead", nextMonday, n)
	}
	
	// last monday
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	lastMonday := time.Date(2016, 1, 25, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("last monday").GetTime()
	if lastMonday != n {
		t.Fatalf("Expected last monday to be %v, got %v instead", lastMonday, n)
	}
	
	// next week
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	nextWeek := time.Date(2016, 2, 6, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("next week").GetTime()
	if nextWeek != n {
		t.Fatalf("Expected next week to be %v, got %v instead", nextWeek, n)
	}
	
	// last week
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	lastWeek := time.Date(2016, 1, 23, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("last week").GetTime()
	if lastWeek != n {
		t.Fatalf("Expected last week to be %v, got %v instead", lastWeek, n)
	}
	
	// next month
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	nextMonth := time.Date(2016, 3, 1, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("next month").GetTime()
	if nextMonth != n {
		t.Fatalf("Expected next month to be %v, got %v instead", nextMonth, n)
	}
	
	// last month
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	lastMonth := time.Date(2015, 12, 30, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("last month").GetTime()
	if lastMonth != n {
		t.Fatalf("Expected last month to be %v, got %v instead", lastMonth, n)
	}
	
	// next year
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	nextYear := time.Date(2017, 1, 30, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("next year").GetTime()
	if nextYear != n {
		t.Fatalf("Expected next year to be %v, got %v instead", nextYear, n)
	}
	
	// last year
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	lastYear := time.Date(2015, 1, 30, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("last year").GetTime()
	if lastYear != n {
		t.Fatalf("Expected last year to be %v, got %v instead", lastYear, n)
	}
}

func TestAgo(t *testing.T) {
	var n, d, newDate time.Time

	// one hour ago
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	newDate = time.Date(2016, 1, 30, 22, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("one hour ago").GetTime()
	if newDate != n {
		t.Fatalf("Expected an hour ago to be %v, got %v instead", newDate, n)
	}

	// one day ago
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	newDate = time.Date(2016, 1, 29, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("one day ago").GetTime()
	if newDate != n {
		t.Fatalf("Expected an day ago to be %v, got %v instead", newDate, n)
	}

	// one week ago
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	newDate = time.Date(2016, 1, 23, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("one week ago").GetTime()
	if newDate != n {
		t.Fatalf("Expected an week ago to be %v, got %v instead", newDate, n)
	}

	// one month ago
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	newDate = time.Date(2015, 12, 30, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("one month ago").GetTime()
	if newDate != n {
		t.Fatalf("Expected an month ago to be %v, got %v instead", newDate, n)
	}
	
	// one year ago
	d = time.Date(2016, 1, 30, 23, 59, 59, 0, time.UTC)
	newDate = time.Date(2015, 1, 30, 23, 59, 59, 0, time.UTC)
	n = NewMoment(d).Strtotime("one year ago").GetTime()
	if newDate != n {
		t.Fatalf("Expected an year ago to be %v, got %v instead", newDate, n)
	}
}
