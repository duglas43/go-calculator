package executor

type ExecuteRequest struct {
	Instructions []Instruction `json:"instructions"`
}
