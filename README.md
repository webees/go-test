##### Generate private key (.key)

```sh
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
```

##### Generation of self-signed(x509) public key (PEM-encodings `.pem`|`.crt`) based on the private (`.key`)

```sh
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

# Git Bash

```
> go env
> setx GOENV "D:\@Go\env"
> setx GOPATH "D:\@Go\GOPATH"
> go env -w GOPROXY=https://goproxy.cn,direct
> go env -w GOCACHE="D:\@Go\GOCACHE"
> go env -w GOMODCACHE="D:\@Go\GOMODCACHE"
> go env -w GO111MODULE=on
> go env -w CGO_ENABLED=0
```

```
curl -H "Content-Type: application/json" -X POST -d  '["t1","t2","t3"]' localhost:1024
```
