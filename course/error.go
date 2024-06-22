package course

//mapear los errores del servicio
type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}
