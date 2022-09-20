package domain

type InternalError struct {
	Err string
}

func (i InternalError) Error() string {
	return i.Err
}
