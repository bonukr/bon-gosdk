package osearchclt

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

var _dft_timeout_msec = 3000
var _dft_use_https_verify = false
var _dft_max_idle_conns_per_host = 3

var osclt struct {
	clt    *http.Client
	urls   []string
	urlidx uint16

	loginId  string
	loginPwd string
}

func Init(urls []string, id, pwd string) {
	osclt.clt = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !_dft_use_https_verify,
			},
			MaxIdleConnsPerHost: _dft_max_idle_conns_per_host,
		},
		Timeout: time.Millisecond * time.Duration(_dft_timeout_msec),
	}
	osclt.urls = urls
	osclt.urlidx = 0
	osclt.loginId = id
	osclt.loginPwd = pwd
}

func Insert(index string, document interface{}) error {
	url := makeUrl(fmt.Sprintf("%s/_doc", index))
	_, rescd, err := rqstPost(url, nil, document, nil)
	if err != nil {
		log.Printf("Insert failed: url(%s): rescd(%d): %v", url, rescd, err)
	}

	return err
}

func InsertRack(document *Msg) error {
	return Insert(IndexName.Rack, document)
}

func InsertServer(document *Msg) error {
	return Insert(IndexName.Server, document)
}

func InsertStorage(document *Msg) error {
	return Insert(IndexName.Storage, document)
}

func InsertOpenstack(document *Msg) error {
	return Insert(IndexName.Openstack, document)
}
