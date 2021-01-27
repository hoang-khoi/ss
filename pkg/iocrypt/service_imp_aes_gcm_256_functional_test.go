package iocrypt

import (
	"crypto/sha256"
	"io/ioutil"
	"log"
	"os"
	"ss/pkg/iocrypt/strategy"
	"testing"

	"github.com/stretchr/testify/assert"
)

const passphrase = "TopSecret"

const (
	originalPath  = "assets/original"
	encryptedPath = "assets/encrypted"
	decryptedPath = "assets/decrypted"
)

var (
	key                    = sha256.Sum256([]byte(passphrase))
	encryptService Service = &ServiceImp{Strategy: &strategy.AesGcm256Encrypt{Key: key}}
	decryptService Service = &ServiceImp{Strategy: &strategy.AesGcm256Decrypt{Key: key}}
)

func TestServiceImp_Process_Functionality(t *testing.T) {
	encrypt(t)
	decrypt(t)
	verify(t)
	cleanup(t)
}

func cleanup(t *testing.T) {
	err := os.Remove(encryptedPath)
	err = os.Remove(decryptedPath)
	if err != nil {
		t.Fatal(err)
	}
}

//goland:noinspection GoUnhandledErrorResult
func verify(t *testing.T) {
	original, err := os.Open(originalPath)
	decrypted, err := os.Open(decryptedPath)
	if err != nil {
		t.Fatal(err)
	}
	defer original.Close()
	defer decrypted.Close()

	originalBuffer, err := ioutil.ReadAll(original)
	decryptedBuffer, err := ioutil.ReadAll(decrypted)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, originalBuffer, decryptedBuffer)
}

func decrypt(t *testing.T) {
	process(t, encryptedPath, decryptedPath, decryptService)
}

func encrypt(t *testing.T) {
	process(t, originalPath, encryptedPath, encryptService)
}

//goland:noinspection ALL
func process(t *testing.T, srcPath string, destPath string, service Service) {
	src, err := os.Open(srcPath)
	dest, err := os.Create(destPath)
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()
	defer dest.Close()

	srcStat, err := src.Stat()
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan int)
	go func() { err = service.Process(dest, src, c) }()

	var n int64
	for i := range c {
		n += int64(i)
	}
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, srcStat.Size(), n)
}
