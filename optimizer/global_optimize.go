package optimize

import (
	"github.com/KazumaTakata/static-typed-language/IR-gen"
)

type Block struct {
	codes []ir_gen.IR_Code
	In    map[string]bool
	Out   map[string]bool
}

func Construct_CFG(codes []ir_gen.IR_Code) ([]*Block, map[string]*Block) {
	blocks := []*Block{}

	block := &Block{}

	labeled_block := map[string]*Block{}

	for _, code := range codes {

		if code.Type == ir_gen.Label {
			blocks = append(blocks, block)
			block = &Block{}
			block.codes = append(block.codes, code)
			labeled_block[code.Right_Operand1.String] = block
		} else if code.Type == ir_gen.Ifz || code.Type == ir_gen.Goto {
			block.codes = append(block.codes, code)
			blocks = append(blocks, block)
			block = &Block{}
		} else {
			block.codes = append(block.codes, code)
		}
	}

	return blocks, labeled_block

}
