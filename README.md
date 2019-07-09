# fileserver
A file server.

## Installation

### From binary

For Mac:

```bash
$ os=darwin arch=x86_64 bash <(curl -s https://raw.githubusercontent.com/broothie/fileserver/master/install.sh)
```

Available OSes:
- darwin
- linux
- windows

Available architectures:
- i386
- x86_64

Releases can also be downloaded directly from [here](https://github.com/broothie/fileserver/releases).

### From source

```bash
$ git clone https://github.com/broothie/fileserver.git
$ cd fileserver/
$ go build
$ cp fileserver /usr/local/bin/ # or wherever you want to put it...
```

### From source (go style)

```
$ go get -u github.com/broothie/fileserver
```

Make sure `$GOPATH/bin/` is on your `$PATH`.
