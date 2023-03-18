package services

//go:generate mockgen -source=interface.go -destination=mock/mock.go -package=mock
type IDomain interface {
	GetResponseBody() []byte
}
