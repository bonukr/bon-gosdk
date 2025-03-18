package osearchclt_test

import (
	"testing"
	"time"

	"github.com/bonukr/bon-gosdk/osearchclt"
)

func TestInsert(t *testing.T) {
	var err error

	// init
	//clt, err := btosearch.NewClient([]string{"https://34.22.77.58:5601", "https://34.22.70.67:5601", "https://34.64.62.201:5601"}, "openserach", "okestro2018")
	//clt, err := btosearch.NewClient([]string{"https://34.64.61.230:5601"}, "openserach", "okestro2018")
	//clt, err := btosearch.NewClient([]string{"https://34.64.62.201:5601"}, "openserach", "okestro2018")
	//clt, err := btosearch.NewClient([]string{"https://34.22.77.58:9200", "https://34.22.70.67:9200", "https://34.64.62.201:9200"}, "openserach", "okestro2018")
	osearchclt.Init([]string{"https://172.10.40.240:30920"}, "admin", "admin")

	// info
	// if err = clt.PrintInfo(); err != nil {
	// 	t.Errorf("PrintInfo failed: %v", err)
	// }

	// insert
	mapData := map[string]int{"apple": 6, "lettuce": 7, "unixtime": int(time.Now().Unix())}
	err = osearchclt.Insert("oke-test-bwlee-metric1", mapData)
	if err != nil {
		t.Errorf("PrintInfo failed: %v", err)
	}
}
