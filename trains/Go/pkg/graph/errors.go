package graph

type PathError struct {
	Path    *Path
	Message string
}

func NewPathError(p *Path, m string) *PathError {
	return &PathError{
		Path:    p,
		Message: m,
	}
}

func (pe *PathError) Error() string {
	return pe.Message
}
