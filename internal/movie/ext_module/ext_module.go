package ext_module

type (
	ExtModule interface {
		NewOMDB() IOmdb
	}
	extModule struct{}
)

func NewExtModule() ExtModule {
	return &extModule{}
}
