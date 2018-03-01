package main

import (
	"fmt"
	"math/rand"
	"time"
//	"sort"
//	"container/list"
)

type Player struct {
	UserId    int
	UserName  string
	HandCards []Card
	Zhuang    bool
}

var (
	nameList = []string{"刘德华", "张学友", "成龙", "周星驰", "张国荣", "李小龙", "皮智文"}
	cards    = []Card{}
	color    = []string{"万", "条", "筒", "字"}
	value    = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	canPlayCardPlayer Player
)

const ()

type Card struct {
	value byte
	color string
}

func InitCards() {
	for i := 0; i < len(color); i++ {
		for j := 0; j < len(value); j++ {
			if color[i] == "字" && value[j] > 7 {

			} else {
				for k := 0; k < 4; k++ {
					cards = append(cards, Card{value[j], color[i]})
				}
			}
		}
	}
}

func Init(num int) (players []Player) {
	l := len(nameList)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	lens := r.Perm(l)
	for i := 0; i < num; i++ {
		if i == 0 {
			p := Player{i + 1, nameList[lens[i]], []Card{}, true}
			players = append(players, p)
			canPlayCardPlayer = p
		} else {
			p := Player{i + 1, nameList[lens[i]], []Card{}, false}
			players = append(players, p)
		}
	}
	return players
}

func XiPai() {
	InitCards()
	newCards := []Card{}
	l := len(cards)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	lens := r.Perm(l)
	for i := 0; i < l; i++ {
		newCards = append(newCards, cards[lens[i]])
	}
	cards = newCards
	fmt.Println(cards)
}

func FaPai(players []Player) {
	for i := 0; i < len(players); i++ {
		p := &players[i]
		isZ := p.Zhuang
		if isZ {
			p.HandCards = append(p.HandCards, cards[0:14]...)
			cards = cards[14:]
		} else {
			p.HandCards = append(p.HandCards, cards[0:13]...)
			cards = cards[13:]
		}
	}
}

func ColorToInt(s string) int{
	var in int
	if s == "万"{
		in = 1
	}else if s == "条"{
		in = 2
	}else if s == "筒"{
		in = 3
	}else if s == "字"{
		in = 4
	}
	return in
}
type cardArray []Card

func (card cardArray) Len() int{
	return len(card)
}
func (card cardArray) Less(i,j int) bool{
	ci := ColorToInt(card[i].color)
	cj := ColorToInt(card[j].color)
	cii := ci*10 + int(card[i].value)
	cjj := cj*10 + int(card[j].value)
	return cii < cjj
}
func (card cardArray) Swap(i,j int) {card[i],card[j] = card[j],card[i]}

func CheckHu(cards []Card) bool{
	if len(cards)%3 != 2 {return false}
	return false
	
}

func Remove(slice []Card,elems ...Card) []Card{
	isInElems := make(map[Card]bool)
	for _,elem := range elems{
		isInElems[elem] = true
	}
	w := 0
	for _,elem := range slice{
		if !isInElems[elem] {
			slice[w] = elem
			w++
		}
	}
	return slice[:w]
}

func main() {
	
	var member int
	member = 3
	fmt.Printf("初始化玩家人数，人数：%v \n", member)
	ps := Init(member)
	fmt.Println("初始玩家成功，开始洗牌")
	XiPai()
	fmt.Println("发牌")
	FaPai(ps)
//	for _, p := range ps {
//		sort.Sort(cardArray(p.HandCards))
//		fmt.Println(p)
//	}
//	//	开始打牌
////	isOver := false
//	Contains()
}
