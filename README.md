##  Functions
### map Functions
```golang

var strs = []string{"peach", "apple", "pear", "plum"}

InArray(needle interface{}, hystack interface{}) 判断needle是否在hystack这个slice中，返回bool
ArrayUnique(arr []string) []string  去重
Index(vs []string, t string) int  返回t在vs中的下标位置，返回int
Any(vs []string, f func(string) bool) bool   遍历切片vs,的每个值，并作为参数传递给函数f调用，只要有一个条件成立，就返回true

fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
```
