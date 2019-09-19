package types

type Request interface {
	Validate() error
}
