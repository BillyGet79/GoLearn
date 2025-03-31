# Go语言加强

## 指针

### 概念

先来看一段代码

```go
func main() {
	var a int = 10
	var p *int = &a
	a = 100
	fmt.Println("a = ", a)
	*p = 250	//借助a变量的地址，操作a对应空间
	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)
    a = 1000
	fmt.Println("*p = ", *p)
}
```

输出结果如下：

```cmd
a =  100
a =  250
*p =  250
*p =  1000
```

指针就是地址。指针变量就是存储地址的变量。

*p: 解引用、间接引用

图解：

<img src="img/image-20250330194310441.png" alt="image-20250330194310441" style="zoom: 80%;" />

### 栈帧

用来给函数运行提供内存空间。 取内存于 stack 上。

当函数调用时，产生栈帧。函数调用结束时，释放栈帧。

栈帧存储：1.局部变量	2.形参（形参与局部变量存储地位等同）	3.内存字段描述值

为了方便描述，我们在上面的代码基础上做一些修改

```go
func test(m int) {
	var b int = 100
	b += m
	fmt.Println(b)
}

func main() {
	var a int = 10
	var p *int = &a
	a = 100
	fmt.Println("a = ", a)
	test(10)
	*p = 250
	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)
	a = 1000
	fmt.Println("*p = ", *p)
}
```

下面是整个程序的运行图：

<img src="img/image-20250330200051665.png" alt="image-20250330200051665" style="zoom:80%;" />

在运行main函数时，会开辟一个空间供main函数存储局部变量，并且栈顶指针向下移动。在代码运行到`test(10)`之后，会开辟一个空间供test函数存储局部变量和形参，然后依旧向下移动指针。当test函数运行完毕后，删除这个栈帧，然后将栈顶指针向上移动，移动回main函数。

### 指针使用注意

空指针：违背初始化的指针。例如`var p *int`，此时`*p --> err`

野指针：被一片无效的地址空间初始化

野指针例子如下：

```go
func main() {
	var p *int = 0
	fmt.Println(*p)
}
```

此时运行会报出如下异常：

```cmd
# command-line-arguments
.\01-指针的应用.go:25:15: cannot use 0 (untyped int constant) as *int value in variable declaration
```

此时p指向的是一个无效的地址空间，编译时一定会报错

但是现在我们并不想通过定义一个变量然后再用指针指向它，而是想定义一个指针，然后让指针指向这个新开辟的内存空间。这个时候我们可以使用`new`关键字

示例如下：

```go
func main() {
	/*	var a int = 10
		fmt.Println("&a", &a)*/
	var p *int

	//在heap上申请一片内存地址空间
	p = new(int) //默认类型的默认值
	fmt.Println(*p)
}
```

这个时候指针就不是空指针和野指针了，也可以用语法来修改这个变量。

### 指针的函数传参

1. 传地址（引用）：将形参的地址值作为函数参数传递
2. 传值（数据）：将实参的值拷贝一份给形参
3. 传引用：在A栈帧内部，修改B栈帧中的变量值

这个在基础语法部分就有讲过（当然其他语言也具有这个特性）

举一个例子来说明：

```go
func swap(a int, b int) {
	a, b = b, a
	fmt.Println(a, b)	//20，10
}

func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Println(a, b)	//10，20
}
```

上述是典型的值传递，所以打印结果也显而易见。

但是如果我们使用地址传递，则可以彻底交换两个值。

```go
func swap(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b)
}
```

上述就是地址传递，两个参数传递的都是地址。

## 切片

### 为什么用切片

1. 数组的容量固定，不能自动扩展
2. 值传递。数组作为函数参数时，将整个数组值拷贝一份给形参

在Go语言当中，我们几乎可以在所有场景中，使用切片替换数组使用

**切片的本质**：不是一个数组的指针，是一种数据结构体，用来操作数组内部元素。

看一下切片的底层结构体定义：

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

第一个`array`类型为`unsafe.Pointer`类型，指向的是具体的数组，`len`表示当前数组长度，`cap`表示当前数组容量。

### 切片的使用

先看一段代码

```go
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	//[]内的三个属性：low,high,max
	//low：起始下标位置
	//high：结束下标位置，len=high-low
	//容量：cap=max-low
    //low（high）都可以省略，默认在从0开始（到末尾结束）
	s := arr[1:3:5]
	fmt.Println(s)
	fmt.Println(cap(s))
	fmt.Println(len(s))
}
```

