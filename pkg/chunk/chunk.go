package chunk

import "io"

type Chunk struct {
	object io.Reader
	size   uint
}

func NewChunk(object io.Reader, size uint) *Chunk {
	return &Chunk{object: object, size: size}
}

func (c Chunk) Split() [][]io.Reader {

	return nil
}
