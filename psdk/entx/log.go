package entx

type Log interface {
	Debug(string, ...any)
}

func NewLog(msg string, log Log) func(...any) {
	return func(args ...any) {
		log.Debug(msg, args...)
	}
}
