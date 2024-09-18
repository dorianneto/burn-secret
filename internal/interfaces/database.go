package interfaces

type KeyPairBased interface {
	Get(key string) (interface{}, error)
	Set(key string, value string) error
}
