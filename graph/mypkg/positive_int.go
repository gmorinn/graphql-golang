package mypkg

import (
	"fmt"
	"io"
)

type PositiveInt int

func (value *PositiveInt) UnmarshalGQL(v interface{}) error {
	switch data := v.(type) {
	case int:
		if data <= 0 {
			return fmt.Errorf("value needs to be positive or 0")
		}
		*value = PositiveInt(data)
	case int64:
		if data <= 0 {
			return fmt.Errorf("value needs to be positive or 0")
		}
		*value = PositiveInt(data)
	default:
		return fmt.Errorf("value must be an int")
	}
	return nil
}

func (value PositiveInt) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("\"%d\"", value)))
}
