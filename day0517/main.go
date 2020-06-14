package main
import(
	"fmt"
	"strings"
)

func suffixImage(s string) func(string)string {
return func(path string) string {
	if strings.HasSuffix(path,s) {
		return path
	}
	return path + s
	
}
}

func addUper() func(int)int {
	base := 10
	return func(n1 int) int{
		res := base + n1
		base ++
		return res
	}
}


func main()  {
	f := addUper()
	fmt.Println("func res = ",f(1))
	fmt.Println("func res = ",f(1))
	fmt.Println("func res = ",f(1))
	fmt.Println("func res = ",f(1))
	fmt.Println("func res = ",f(1))
	f2 := suffixImage(".jpg")
	fmt.Println("image1 ",f2("1.jpg"))
	fmt.Println("image1 ",f2("1"))
	f2 = suffixImage(".png")
	fmt.Println("image1 ",f2("1"))
	anon := func(a int,b int)int{
		return a + b
	}

	annof := func() func(int,int) int{
		return func(a1 int, b1 int) int{
			return a1 + b1
		}
	}
	anon2 := annof()
	
	anon3 := func(param string)func(int,int) int {
		fmt.Println("param",param)
		return func(a1 int,b1 int) int{
			return a1 + b1
		}
	}("Hello World!")
	type returnFunctionType func(int,int) int
	anon4 := func() returnFunctionType {
		return func(a1 int, b1 int) int{
			return a1 + b1
		}
	}()
	fmt.Println("anon func =",anon(1,2))
	fmt.Println("anon func2 = ",anon2(1,2))
	fmt.Println("anon func3 = ",anon3(1,2))
	fmt.Println("anon func4 = ",anon4(1,2))
}