/*Реализовать конкурентную запись данных в map.*/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
var wg sync.WaitGroup
    m1 := map[string]string{
		  "typica": "Ramazan",
		}
		 rName := []string {"Сын маминой подруги", "lilia", "Sergey", "Harry","Pupkin", "Sasha" }
		 Name := []string {"делает 300к/с", "Zver", "love hachimuchi", "Учится в хогвартсе","15", "душнила" }
		 u1 := len(rName)

		 for i := 0; i < 5; i++ {
			
		 u := rand.Intn(u1)
		 value1 := rName[u]
		 key1:= Name[u]
		  
		 go func() {
			wg.Add(1)
			defer wg.Done()
			for key, value := range m1 { // Порядок не определен 
			time.Sleep(time.Millisecond*15)
				m1[key] = value
				m1[key1] = value1
				}
		   }()
	     }
		 wg.Wait()
		 fmt.Println("конечный", m1)
}

