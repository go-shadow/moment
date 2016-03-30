package moment

import (
	"testing"
	"time"
)

var (
	oneDay = time.Duration(time.Hour * 24)
)

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

	yesterday := New().Strtotime("yesterday").GetTime()
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

	now := time.Now().Truncate(time.Second).Add(time.Nanosecond*112233445)

	for _, p := range pairs {
		if timeToCheck, ok := checkTimeFormat(now, p.Format, p.Value); !ok {
			t.Fatalf("Expected %s, got %s", p.Value, timeToCheck.Format(p.Format))
		}

	}
}
