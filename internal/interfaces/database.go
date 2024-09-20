package interfaces

type KeyPairBased interface {
	Get(key string) (interface{}, error)
	Set(key string, value string) error
	Insert(key string, data interface{}) error
	Select(key string, field string) (string, error)
	SelectAll(key string, output interface{}) error
	Delete(key string) (int64, error)
}
