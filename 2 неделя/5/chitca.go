/*Разработать программу, которая переворачивает подаваемую на ход строку 
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/
package main 
import "fmt"
func main() { 
 
        input()
}

func input () {

        var input, word, word2, ff, word5 string 
        var end int
        for  end != 2 { 
                fmt.Println("Введи 4  слова через интер")
                fmt.Scanln(&input)
                fmt.Scanln(&word2) 
                fmt.Scanln(&word) 
                fmt.Scanln(&word5) 
                // vers1
                output  :=   Reverse1(input)
                fmt.Println("version(1) ",output)

                // vers 2
                fmt.Println("version(2) ", Reverse2(word2))

                // vers 3
            
                 fmt.Println("version(3) ", Reverse3(word))
                 
                 //vers 4
                 fmt.Println("version(4) ", ReverseString(word5))


        fmt.Println("желаете продолжить ? (y/n)")
        fmt.Scanln(&ff) 
        if ff == "y" {
                end = 1
        } else {
                end = 2
        }
}
}

func Reverse1(s string) string {
        r := []rune(s)
        // i,j - начало и конец строки 
        for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
            r[i], r[j] = r[j], r[i]
        }
        return string(r)
    }

func Reverse2(s string) string {
        n := len(s)
        runes := make([]rune, n)
        for _, rune := range s {
            n--
            runes[n] = rune
        }
        return string(runes[n:])
    }

func Reverse3(s string) (result string) {
        for _,v := range s {
          result = string(v) + result
        }
        return 
      }

      // ReverseStrint переворачивает строку
func ReverseString(s string) (result string) {
	for _, val := range s {
		// result = Текущий элемент + предыдущий result
		result = string(val) + result
	}
	return
}