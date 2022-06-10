package psp

type Connection interface {
	Do(Request) ([]byte, error)
}
