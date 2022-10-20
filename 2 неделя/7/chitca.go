/*Реализовать собственную функцию sleep.*/

package main

import (
	"fmt"
	"log"
	"time"
)

func Sleep1(sec int) {
	<-time.After(time.Second * time.Duration(sec))
}
func main() {
	log.Println(" begin")
	fmt.Printf("Заснул в - %v, время сна - %d\n", time.Now())
	Sleep1(3)
	log.Println(" end")
}

/*package main

import (
"fmt"
"time"
)

func sleepFunc(d time.Duration) {
	timer := time.NewTimer(d) //таймер ждет окончание отсчета и завершает выполнение функции

	<-timer.C

	return
}

func main() {
	d := time.Duration(3 * time.Second) //засыпает на 3 секунды

	fmt.Printf("Заснул в - %v, время сна - %v\n", time.Now(), d)

	sleepFunc(d)

	fmt.Printf("Проснулся в - %v", time.Now())
}
*/
