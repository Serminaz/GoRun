package main

import  (
	"fmt"
	"math/rand"
	 "os"
	"strconv"
	"sync"
	 "time"
	"github.com/joho/godotenv"
	//"crypto/rand"
         )
type Workers struct {
	name  string
	age   uint8
	high uint8
}

var arr = []Workers{}

func main() {
	N, err := strconv.Atoi(ReadingENV("N")) // для конвертации из строки to int
	if err != nil {
		fmt.Println("Нет данных о кол-ве работников!")
	}
	wg := sync.WaitGroup{} // объединяет горутины в группы 
	m := new(sync.Mutex)
	for i := 0; i < N; i++ {
		wg.Add(1)      // в группе 1 горутина
		go func(m *sync.Mutex) {
			m.Lock()
			AppendToArray()
			m.Unlock()
			defer wg.Done()
		}(m)
	}

	wg.Wait()     // ожидаем завершения обоих горутин
	fmt.Println(arr)
}
func init() {
	rand.Seed(time.Now().UnixNano())
	//rand.Int(rand.Reader, big.NewInt())
}

func AppendToArray() {
	arr = append(arr, Workers{
		name: randomName(),
		age:  uint8(randomAge()),
		high: uint8(randomHihgt()),
	})
}


func randomAge() int {
	n := rand.Intn(60)
	return n
}
func randomHihgt() int {
	m := 150 + rand.Intn(190-150)
	h :=  m  
	return h
}
// func randomHihgt() int {
// 	min := 150
// 	max := 190
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	h:= min + rand.Intn(max-min)
// 	return h
// }

func   randomName() string {
	rName := []string {"Eva", "Alica", "Sergey", "Harry Potter","Pupkin", "Sasha" }
	arr := len(rName)
	rArr := rand.Intn(arr) 
	value := rName[rArr]
	return value 
}
func ReadingENV(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Нет файла .env")
	}
	return os.Getenv(key)
}



/*https://metanit.com/go/tutorial/7.7.php*/