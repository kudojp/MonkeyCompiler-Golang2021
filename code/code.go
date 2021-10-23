package code

import "fmt"

type Instructions []byte
type Opcode byte

const (
	OpConstant Opcode = iota
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}}, // Thus, up to 65536 constants could be defied.
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

/*
Build byte array (instrcutions) from Opcode + Operands
*/
func Make(op Opcode, operands ...int) []byte {
	return nil
}
