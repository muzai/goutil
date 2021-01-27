package pkg

/*
#include <stdio.h>

void sayHello(){
    printf("Hello, XiaoC!");
}
*/
import "C"

func printc(){
    C.sayHello()
}
