package auth

type Auth interface {
	Login(userName, password string) (string, error)
}
