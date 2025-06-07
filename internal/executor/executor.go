package executor

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type InstructionType string

const (
	CalcType  InstructionType = "calc"
	PrintType InstructionType = "print"
)

type Instruction struct {
	Type  InstructionType `json:"type"`
	Op    string          `json:"op,omitempty"`
	Var   string          `json:"var"`
	Left  any             `json:"left,omitempty"`
	Right any             `json:"right,omitempty"`
}

type Result struct {
	Var   string `json:"var"`
	Value int64  `json:"value"`
}

type Executor struct {
	vars map[string]int64
	mu   sync.RWMutex
}

func NewExecutor() *Executor {
	return &Executor{
		vars: make(map[string]int64),
	}
}

func (e *Executor) Execute(instructions []Instruction) ([]Result, error) {
	results := []Result{}

	for _, instr := range instructions {
		switch instr.Type {
		case CalcType:
			left, err := e.resolveOperand(instr.Left)
			if err != nil {
				return nil, fmt.Errorf("error resolving left operand %s: %w", instr.Var, err)
			}
			right, err := e.resolveOperand(instr.Right)
			if err != nil {
				return nil, fmt.Errorf("error resolving right operand %s: %w", instr.Var, err)
			}

			time.Sleep(50 * time.Millisecond)

			var val int64
			switch instr.Op {
			case "+":
				val = left + right
			case "-":
				val = left - right
			case "*":
				val = left * right
			default:
				return nil, fmt.Errorf("unknown operation %q", instr.Op)
			}

			e.mu.Lock()
			e.vars[instr.Var] = val
			e.mu.Unlock()

		case PrintType:
			val, err := e.resolveVar(instr.Var)
			if err != nil {
				return nil, fmt.Errorf("error resolving variable %q for print: %w", instr.Var, err)
			}
			results = append(results, Result{Var: instr.Var, Value: val})

		default:
			return nil, fmt.Errorf("unknown instruction type %q", instr.Type)
		}
	}

	return results, nil
}

func (e *Executor) resolveOperand(op any) (int64, error) {
	switch v := op.(type) {
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case string:
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			return n, nil
		}
		return e.resolveVar(v)
	default:
		return 0, fmt.Errorf("unsupported operand type %T", v)
	}
}

func (e *Executor) resolveVar(name string) (int64, error) {
	e.mu.RLock()
	val, ok := e.vars[name]
	e.mu.RUnlock()
	if !ok {
		return 0, nil
	}
	return val, nil
}
