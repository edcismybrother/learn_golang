package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	UserId int
	UserName string
	HandCards []Card
	Zhuang bool
}

var(
	nameList = []string{"刘德华","张学友","成龙","周星驰","张国荣","李小龙","皮智文"}
	cards = []Card{}
	color = []string{"万","条","筒","字"}
	value = []byte{1,2,3,4,5,6,7,8,9}
)

const(
)

type Card struct {
	value byte
	color string
}

func InitCards()([]Card){
	for i := 0;i<len(color);i++{
		for j := 0;j<len(value);j++{
			if color[i] == "字" && value[j]>7 {
			
			}else{
				cards = append(cards, Card{value[j],color[i]}) 
			}
		}
	}
	return cards
}


func Init(num int)([]Player) {
	l := len(nameList)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	lens := r.Perm(l)
	players := []Player{}
	for i := 0; i < num; i++ {
		if i == 0{
			p := Player{i+1,nameList[lens[i]],[]Card{},true}
			players = append(players, p)
		}else{
			p := Player{i+1,nameList[lens[i]],[]Card{},false}
			players = append(players, p)
		}
	}
	return players
}

func XiPai(){
	cards = InitCards()
	newCards := []Card{}
	l := len(cards)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	lens := r.Perm(l)
	for i := 0;i<l;i++ {
		newCards = append(newCards, cards[lens[i]])
	}
	cards = newCards
}

func FaPai(players []Player){
	fmt.Println(cards)
	for _,p :=  range players{
		isZ := p.Zhuang
		if isZ {
			for i:= 0;i<14;i++{
				p.HandCards = append(p.HandCards, cards[i])
				cards = cards[0:]
			}
		}else{
			for i:= 0;i<13;i++{
				p.HandCards = append(p.HandCards, cards[i])
				cards = cards[0:]
			}
		}
	}
}

func main() {
//	p := new(Player{UserId : 1,UserName : "刘德华",HandCards : []byte{11,11,11,12,12,12,13,13,13,14,14,14,15,15}}))
//	p := &Player{
//		UserId: 1,
//		UserName: "刘德华",
//		HandCards: []byte{11,11,11,12,12,12,13,13,13,14,14,14,15,15},
//	}
	ps := Init(3)
	fmt.Println(ps)
	XiPai()
	FaPai(ps)
	fmt.Println(ps)
//	name := p.UserName
//	cards := p.HandCards
//	fmt.Printf("玩家 %v 的手牌是: %v", name ,cards)
	
//	发牌
	
}

