package strategy

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// AesGcm256Encrypt encrypting strategy implemented with AES-GCM-256.
type AesGcm256Encrypt struct {
	Key [32]byte
}

const bufferSize = 4096

// Process pulls bufferSize bytes from src, encrypts and dumps to dest. Returns the number of bytes read.
func (s *AesGcm256Encrypt) Process(dest io.Writer, src io.Reader) (n int, err error) {
	buffer := make([]byte, bufferSize)

	// Trying to pull a block of data
	n, err = src.Read(buffer)
	if err != nil {
		return
	}

	// Prepare encryption
	block, err := aes.NewCipher(s.Key[:])
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Reader.Read(nonce)
	if err != nil {
		return
	}

	// Encrypt the data block and put it into an EncryptedChunk instance
	chunk := &EncryptedChunk{
		Nonce: nonce,
		Block: gcm.Seal(nil, nonce, buffer[:n], nil),
	}

	err = chunk.Serialize(dest)
	return
}
