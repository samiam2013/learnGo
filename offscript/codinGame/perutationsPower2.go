package codingame

import (
	"fmt"
)

const op = "?"
const vn = "_"

func getFormulas(varName string, numVars int) []string {

	operators := []string{"**", "+", "-"}
	operands := []string{"2", varName}

	formulaTemplate := ""
	for i := 0; i < numVars-1; i++ {
		formulaTemplate += vn + op
	}
	formulaTemplate += vn

	// pass the formula template, list of operators & operands, get back formulas
	output := recurseFillOperand(operators, operands, 2, "")
	fmt.Printf("%#v\n", output)

	return output
}

//and DFS search and return possible
func recurseFillOperand(operators, operands []string, nOps int, parent string) (output []string) {
	if nOps < 1 {
		return []string{parent}
	}
	for _, operand := range operands {
		output = append(output, recurseFillOperator(operators, operands, nOps-1, parent+operand)...)
	}
	return output
}

func recurseFillOperator(operators, operands []string, nOps int, parent string) (output []string) {
	if nOps < 1 {
		return []string{parent}
	}
	for _, operator := range operators {
		output = append(output, recurseFillOperand(operators, operands, nOps, parent+operator)...)
	}
	return output
}
