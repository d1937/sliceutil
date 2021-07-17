##  Functions
### map Functions
```golang

var strs = []string{"peach", "apple", "pear", "plum"}

InArray(needle string, hystack string) bool 判断needle是否在hystack这个slice中，返回bool

ArrayUnique(arr []string) []string  去重

Index(vs []sting, t string)  返回t在vs中的下标位置，返回int

Any(vs []string, f func(string) bool) bool   遍历切片vs,的每个值，并作为参数传递给函数f调用，只要有一个条件成立，就返回true

fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
})) // output: true


All(vs []string, f func(string) bool) bool  遍历切片vs,的每个值，并作为参数传递给函数f调用，全部条件都为true，则返回true,否则返回false

fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
})) // output: false

Filter(vs []string, f func(string) bool) []string  遍历切片vs,的每个值，并作为参数传递给函数f调用,返回符合条件的值

Map(vs []string, f func(string) string) []string  遍历切片vs,的每个值，并作为参数传递给函数f调用并返回
```
