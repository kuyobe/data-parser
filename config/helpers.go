package dataparser

import (
	"fmt"
	"time"
)

func NewDateParser(layout, dateStr string) (DateParser, error) {
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, err
	}
	return NewDateParserFromTime(t), nil
}

func NewDateParserFromTime(t time.Time) DateParser {
	return func(dateStr string) (time.Time, error) {
		t, err := time.Parse(time.Layout, dateStr)
		return t, err
	}
}

func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(time.Layout, dateStr)
}

func ParseDateWithLayout(layout, dateStr string) (time.Time, error) {
	return time.Parse(layout, dateStr)
}

type DateParser func(dateStr string) (time.Time, error)

func (p DateParser) Parse(dateStr string) (time.Time, error) {
	return p(dateStr)
}

func (p DateParser) String() string {
	return fmt.Sprintf("DateParser(%s)", reflect.TypeOf(p).Elem().Name())
}