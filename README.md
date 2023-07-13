# Hello dagger

### Requires
  - Above Go version 1.20
  - OCI container environment
  - dagger

### steps
1. Init source code
``` bash
$ mkdir src 
$ cd src
```
``` bash
$ go mod init hello_dagger
```
- Write a source code

2. Init CI code
``` bash
# Must move a project root directory
$ go mod init ci
```
``` bash
# Install degger
$ go get dagger.io/dagger@latest
```

- Write a ci code

3. Add Github Action
``` bash
# make directories
$ mkdir .github
$ mkdir .github/workflows/
```
- Write a github action yaml
