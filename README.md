
### Какой самый эффективный способ конкатенации строк?
Ответ: Использование strings.Builder. Он позволяет конкатенировать строки без создания новых.
___
### Что такое интерфейсы, как они применяются в Go?
Ответ: Интерфейс это своего рода определение. Он определяет и описывает конкретные методы, которые должны быть у какого-то другого типа.
___
### Чем отличаются RWMutex от Mutex?
Ответ: Оба нужны для обеспечения безопасности при доступе к разделяемым данным из нескольких горутин, но RWMutex различает между собой операции записи и чтения. Позволяет нескольким горутинам одновременно выполнять операции чтения, но только одной горутине выполнять операцию записи. 
___
### Чем отличаются буферизированные и не буферизированные каналы?
Ответ: Отличие между буферизированными и не буферизированными каналами в Go заключается в том, как они управляют передачей данных между горутинами. Не буферизированные каналы не имеют внутреннего буфера для хранения данных. Отправка и получение происходят синхронно
___
### Какой размер у структуры struct{}{}?
Ответ: Пустая структура не занимает места в памяти 
___
### Есть ли в Go перегрузка методов или операторов?
Ответ: В Go нет поддержки перегрузки методов или операторов. Перегрузка методов означает, что можно иметь несколько методов с одним и тем же именем, но с разными параметрами или типами возвращаемых значений. Перегрузка операторов позволяет определить разные версии оператора для разных типов данных.
___
### В какой последовательности будут выведены элементы map[int]int?
```go
m[0]=1
m[1]=124
m[2]=281
```
Ответ: Порядок не гарантирован. Не стоит опираться на него при разработке.
___
### В чем разница make и new?
Ответ: make используется для создания слайсов, карт и каналов. make возвращает инициализированный тип, готовый к использованию после создания, а new – указатель на тип с его нулевым значением
___
### Сколько существует способов задать переменную типа slice или map?
Ответ: Литерал, make.
___
### Что выведет данная программа и почему?
```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```
Ответ: 1 и 1. Внутри update p - локальная копия указателя, так что изменения не применяются.
___
### Что выведет данная программа и почему?
```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```
Ответ: Мы получим deadlock, так как надо передать wg по указателю. Exit никогда не выведется, если этого не сделать.
___
### Что выведет данная программа и почему?
```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
Ответ: Выведется 0, потому что локальная область видимости if'a никак не влияет на переменную в области видимости функции. Если бы внутри if'а было не `n := 1` а `n = 1`, вывод был бы 2
___
### Что выведет данная программа и почему?
```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
Ответ: Слайс передаются по значению, но имеют ссылку на лежащий под капотом массив. При передаче слайса в функцию, он копируется, но ссылка на массив остаётся неизменна. Вывод: [100 2 3 4 5]
___
### Что выведет данная программа и почему?
```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
Ответ: В данном случае вывод будет `[b b a][a a]`, так как после вызова append создаётся новый слайс и исходный не изменяется
___