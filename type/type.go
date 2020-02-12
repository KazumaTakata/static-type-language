package basic_type

import "fmt"

type Type int

const (
	INT Type = iota + 1
	DOUBLE
	STRING
)

func (e Type) String() string {

	switch e {
	case INT:
		return "INT"
	case DOUBLE:
		return "DOUBLE"
	case STRING:
		return "STRING"

	default:
		return fmt.Sprintf("NULL")
	}
}