输出结果如下：

```cmd
[2 3]
4
2
```

上述是最基本的切片的使用，当然，切片可以去切取数组。

请注意下述切片使用场景：

```go
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[2:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))

	s2 := s[2:7]
	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2))
	fmt.Println("cap(s2) = ", cap(s2))
}
```

此时你猜猜输出结果是什么？

```go
s =  [3 4 5]
len(s) =  3
cap(s) =  8
s2 =  [5 6 7 8 9]
len(s2) =  5
cap(s2) =  6
```

对于这个结果，我们可以从切片的结构体定义方面来理解：

- 当定义`s`时，由于其`low=2`，所以它指向数组下标为2的位置，将其视为起始位置，然后设置`len`和`cap`，为3和8
- 当定义`s2`时，由于其`low=2`，所以它指向`s`的起始位置后面的的2的位置，将其视为起始位置，然后设置`len`和`cap`，为5和6

上述两个操作都没有指定`max`，而指向的都是`arr`数组，所以可以理解为，我们这样使用切片只是使用了一个视图，并未真正在底层创建一个空间存储这些内容。

由于没有指定`max`，所以上述我们并未发现在指定切片时的越界行为。但是请注意，当我们定义切片，在切取部分内容时，一定要注意指定的数组或切片的容量，如果超出了一定会报异常。

### 创建切片

看代码：

```go
func main() {
	//1.自动推导赋初值
	s1 := []int{1, 2, 2, 3}
	fmt.Println("s1 = ", s1)
	//2.make创建
	//第一个参数指定类型
	//第二个参数指定len
	//第三个参数指定cap
	s2 := make([]int, 5, 10)
	fmt.Println("s2 = ", s2)
	//3.make创建不指定cap
	//此为最常用的方式
	s3 := make([]int, 10)
	fmt.Println("s3 = ", s3)
}
```

输出结果如下：

```cmd
s1 =  [1 2 2 3]
s2 =  [0 0 0 0 0]
s3 =  [0 0 0 0 0 0 0 0 0 0]
```

注意：make只能创建slice、map和channel，并且返回一个有初始值（非0）的对象

当切片做函数参数时，为传引用方式（传地址）

### append使用

看代码

```go
func main() {
	s1 := []int{1, 2, 2, 3}
	//append(切片对象， 待追加元素）
	//返回值类型为切片类型
	s1 = append(s1, 88)
	s1 = append(s1, 88)
	s1 = append(s1, 88)
	s1 = append(s1, 88)
	s1 = append(s1, 88)
	fmt.Println("s1:", s1)
}
```

输出结果：

```cmd
s1: [1 2 2 3 88 88 88 88 88]
```

添加元素必然会涉及到扩容机制。

向切片增加元素时，切片的容量会自动增长。1024以下时，以两倍方式增长。

练习题：

```go
/**
练习：给定一个字符串列表，在原有slice上返回不包含空字符串的列表，如：
{"red", "", "black", "", "", "pink", "blue"}
--> {"red", "black", "pink", "blue"}
 */
func main() {
	slice := []string{"red", "", "black", "", "", "pink", "blue"}
	var res []string
	for _, s := range slice {
		if s != "" {
			res = append(res, s)
		}
	}
	fmt.Println(res)
}
```

```go
/**
练习：写一个函数，消除[]string中重复字符串，如：
{"red", "black", "red", "pink", "blue", "pink", "blue"}
--> {"red", "black", "pink", "blue"}
*/
func main() {
	slice := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	var res []string
	set := make(map[string]struct{})
	for _, v := range slice {
		if _, exists := set[v]; !exists {
			set[v] = struct{}{}
			res = append(res, v)
		}
	}
	fmt.Println(res)
}
```

### copy使用

```go
func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s1 := data[8:]
	s2 := data[0:5]
	//copy(目标位置切片, 源切片)
	//拷贝过程中，直接对应位置拷贝
	copy(s2, s1)
	fmt.Println("s2 = ", s2)
	fmt.Println(data)
}
```

输出结果如下：

```cmd
s2 =  [8 9 10 11 12]
[8 9 10 11 12 5 6 7 8 9 10 11 12 13 14 15]
```

看一道练习题：

