// hello world v2.1.2-release by 冬冬
package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// MAGIC GIRL
const MAGIC = 8579384335739694
const GIRL = "3bX=<* \\"

func main() {
	elPsyCongroo(GIRL)
}

// elPsyCongroo 233333.
func elPsyCongroo(TAT string) {
	OAO := []byte(TAT)
	oAo := ""
	for _, qwq := range OAO {
		oAo += strconv.Itoa(int(qwq))
	}
	_233, _ := strconv.Atoi(oAo)
	_2333 := len(fmt.Sprint(MAGIC))
	_23333, _ := strconv.Atoi(strconv.FormatInt(int64(_2333), 16))

	el := MAGIC ^ _233
	psy := el / int(math.Pow(float64(_23333), float64(_2333>>1)))
	congroo := ""
	rand.Seed(int64(psy))
	for QAQ := 0; QAQ < _23333; QAQ++ {
		if QAQ == _23333>>1 {
			congroo += " "
			rand.Seed(int64(el%psy + psy))
		}
		congroo += string(rand.Int()%'!' + 87)
	}
	// el psy congroo
	fmt.Print(congroo)
}
