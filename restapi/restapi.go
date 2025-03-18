package restapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

var _dft_conn_timeout_msec = 7000
var _dft_send_timeout_msec = (10 * 1000)
var clt *http.Client

func SetTimeout(connTimeoutMsec, sendTimeoutMsec int) {
	_dft_conn_timeout_msec = connTimeoutMsec
	_dft_send_timeout_msec = sendTimeoutMsec
}

func init() {
	// option
	clt = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable security check for a client: (x509) certificate signed by unknown authority
			Dial: (&net.Dialer{
				Timeout: time.Duration(_dft_conn_timeout_msec) * time.Millisecond,
			}).Dial,
			//TLSHandshakeTimeout: time.Duration(_dft_conn_timeout_msec) * time.Millisecond,
		},
		Timeout: time.Duration(_dft_send_timeout_msec) * time.Millisecond,
	}
}

func RqstGet(uri string, hdr map[string]string, param map[string]string, respData interface{}) (http.Header, int, error) {
	// set param
	if len(param) > 0 {
		urlA, _ := url.Parse(uri)
		values := urlA.Query()
		for k, v := range param {
			values.Add(k, v)
		}
		urlA.RawQuery = values.Encode()
	}

	// alloc
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		//log.Printf("%s: http.NewRequest error: url(%s): %s", fnc, reqUrl, err)
		return nil, -1, err
	}

	// set header
	for k, v := range hdr {
		req.Header.Add(k, v)
	}

	// send
	res, err := clt.Do(req)
	if err != nil {
		//log.Printf("%s: http.PostForm error: url(%s): %s", fnc, url, err)
		return nil, -1, err
	}

	// close th connection to reuse it
	defer io.ReadAll(res.Body)
	defer res.Body.Close()

	// check result
	if (res.StatusCode / 100) != 2 {
		//log.Printf("%s: resp.StatusCode != 2xx: %s: \n%s", fnc, res.Status, res.Body)
		return res.Header, res.StatusCode, errors.New(fmt.Sprintf("not 2XX(%d)", res.StatusCode))
	}

	// unmarshal
	if respData != nil {
		err = json.NewDecoder(res.Body).Decode(respData)
		if err != nil {
			//log.Printf("%s: json.Decode error: %s \n%s", fnc, res.Status, res.Body)
			//return res.Header, res.StatusCode, err
		}
	}

	return res.Header, res.StatusCode, nil
}

func RqstPost(uri string, hdr map[string]string, data interface{}, respData interface{}) (http.Header, int, error) {
	// alloc
	sndData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(sndData))
	if err != nil {
		//log.Printf("%s: http.NewRequest error: url(%s): %s", fnc, reqUrl, err)
		return nil, -1, err
	}

	// set header
	for k, v := range hdr {
		req.Header.Add(k, v)
	}

	// send
	res, err := clt.Do(req)
	if err != nil {
		//log.Printf("%s: http.PostForm error: url(%s): %s", fnc, url, err)
		return nil, -1, err
	}
	defer res.Body.Close()

	// check result
	if (res.StatusCode / 100) != 2 {
		//log.Printf("%s: resp.StatusCode != 2xx: %s: \n%s", fnc, res.Status, res.Body)
		return res.Header, res.StatusCode, errors.New(fmt.Sprintf("not 2XX(%d)", res.StatusCode))
	}

	// unmarshal
	if respData != nil {
		err = json.NewDecoder(res.Body).Decode(respData)
		if err != nil {
			//log.Printf("%s: json.Decode error: %s \n%s", fnc, res.Status, res.Body)
			//return res.Header, res.StatusCode, err
		}
	}

	return res.Header, res.StatusCode, nil
}
