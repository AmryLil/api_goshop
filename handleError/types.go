package handleError

type NotFoundError struct {
	Message     string
	StatusError string
}
type BadRequestError struct {
	Message     string
	StatusError string
}
type InternalServerError struct {
	Message string
}
type UnauthorizedError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}
func (e BadRequestError) Error() string {
	return e.Message
}
func (e InternalServerError) Error() string {
	return e.Message
}
func (e UnauthorizedError) Error() string {
	return e.Message
}
