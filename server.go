package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

// File contains name of a file
type File struct {
	Name string `json:"name,omitempty"`
}

// FileListResponse defines JSON response to list files request
type FileListResponse struct {
	Files []File `json:"files,omitempty"`
}

var files []File

var targetDirectory = os.Getenv("TARGET_DIRECTORY")
var passphrase = os.Getenv("PASSPHRASE")

// GetFileEndpoint Controller logic to get files for the given filename
func GetFileEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	filename := params["filename"]
	filepath := filepath.Join(targetDirectory, filename)

	decryptedData := decryptFile(filepath, "testtesttesttest")

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	io.Copy(w, bytes.NewReader(decryptedData))
}

// GetFilesEndpoint Controller logic to get all files within targetdirectory
func GetFilesEndpoint(w http.ResponseWriter, req *http.Request) {
	_files, err := ioutil.ReadDir(targetDirectory)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range _files {
		files = append(files, File{Name: f.Name()})
	}

	json.NewEncoder(w).Encode(FileListResponse{Files: files})
}

func decryptFile(filepath string, passphrase string) []byte {
	ciphertext, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err.Error())
	}

	key := []byte(passphrase)
	data := []byte(ciphertext)

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

	nonce, cipherbytes := data[:nonceSize], data[nonceSize:]

	decryptedBytes, err := gcm.Open(nil, nonce, cipherbytes, nil)
	if err != nil {
		panic(err.Error())
	}

	return decryptedBytes
}

func main() {
	fmt.Println("Within Go Server App")
	router := mux.NewRouter()

	router.HandleFunc("/v1/files", GetFilesEndpoint).Methods("GET")
	router.HandleFunc("/v1/files/{filename}", GetFileEndpoint).Methods("GET")

	fmt.Println("Starting server on: http://localhost:4000/")
	log.Fatal(http.ListenAndServe(":4000", router))
}
