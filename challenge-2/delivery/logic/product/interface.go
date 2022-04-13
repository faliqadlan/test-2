package product

type Product interface {
	ValidationStruct(req Req) error
}
