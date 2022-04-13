package movie

type Movie interface {
	ValidationStruct(req Req) error
}
