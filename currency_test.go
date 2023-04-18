package currency

import "testing"

func TestTodayDate(t *testing.T) {
	// TODO Hardcoded date will be problem.
	today := "18042023"
	got := todayDate()
	if today != got {
		t.Errorf("expected: %s\ngot: %s", today, got)
	}
	t.Logf("got: %s", got)
}

func TestNew(t *testing.T) {
	c := New("USD")
	t.Log(c.currency)
}
