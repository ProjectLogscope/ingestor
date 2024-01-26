package logger

type handlerType int8

const (
	TypeText handlerType = iota
	TypeJSON
)
