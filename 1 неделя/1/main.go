package main

import  (
	"fmt"
    "time"
)
func main() {

	power()

}

func power() {
	//start := time.Now()
	arr :=[]int {7,12,5}
	 high := 0
	 for _, sc := range arr {
		high = sc*sc  
		go func (high int)  {
			fmt.Println(high)
			}(high)
		}
	// elapsedTime := time.Since(start)
	// fmt.Println("Total Time For Execution: " + elapsedTime.String())
	time.Sleep(time.Millisecond)
}
