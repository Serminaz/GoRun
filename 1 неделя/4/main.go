/*Реализовать все возможные способы остановки выполнения горутины.
что такое руна ?*/
package main
import (
"fmt"
"sync"
"time"
)
/*. Блокировки могут происходить при использовании
примитивов для блокировок или при чтении/записи в канал. */
    
func main() { 
  WithWG()    
  withChanel() 
  withTime()
  withMutex()
 
}


func WithWG () {
	var wg sync.WaitGroup 
    wg.Add(2)     // в группе две горутины
    work := func(id int) { 
        defer wg.Done()
        fmt.Printf("Горутина %d начала выполнение \n", id) 
        time.Sleep(2 * time.Second)
        fmt.Printf("Горутина %d завершила выполнение \n", id) 
   } 
   // вызываем горутины
   go work(1) 
   go work(2) 
   
   wg.Wait()        // ожидаем завершения обоих горутин
   fmt.Println("WG Горутины завершили выполнение") 
   fmt.Println("/****************/")
}


 func   withChanel() {
    Ch := make(chan string) // создаем канал, в который будем отправлять сообщения
    go printer(Ch)  // вызываем функцию асинхронно в горутине
    Ch <- "Горутина 1 начала работу "   //записываем в канал 
    Ch <- "Горутина 2 заблокирована"
    Ch <- "Горутина 1 завершила работу"
    close(Ch)  // закрываем канал
    //time.Sleep(100 * time.Millisecond) // и ждем, пока printer закончит работу
    fmt.Println("Chanel Горутины завершили выполнение ")
    fmt.Println("/****************/")
}

func printer(Ch chan string) {
    for msg := range Ch {  // читаем из канала, пока он открыт
        fmt.Println(msg)
    }
     fmt.Println("Горутина 2 разблокирована ")
};

 
func withTime() {
  fmt.Println("Горутина 1 начала работу ")
	go calculateSomething(1000)
 fmt.Println("Горутина 1 заблокирована")
 //Функция time.Sleep() также блокирует выполнение функции
time.Sleep(500*time.Millisecond) 
 fmt.Println("Time Горутины завершили выполнение ")
 fmt.Println("/****************/")
}

func calculateSomething(n int) {
  fmt.Println("Горутина 2 начала работу ")
	result := 0
	for i := 0; i <= n; i++ {
		result += i * 2
	}
  fmt.Println("Горутина 2 закончила работу ")
	
}


 func withMutex(){
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mu.Lock() // из 10, только одна получит доступ к счетчику,
                // остальные 9ть заблокированы и ждут своей очереди
                fmt.Println("Горутина %d начала работу ", counter+1)
			counter++
			mu.Unlock()
      fmt.Println("Горутина %d закончила работу ", counter)
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

 