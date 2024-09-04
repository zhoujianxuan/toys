package main

import "github.com/zhoujianxuan/toys/codetop/no3"

func main() {

}

type Nil struct {
}

var NilPoint = Nil{}

var codeMap = make(map[int32]Nil)

func initCodeTop() {
	codeMap[no3.GetNo()] = NilPoint
}
