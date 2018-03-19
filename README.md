## Doc

### Clone project from github
```bash
$ mkdir -p $GOPATH/src/github.com/myrubapa/homework && cd "$_"
$ git clone https://github.com/myrubapa/homework .
$ chmod +x ./run.sh
```

### Install dependency
```bash
$ govendor sync
```

### Run test
```bash
$ ./run.sh test
```

### Run project
```bash
$ ./run.sh
```