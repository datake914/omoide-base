package exception

type ServiceError interface {
	Error() string
	HttpStatusCode() int
	MessageCode() string
	Message() string
}
