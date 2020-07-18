# Golang: Addition

- [Golang: Addition](#golang-addition)
  - [Building the source](#building-the-source)
    - [Prerequisites](#prerequisites)
    - [Build](#build)
  - [Usage](#usage)

## Building the source

### Prerequisites

- goLang
- dep: `go get -u github.com/golang/dep/cmd/dep`

### Build

Once you have unzipped the project are you are within project repository:

```sh
cd path/to/project

mv $PWD $GOPATH/src/github.com/file_watcher
cd $GOPATH/src/github.com/file_watcher
```

The build process is set to be automatic you just need to install the package using:

```sh
dep ensure

make tools
```

## Usage

### watcher

To run the watcher you need to execute the following command: `make watcher`

Sample output:

```sh
Enter the source directory path: /home/rails/work/projects/BC/fleek/fleek-test-task/source-folder
Enter the target directory path: /home/rails/work/projects/BC/fleek/fleek-test-task/target-folder
Enter the passphrase for encryption: testtesttesttest
```

As you can observe the command asks for source directory, target directory and the passphrase used in encryption.

> Note: the passphrase should have minimum 16 charcters.

Make the changes in the source directory and the encrypted changes will be reflected in target directory.

### server

You need to specify the `TARGET_DIRECTORY` and `PASSPHRASE` env variables to run our server.

Sample exports:

```sh
TARGET_DIRECTORY=/home/rails/work/projects/BC/fleek/fleek-test-task/source-folder
PASSPHRASE=testtesttesttest

echo $TARGET_DIRECTORY
echo $PASSPHRASE
```

Now you can run the server using:

```sh
make server

# Sample Output:go1.13 run server.go
Within Go Server App
Starting server on: http://localhost:4000/
```

#### playing with the APIs

You should be able to access the APIs using:

##### listFiles

Once you start the server go to the following link: [listFiles](http://localhost:4000/v1/files)

##### getFile

Append one of the filenames mentioned in the listFiles url to get the required file. For example if `a.txt` is name of one of the files use this link: [getFile: a.txt](http://localhost:4000/v1/files/a.txt)
