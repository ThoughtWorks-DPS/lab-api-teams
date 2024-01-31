package datastore

type Datastore interface {
	Create(data interface{}) error
	ReadByID(id string, out interface{}) error
	ReadByAttributes(filter Filter, out interface{}) error
	ReadByAttributesWithPagination(filter Filter, out interface{}, page int, maxResult int) error
	ReadAll(out interface{}) error
	Update(data interface{}) error
	Delete(data interface{}) error
	IsDatabaseAvailable() (bool, error)
}

type Migratable interface {
	Migrate(models ...interface{}) error
}

type Filter map[string]interface{}
