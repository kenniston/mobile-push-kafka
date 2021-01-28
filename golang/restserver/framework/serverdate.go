package framework

import (
	"errors"
	"strings"
	"time"
)

//===============================================================================
// Date is a structure which defines Marshal and Unmarshal functions
// for specific date format "2006-01-02"
//
type Date struct {
	time.Time
}

func (s *Date) MarshalJSON() ([]byte, error) {
	if y := s.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(s.Time.Format(`"2006-01-02"`)), nil
}

func (s *Date) MarshalText() ([]byte, error) {
	if y := s.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(s.Time.Format(`2006-01-02`)), nil
}

func (s *Date) UnmarshalJSON(input []byte) error {
	str := strings.Trim(string(input), "\"")
	if str == "null" {
		s.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	s.Time = t
	return nil
}

func (s *Date) Format(str string) string {
	return s.Time.Format(str)
}

func UnserializeDate(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		t := Date{}
		err := t.UnmarshalJSON(value)
		if err != nil {
			return nil
		}
		return t
	case string:
		return UnserializeDate([]byte(value))
	case *string:
		if value == nil {
			return nil
		}
		return UnserializeDate([]byte(*value))
	case Date:
		return value
	default:
		return nil
	}
}

func SerializeDate(value interface{}) interface{} {
	switch value := value.(type) {
	case Date:
		return SerializeDate(&value)
	case *Date:
		if value == nil {
			return nil
		}

		buff, err := value.MarshalText()
		if err != nil {
			return nil
		}
		return string(buff)
	default:
		return nil
	}
}

//===============================================================================
// DateTime is a structure which defines Marshal and Unmarshal functions
// for specific date format "2006-01-02T15:04:05"
//
type DateTime struct {
	time.Time
}

func (s *DateTime) MarshalJSON() ([]byte, error) {
	if y := s.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("DateTime.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(s.Time.Format(`"2006-01-02T15:04:05"`)), nil
}

func (s *DateTime) MarshalText() ([]byte, error) {
	if y := s.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("DateTime.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(s.Time.Format(`2006-01-02T15:04:05`)), nil
}

func (s *DateTime) UnmarshalJSON(input []byte) error {
	str := strings.Trim(string(input), "\"")
	if str == "null" {
		s.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	s.Time = t
	return nil
}

func (s *DateTime) Format(str string) string {
	return s.Time.Format(str)
}

func UnserializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		t := DateTime{}
		err := t.UnmarshalJSON(value)
		if err != nil {
			return nil
		}
		return t
	case string:
		return UnserializeDate([]byte(value))
	case *string:
		if value == nil {
			return nil
		}
		return UnserializeDate([]byte(*value))
	case DateTime:
		return value
	default:
		return nil
	}
}

func SerializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case DateTime:
		return SerializeDate(&value)
	case *DateTime:
		if value == nil {
			return nil
		}

		buff, err := value.MarshalText()
		if err != nil {
			return nil
		}
		return string(buff)
	default:
		return nil
	}
}

//===============================================================================
// ISO8601DateTime is a structure which defines Marshal and Unmarshal functions
// for specific date format "2006-01-02T15:04:05-0700"
//
type ISO8601DateTime struct {
	time.Time
}

func (s *ISO8601DateTime) MarshalJSON() ([]byte, error) {
	if y := s.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("ISO8601DateTime.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(s.Time.Format(`"2006-01-02T15:04:05-0700"`)), nil
}

func (s *ISO8601DateTime) MarshalText() ([]byte, error) {
	if y := s.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("ISO8601DateTime.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(s.Time.Format(`2006-01-02T15:04:05-0700`)), nil
}

func (s *ISO8601DateTime) UnmarshalJSON(input []byte) error {
	str := strings.Trim(string(input), "\"")
	if str == "null" {
		s.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05-0700", str)
	if err != nil {
		return err
	}
	s.Time = t
	return nil
}

func (s *ISO8601DateTime) Format(str string) string {
	return s.Time.Format(str)
}

func UnserializeISO8601DateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		t := ISO8601DateTime{}
		err := t.UnmarshalJSON(value)
		if err != nil {
			return nil
		}
		return t
	case string:
		return UnserializeISO8601DateTime([]byte(value))
	case *string:
		if value == nil {
			return nil
		}
		return UnserializeISO8601DateTime([]byte(*value))
	case ISO8601DateTime:
		return value
	default:
		return nil
	}
}

func SerializeISO8601DateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case ISO8601DateTime:
		return SerializeISO8601DateTime(&value)
	case *ISO8601DateTime:
		if value == nil {
			return nil
		}
		buff, err := value.MarshalText()
		if err != nil {
			return nil
		}
		return string(buff)
	default:
		return nil
	}
}

//===============================================================================
// ISO8601DateTimeNano is a structure which defines Marshal and Unmarshal functions
// for specific date format "2006-01-02T15:04:05.999-0700"
//
type ISO8601DateTimeNano struct {
	time.Time
}

func (s *ISO8601DateTimeNano) MarshalJSON() ([]byte, error) {
	if y := s.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("ISO8601DateTimeNano.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(s.Time.Format(`"2006-01-02T15:04:05.999-0700"`)), nil
}

func (s *ISO8601DateTimeNano) MarshalText() ([]byte, error) {
	if y := s.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("ISO8601DateTimeNano.MarshalJSON: year outside of range [0,9999]")
	}
	return []byte(s.Time.Format(`2006-01-02T15:04:05.999-0700`)), nil
}

func (s *ISO8601DateTimeNano) UnmarshalJSON(input []byte) error {
	str := strings.Trim(string(input), "\"")
	if str == "null" {
		s.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05.999-0700", str)
	if err != nil {
		return err
	}
	s.Time = t
	return nil
}

func (s *ISO8601DateTimeNano) Format(str string) string {
	return s.Time.Format(str)
}

func UnserializeISO8601DateTimeNano(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		t := ISO8601DateTimeNano{}
		err := t.UnmarshalJSON(value)
		if err != nil {
			return nil
		}
		return t
	case string:
		return UnserializeISO8601DateTimeNano([]byte(value))
	case *string:
		if value == nil {
			return nil
		}
		return UnserializeISO8601DateTimeNano([]byte(*value))
	case ISO8601DateTimeNano:
		return value
	default:
		return nil
	}
}

func SerializeISO8601DateTimeNano(value interface{}) interface{} {
	switch value := value.(type) {
	case ISO8601DateTimeNano:
		return SerializeISO8601DateTimeNano(&value)
	case *ISO8601DateTimeNano:
		if value == nil {
			return nil
		}

		buff, err := value.MarshalText()
		if err != nil {
			return nil
		}
		return string(buff)
	default:
		return nil
	}
}
