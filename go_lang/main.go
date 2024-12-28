package main

import (
	"fmt"
)

var (
  name_sport string = "Ivan"
  // age int = 25
  // height float32 = 12.2
  // c complex64 = complex(1, 2)
)

// func add(a complex64, b complex64) complex64 {
//   return a + b
// }


func main() {
  // fmt.Println(age, name_sport, height, c, real(c), imag(c))
  // fmt.Println(add(c, c))
  
  // for i:=0; i<10; i++{
  //   fmt.Print(add(c, c))
  // }
  // println()

// Array and slice
  nums := []int{1, 2, 3}
  // nums = append(nums, 10)
  // fmt.Println(nums[2])

  // nums_arr := [5]int{1, 2, 3}
  // fmt.Println(nums_arr)

  // arr := nums_arr[2:3]
  // fmt.Println(arr)
// При отправки slide в ф-ю изменяется slide в отличее от массива
  fmt.Println(nums)
  copy_nums := make([]int, len(nums))
  copy(copy_nums, nums)
  fmt.Println(reverseSlideInt(copy_nums))
  fmt.Println(nums)

  inter := []interface{}{"Sergey", 12, 12.5, []int{1,2,3}}
  // run := []rune{'a', 'b', 'c'}
  fmt.Println(inter[0])
// Struct
  person1 := Person{"Name", 12}
  fmt.Println(person1.Name)
  person1.conterAge()
  person1.printPerson()
// Map
  dcit_person := map[string]int{
    "Ivan": 12,
    "Alexey": 14,
  }
  
  // Пример перебора slise arr map and str 
  for index, num := range name_sport {
    fmt.Printf("%d %c\n", index, num)
  }

  fmt.Println(dcit_person["Ivan"])

  fmt.Println(20 / 7 + 1)
  fmt.Println(96)
}

func totalMoney(n int) int {
  var count int
  for weak:=1; weak<=n/7+1; weak++ {
      for dey:=1; dey<=7; dey++ {
          count++;
          if weak==n/7+1 && dey==(n*7)-((n/7)*7){
              fmt.Println((n*7)-((n/7)*7))
              break;
          }
        }
      }
      return count

      n := 20
      
}
// count = n // 7
// (28 * count) + ((7 * (count - 1)) / 2) * count + (((count + 1) + ((n - (count * 7)) + count)) / 2) * (n - (count * 7))

// Можно сказать новый тип, который состоит из других типов
type Person struct {
  Name string
  age int
}

func (p *Person) conterAge() {
  p.age++
}

func (p Person) printPerson() {
  fmt.Printf("Меня зовут %s и мне %d лет\n", p.Name, p.age)
} 


func reverseSlideInt(arr []int) []int{
    for i, j:=0, len(arr)-1; i < j; i,j = i+1, j-1{
        arr[i], arr[j] = arr[j], arr[i]
    } 

    return arr
}