package osearchclt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func makeUrl(dir string) string {
	if len(osclt.urls) <= 0 {
		return ""
	} else {
		return fmt.Sprintf("%s/%s", osclt.urls[(osclt.urlidx%uint16(len(osclt.urls)))], dir)
	}
}

func rqstPost(uri string, hdr map[string]string, data interface{}, respData interface{}) (http.Header, int, error) {
	// alloc
	sndData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(sndData))
	if err != nil {
		// increse urlidx
		osclt.urlidx++

		return nil, -1, err
	}

	// set auth
	if len(osclt.loginId) > 0 {
		req.SetBasicAuth(osclt.loginId, osclt.loginPwd)
	}

	// set headr
	//req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")

	// set header
	if len(hdr) != 0 {
		for k, v := range hdr {
			req.Header.Add(k, v)
		}
	}

	// send
	res, err := osclt.clt.Do(req)
	if err != nil {
		// increse urlidx
		osclt.urlidx++

		return nil, -1, err
	}

	// Close the connection to reuse it
	defer io.ReadAll(res.Body)
	defer res.Body.Close()

	// check result
	if (res.StatusCode / 100) != 2 {
		// increse urlidx
		osclt.urlidx++

		return res.Header, res.StatusCode, fmt.Errorf("not 2XX(%d)", res.StatusCode)
	}

	// unmarshal
	if respData != nil {
		_ = json.NewDecoder(res.Body).Decode(respData)
		//err = json.NewDecoder(res.Body).Decode(respData)
		// if err != nil {
		// 	return res.Header, res.StatusCode, err
		// }
	}

	return res.Header, res.StatusCode, nil
}
