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

	// Title for specifying herald in flags
	Title string = "etcd"
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
func ProcessEtcdErrors(err error) error
```

#### func  Register

```go
func Register()
```
Register this herald with consul

#### type Etcd

```go
type Etcd struct {
	KeysAPI client.KeysAPI
}
```

Etcd structure stores the etcd clients KeysAPI

#### func  New

```go
func New() *Etcd
```
New returns a new Etcd struct

#### func (*Etcd) Declare

```go
func (e *Etcd) Declare(s *service.Service) error
```
Declare the service per the Declare interface in Herald.

#### func (*Etcd) Get

```go
func (e *Etcd) Get(s *service.Service) error
```
Get retrieves a service per the Declare interface in Herald

#### func (*Etcd) Init

```go
func (e *Etcd) Init() error
```
