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
}

func Contains(ls *list.List,card interface{})(bool, *list.Element){
	for le := ls.Front();le != nil;le = le.Next(){
		if le.Value == card{
			return true,le
		}
	}
	return false,nil
}

func Removes(ls *list.List,card interface{}){
	if contain,e := Contains(ls, card);contain {
		ls.Remove(e)
	}
}