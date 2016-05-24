// Copyright © 2015 Jason Smith <jasonrichardsmith@gmail.com>.
//
// Use of this source code is governed by the LGPL-3
// license that can be found in the LICENSE file.

package etcd

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"net/context"

	"github.com/coreos/etcd/client"
	"github.com/duckbunny/herald"
	"github.com/duckbunny/service"
)

var (
	etcdMachines string
	uname        string
	upass        string
	// Where the ServiceKVPath resides
	KVpath string = "services"

	// Title for specifying herald in flags
	Title string = "etcd"
)

func init() {
	flag.StringVar(&etcdMachines, "etcd-machines", os.Getenv("ETCD_MACHINES"), "The etcd machines.")
	flag.StringVar(&uname, "etcd-username", os.Getenv("ETCD_USER"), "The etcd username if secured.")
	flag.StringVar(&upass, "etcd-pass", os.Getenv("ETCD_PASS"), "The etcd password if secured.")
}

// Etcd structure stores the etcd clients KeysAPI
type Etcd struct {
	KeysAPI client.KeysAPI
}

// New returns a new Etcd struct
func New() *Etcd {
	return new(Etcd)
}

// Declare the service per the Declare interface in Herald.
func (e *Etcd) Declare(s *service.Service) error {
	js, err := json.Marshal(s)
	if err != nil {
		return err
	}
	key := FormattedKey(s)
	_, err = e.KeysAPI.Set(context.Background(), key, string(js), nil)
	if err != nil {
		return ProcessEtcdErrors(err)
	}
	return nil
}

// Get retrieves a service per the Declare interface in Herald
func (e *Etcd) Get(s *service.Service) error {
	key := FormattedKey(s)
	resp, err := e.KeysAPI.Get(context.Background(), key, nil)
	if err != nil {
		return ProcessEtcdErrors(err)
	}
	return json.Unmarshal([]byte(resp.Node.Value), s)
}

func ProcessEtcdErrors(err error) error {
	if err == context.Canceled {
		return errors.New("Context cancelled by another routine")
	} else if err == context.DeadlineExceeded {
		return errors.New("Context deadline exceeded")

		//Need clarification on this one
		//} else if cerr, ok := err.(*client.ClusterError); ok {
		// process (cerr.Errors)
	} else {
		return errors.New("Bad cluster endpoints, which are not etcd servers")
	}

}

func (e *Etcd) Init() error {
	config := client.Config{
		Endpoints: Machines(),
	}
	if uname != "" {
		config.Username = uname
		if upass != "" {
			config.Password = upass
		} else {
			return errors.New("Etcd username provided but no password")
		}
	}
	cl, err := client.New(config)
	if err != nil {
		return err
	}
	e.KeysAPI = client.NewKeysAPI(cl)
	return nil
}

func Machines() []string {
	return strings.Split(etcdMachines, "|")
}

// FormattedKey returns correctly formatted key of the service
func FormattedKey(s *service.Service) string {
	return fmt.Sprintf("/%v/%v/%v/%v/definition", KVpath, s.Domain, s.Title, s.Version)
}

// Register this herald with consul
func Register() {
	c := New()
	herald.AddDeclaration(Title, c)
}
