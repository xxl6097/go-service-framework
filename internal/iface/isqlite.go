package iface

type ISqlite[V any] interface {
	AddJson(v map[string]interface{}) error
	Add(*V) error
	Update(map[string]interface{}) error
	Delete(*V) error
	DeleteByUniqueKey(string) error
	Find(*V) (*V, error)
	FindAll() ([]V, error)
	First() (*V, error)
	Save(procModel *V) error
}
