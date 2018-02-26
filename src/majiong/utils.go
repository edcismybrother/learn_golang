package main

import (
	"container/list"
	"fmt"
)

func Contain1s(){
	l := list.New()
	l.PushBack("one")
	l.PushBack("two")
	for le := l.Front();le != nil;le = le.Next(){
		fmt.Println(le.Value)
	}
func Contains(ls *list.List,value ){
	l := list.New()
	l.PushBack("one")
	l.PushBack("two")
	for le := l.Front();le != nil;le = le.Next(){
		fmt.Println(le.Value)
	}
}