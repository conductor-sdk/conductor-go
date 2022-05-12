package def

type ConductorWorkflow struct {
	Name    string
	Version int32
}

func (workflow *ConductorWorkflow) register(overwrite bool) (string, error) {
	return "", nil
}

func (workflow *ConductorWorkflow) execute() (string, error) {
	return "", nil
}
