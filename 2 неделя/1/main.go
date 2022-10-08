/* Дана последовательность температурных колебаний:
 -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
 Последовательность в подмножноствах не важна.*/

 package main

 
 import (
	 "fmt"
 )
 

 func main() {
	 plenty := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	 fmt.Println(defrag(plenty))
 }
 


 func defrag(arr [8]float32) map[int][]float32 {
	 temp := make(map[int][]float32)
 
	 for _, content := range arr {
 
		 key := int(content / 10)
		 if key < 0 {
			 key = (key - 1) * 10
		 } else {
			 key = key * 10
		 }
 
		 temp[key] = append(temp[key], content)
	 }
 
	 return temp
 }
 