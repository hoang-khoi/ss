package strategy

import (
	"encoding/binary"
	"io"
	"unsafe"
)

// EncryptedChunk wraps a block of encrypted bytes.
// [NonceSize][Nonce...][BlockSize][Block...]
// NonceSize: 4 bytes, int32, little endian
// Size: 4 bytes, int32, little endian
type EncryptedChunk struct {
	Nonce []byte
	Block []byte
}

// Serialize writes an EncryptedChunk to an io.Writer instance.
func (c *EncryptedChunk) Serialize(w io.Writer) (err error) {
	// Write [NonceSize]
	err = binary.Write(w, binary.LittleEndian, int32(len(c.Nonce)))

	if err != nil {
		return
	}

	// Write [Nonce]
	_, err = w.Write(c.Nonce)
	if err != nil {
		return
	}

	// Write [BlockSize]
	err = binary.Write(w, binary.LittleEndian, int32(len(c.Block)))
	if err != nil {
		return
	}

	// Write [Block]
	_, err = w.Write(c.Block)
	if err != nil {
		return
	}

	return
}

// Deserialize parses an EncryptedChunk from io.Reader, returns the number of bytes read.
func (c *EncryptedChunk) Deserialize(r io.Reader) (n int, err error) {
	var nonceSize, blockSize int32

	// Read nonce size
	err = binary.Read(r, binary.LittleEndian, &nonceSize)
	n += int(unsafe.Sizeof(nonceSize))
	if err != nil {
		return
	}

	// Load nonce
	c.Nonce = make([]byte, nonceSize)
	m, err := r.Read(c.Nonce)
	n += m
	if err != nil {
		return
	}

	// Read block size
	err = binary.Read(r, binary.LittleEndian, &blockSize)
	n += int(unsafe.Sizeof(blockSize))
	if err != nil {
		return
	}

	// Load block
	c.Block = make([]byte, blockSize)
	m, err = r.Read(c.Block)
	n += m
	if err != nil {
		return
	}

	return
}
