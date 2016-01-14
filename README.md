###Etcd

Is the etcd implementation for herald declare.

[![GoDoc](https://godoc.org/github.com/duckbunny/etcd?status.svg)](https://godoc.org/github.com/duckbunny/etcd)


# etcd
--
    import "github.com/duckbunny/etcd"




## Usage

```go
var (

	// Where the ServiceKVPath resides
	KVpath string = "services"
)
```

#### func  FormattedKey

```go
func FormattedKey(s *service.Service) string
```
FormattedKey returns correctly formatted key of the service

#### func  Machines

```go
func Machines() []string
```

#### func  ProcessEtcdErrors

```go
func ProcessEtcdErrors(error) error
```

#### type Etcd

```go
type Etcd struct {
	KeysAPI *client.KeysAPI
}
```


#### func  New

```go
func New() *Etcd
```

#### func (*Etcd) Declare

```go
func (e *Etcd) Declare(s *service.Service) error
```

#### func (*Etcd) Get

```go
func (e *Etcd) Get(s *service.Service) error
```

#### func (*Etcd) Init

```go
func (e *Etcd) Init() error
```
