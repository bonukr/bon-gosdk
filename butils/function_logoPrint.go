package butils

import (
	"log"

	"github.com/common-nighthawk/go-figure"
)

// // 각 프로젝트 소스코드 내에서 프로젝트명 가져오는 코드
// func packageName() string {
// 	pc, _, _, _ := runtime.Caller(1)
// 	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
// 	pl := len(parts)
// 	pkage := ""
// 	if parts[pl-2][0] == '(' {
// 		pkage = strings.Join(parts[0:pl-2], ".")
// 	} else {
// 		pkage = strings.Join(parts[0:pl-1], ".")
// 	}
// 	packageNames := strings.Split(pkage, "/")
// 	return packageNames[0]
// }

func LogoPrint(pName string) {
	serviceLogo := `` + pName
	log.Println(figure.NewFigure(serviceLogo, "doom", true))
}
