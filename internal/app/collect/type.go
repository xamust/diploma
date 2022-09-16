package collect

// ParsingF интерфейс для реализации Collect
type ParsingF interface {
	initMap()
	checkData() error
	FindFiles() (bool, error)
	readDir(directory string) error
}

// Coll интерфейс для реализации Collect
type Coll interface {
	Start() error
	searchDataFiles() (map[string]string, error)
}