```go
/*
练习：要删除slice中间的某个元素并保存现有的元素顺序。如：
{5, 6, 7, 8, 9} --> {5, 6, 8, 9}
*/
func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	data = deleteSliceElem(data, 5)
	fmt.Println(data)
}

func deleteSliceElem(data []int, i int) []int {
	slice1 := data[i:]
	slice2 := data[i+1:]
	copy(slice1, slice2)
	return data[:len(data)-1]
}
```

输出结果：

```go
[0 1 2 3 4 6 7 8 9 10 11 12 13 14 15]
```

## map

### map的创建与初始化

#### 创建

```go
func main() {

	//1.var定义
	var m1 map[int]string
	fmt.Println(m1 == nil)
	//此定义方法定义后并不能存储键值对，因为其变量是一个nil，即为空
	//m1[100] = "tiny" //err

	//2.赋值定义
	m2 := map[int]string{}
	fmt.Println(len(m2))
	//此时就可以插入值了
	m2[4] = "red"
	fmt.Println("m2 = ", m2)
	//3.make定义
	m3 := make(map[int]string)
	fmt.Println(len(m3))
	fmt.Println("m3 = ", m3)
	//4.make定义指定大小
	m4 := make(map[int]string, 10)
	fmt.Println(len(m4))
	fmt.Println("m4 = ", m4)
}
```

输出结果如下：

```cmd
true
0
m2 =  map[4:red]
0
m3 =  map[]
0
m4 =  map[]
```

#### 初始化

一般来讲，初始化语法用不上，这里展示两种方式

```go
	//初始化map
	//1.定义同时初始化
	var m5 map[int]string = map[int]string{1: "Luffy", 250: "Sanji", 130: "Zero"}
	fmt.Println("m5 = ", m5)
	//2.自动推导类型
	m6 := map[int]string{1: "Luffy", 250: "Sanji", 130: "Zero"}
	fmt.Println("m6 = ", m6)
```

### map赋值

看代码

```go
func main() {
	m1 := make(map[int]string, 1)
	m1[100] = "Nami"
	m1[20] = "Hello"
	m1[3] = "World"
	fmt.Println("m1:", m1)

	m1[3] = "yellow" //成功，将原map中key值为3的map元素替换
	fmt.Println("m1:", m1)
}
```

输出结果

```cmd
m1: map[3:World 20:Hello 100:Nami]
m1: map[3:yellow 20:Hello 100:Nami]
```

赋值过程中，如果新map元素的key与原map元素key相同，则替换value；如果不同，则添加

### map的使用

#### 遍历map方式

```go
	m1 := make(map[int]string, 1)
	m1[100] = "Nami"
	m1[20] = "Hello"
	m1[3] = "World"
	//遍历map
	for k, v := range m1 {
		fmt.Println("key:", k, "value:", v)
	}
```

输出结果：

```cmd
key: 20 value: Hello
key: 3 value: World
key: 100 value: Nami
```

#### 判断key是否存在

```go
	//判断map中的key是否存在
	if value, exists := m1[3]; !exists {
		fmt.Println("value=", value, "exists=", exists)
	} else {
		fmt.Println("value=", value, "exists=", exists)
	}
```

输出结果：

```cmd
value= World exists= true
```

### map删除和传参

```go
func main() {
	m1 := make(map[int]string, 1)
	m1[100] = "Nami"
	m1[20] = "Hello"
	m1[3] = "World"

	//map做函数参数、返回值，传引用
	m1 = mapDelete(m1, 20)
	fmt.Println(m1)
}

func mapDelete(m map[int]string, key int) map[int]string {
	//该方法没有任何返回值，直接删除使用即可
	delete(m, key)
	return m
}
```

输出结果：

```go
map[3:World 100:Nami]
```

### 练习

```go
/*
练习：封装wcFunc()函数。接收一段英文字符串str。返回一个map，记录str中每个"词"出现的次数
如："I love you and I love my family too"
*/
func main() {
	str := "I love my work and I love my family too"
	m := wcFunc(str)
	fmt.Println(m)
}

func wcFunc(str string) map[string]int {
	slice := strings.Split(str, " ")
	res := make(map[string]int)
	for _, s := range slice {
		if value, exists := res[s]; exists {
			res[s] = value + 1
		} else {
			res[s] = 1
		}
	}
	return res
}
```

输出结果

```cmd
map[I:2 and:1 family:1 love:2 my:2 too:1 work:1]
```

