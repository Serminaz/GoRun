/*N - количество работяг которые будут генерировать data*/
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
	wg := sync.WaitGroup{}      // объединяет (N)горутин в группу
	wg.Add(N) 
    m := new(sync.Mutex)   
	for i := 0; i < N; i++ {
	//	wg.Add(1)               // 1 цикл -> 1 горутина
		go func(m *sync.Mutex) {
			m.Lock()           // блокирует одну из N горутин
			AppendToArray()    // изменяет значения 
			m.Unlock()         // отпускает 
			defer wg.Done()    // элемент группы завершил свое выполнение
		}(m)
	}

	wg.Wait()                  // ожидаеt завершения всех(N) горутин в wg
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
	return m
}

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
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}



/*https://metanit.com/go/tutorial/7.7.php*/