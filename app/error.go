package app

type BusinessError struct {
	message string
}

func (be BusinessError) Error() string {
	return be.message
}
