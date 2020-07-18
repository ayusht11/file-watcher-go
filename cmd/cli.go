package cmd

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Cmd is to run command line program
type Cmd struct{}

// Add will run the addition program
func (c *Cmd) Start_watcher() {

	//Scan source directory
	fmt.Print("Enter the source directory path: ")
	sourceDirectory := fetchFilePath()

	//Scan target directory
	fmt.Print("Enter the target directory path: ")
	targetDirectory := fetchFilePath()

	//Scan the AES key for encryption
	fmt.Print("Enter the passphrase for encryption: ")
	passphrase := fetchFilePath()

	//Initializing the watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error while starting the watcher ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event:", event)
				encryptFile(event.Name, targetDirectory, passphrase)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(sourceDirectory)
	if err != nil {
		fmt.Println("error:", err)
	}
	<-done

}

// Scan the entries
func fetchFilePath() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func encryptFile(sourceFile string, targetDirectory string, passphrase string) {
	os.Setenv("TARGET_DIRECTORY", targetDirectory)
	os.Setenv("PASSPHRASE", passphrase)

	plaintext, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		panic(err.Error())
	}

	// this is a key
	key := []byte(passphrase)
	data := []byte(plaintext)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	nonceSize := gcm.NonceSize()

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	s := string(ciphertext)
	filename := filepath.Base(sourceFile)
	targetFile := filepath.Join(targetDirectory, filename)

	// create a new file for saving the encrypted data.
	f, err := os.Create(targetFile)

	f.Write(ciphertext)
	if err != nil {
		panic(err.Error())
	}
}
