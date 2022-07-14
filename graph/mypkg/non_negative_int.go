package mypkg

import (
	"fmt"
	"io"
)

type NonNegativeInt int

func (value *NonNegativeInt) UnmarshalGQL(v interface{}) error {
	switch data := v.(type) {
	case int:
		if data < 0 {
			return fmt.Errorf("value is inferior to 0")
		}
		*value = NonNegativeInt(data)
	case int64:
		if data < 0 {
			return fmt.Errorf("value is inferior to 0")
		}
		*value = NonNegativeInt(data)
	default:
		return fmt.Errorf("value must be an int")
	}
	return nil
}

func (value NonNegativeInt) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("\"%d\"", value)))
}
