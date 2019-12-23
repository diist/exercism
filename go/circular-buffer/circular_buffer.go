// Package circular contains the logic for the Exercism go exercise "circular buffer"
package circular

import "errors"

// Buffer is a buffer of bytes
type Buffer struct {
	c chan byte
}

// NewBuffer creates a new buffer of a given size
func NewBuffer(size int) *Buffer {
	b := &Buffer{make(chan byte, size)}
	return b
}

// ReadByte returns a byte from the buffer
func (b *Buffer) ReadByte() (byte, error) {
	if len(b.c) == 0 {
		return byte(0), errors.New("nothing to read, buffer empty")
	}
	return <-b.c, nil
}

// WriteByte writes a byte to the buffer
func (b *Buffer) WriteByte(c byte) error {
	if b.full() {
		return errors.New("cannot write, buffer is full")
	}
	b.c <- c
	return nil
}

// Overwrite pushes a value out if the buffer is empty
func (b *Buffer) Overwrite(c byte) {
	if b.full() {
		<-b.c
	}
	b.c <- c
}

// Reset flushes out the values from the buffer
func (b *Buffer) Reset() {
	b.c = make(chan byte, cap(b.c))
}

func (b *Buffer) full() bool {
	return len(b.c) == cap(b.c)
}
