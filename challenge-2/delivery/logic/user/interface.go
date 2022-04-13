package user

type User interface {
	ValidationStruct(req Req) error
}
