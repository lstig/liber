package dal

type Metadata struct {
	Size int
}

// Operator provides a unified interface for reading and writing data from various backends
type Operator interface {
	// Read data into a byte array from a path
	Read(path string, p []byte) (n int, err error)
	// Write data from a byte array to a path
	Write(path string, p []byte) (n int, err error)
	// Stat return information about a file at the given path
	Stat(path string) (meta *Metadata, err error)
}

// Builder defines the Build function which is used to initialize an operator
type Builder interface {
	Build() (Operator, error)
}

// NewOperator given a builder, returns an intialized Operator
func NewOperator(builder Builder) (Operator, error) {
	op, err := builder.Build()
	if err != nil {
		return nil, err
	}
	return op, nil
}
