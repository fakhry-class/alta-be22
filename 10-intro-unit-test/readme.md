# Unit testing

### go.mod
```bash
go mod init nama-module
```

### Menjalankan file test
```bash
go test ./calculate/... -cover
```
atau
```bash
go test ./calculate/... -coverprofile=cover.out && go tool cover -html=cover.out
```

### download testify
```bash
go get github.com/stretchr/testify
```