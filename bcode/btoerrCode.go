package bcode

//NOTE - 수집기 로그 코드

/*
- 수집 성공
- 수집 실패
- 수집기 함수 실행 성공
- 수집기 함수 실행 실패
- 내부 요청 오류
- 공급자 통신 오류
- 공급자 데이터 요청 성공
- 공급자 데이터 요청 실패
- logstash 전송 성공
- logstash 전송 실패
*/

type btlogcode struct {
	CollectSucess         string //수집 성공
	CollectFail           string //수집 실패
	CollectorfuncSucess   string //수집기 함수 실행 성공
	CollectorfuncFail     string //수집기 함수 실행 실패
	InternalServerErr     string //내부 요청 오류
	ProviderConnectErr    string //공급자 통신 오류
	ProviderDataApiSucess string //공급자 데이터 요청 성공
	ProviderDataApiFail   string //공급자 데이터 요청 실패
	DataPiplineReqSucess  string //logstash 전송 성공
	DataPiplineReqFail    string //logstash 전송 실패
}

var Btlogcode = btlogcode{
	CollectSucess:         "BC0001",
	CollectFail:           "BC0002",
	CollectorfuncSucess:   "BC0003",
	CollectorfuncFail:     "BC0004",
	InternalServerErr:     "BC0005",
	ProviderConnectErr:    "BC0006",
	ProviderDataApiSucess: "BC0007",
	ProviderDataApiFail:   "BC0008",
	DataPiplineReqSucess:  "BC0009",
	DataPiplineReqFail:    "BC0010",
}
