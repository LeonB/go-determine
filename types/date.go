package types

import (
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalText(text []byte) error {
	var err error
	s := string(text)
	if s == "" {
		return nil
	}

	// first try standard date
	layout := time.RFC3339
	d.Time, err = time.Parse(layout, s)
	if err == nil {
		return nil
	}

	layout = "2006-01-02T15:04"
	d.Time, err = time.Parse(layout, s)
	if err == nil {
		return err
	}

	return nil
}

func (d Date) MarshalText() ([]byte, error) {
	if d.IsZero() {
		return []byte{}, nil
	}

	layout := "2006-01-02T15:04"
	return []byte(d.Time.Format(layout)), nil
}
