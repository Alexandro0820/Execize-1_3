//Циклические сдвиги в GOLANG

package main

import (
	"fmt"
)

func PRINT_ARRAY(arr []int) {
	for i := range arr {
		fmt.Print(arr[i], " ")
	}
}
func LEFT(arr []int, size int, step int) {
	var temp int
	for i := 0; i < step; i++ {
		for j := 0; j < size-1; j++ {
			temp = arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = temp
		}
	}
}
func RIGHT(arr []int, size int, step int) {
	var temp int
	for i := 0; i < step; i++ {
		for j := size - 1; j > 0; j-- {
			temp = arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp
		}
	}
}

func main() {
	const size = 5
	arr := [size]int{1, 2, 3, 4, 5}
	fmt.Println(" ")
	fmt.Print("Ваш исходный массив: ")
	PRINT_ARRAY(arr[:])
	fmt.Println(" ")
	var step int
	fmt.Print("Введите кол - во шагов: ")
	fmt.Scan(&step)
	var change string
	fmt.Print("Введите направление сдвига (влево или вправо): ")
	fmt.Scan(&change)
	switch change {
	case "влево":
		fmt.Print("Направление - влево на ", step, " шаг(-а)")
		fmt.Println(" ")
		LEFT(arr[:], size, step)
		fmt.Print("Ваш новый массив: ")
		PRINT_ARRAY(arr[:])
	case "вправо":
		fmt.Print("Направление - вправо на ", step, " шаг(-а)")
		fmt.Println(" ")
		RIGHT(arr[:], size, step)
		fmt.Print("Ваш новый массив: ")
		PRINT_ARRAY(arr[:])
	default:
		fmt.Println("Неверный ввод!")
	}
	fmt.Println(" ")
}
