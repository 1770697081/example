package main

import( 
	"fmt" 
)


func main(){
    a := 1
    b := 2
    defer calc(a, calc(a,b))    //先执行calc(1,2) p="1,2,3" ,  1,3 4
	
	a = 0
    defer calc(a, calc(a,b))     //1、先执行这个  calc(0,2) p="0,2,2"     calc(0,2) p = 0 ,2,2
}

func calc(x,y int)int{
    fmt.Println(x,y,x+y)
    return x+y
}

/*
1 2 3
0 2 2
0 2 2
1 3 4
考察两个知识点：
1.defer是栈调用，后写的先执行
2.defer的函数调用语句会在父函数调用后执行，但是用到的参数会在当时就执行得出
*/