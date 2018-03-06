package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
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
	canPlayCardPlayer *Player
	end bool =  true
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
			canPlayCardPlayer = &p
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
func IntToColor(in int) string{
	var s string
	if in == 1{
		s = "万"
	}else if in == 2{
		s = "条"
	}else if in == 3{
		s = "筒"
	}else if in == 4{
		s = "字"
	}
	return s
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

/**下一个出牌的玩家**/
func nextPlayer(players *[]Player,actionPlayer *Player) *Player{
	if actionPlayer == nil{
//		pp :=  *players
		return &(*players)[0]
	}else{
		for k,p := range *players{
			if p.UserId == actionPlayer.UserId {
				index := (k+1)%len(*players)
				return &(*players)[index]
			}
		}
	}
	return nil
}

func intToCard(number int) Card{
	value := byte(number%100)
	color := IntToColor(number/100)
	return Card{value , color}
}

/**检查出牌是否存在于手牌中**/
func checkOutCard(cards *[]Card,outCard Card) bool{
	for i := 0;i<len(*cards);i++{
		
		if(*cards)[i] == outCard{
			*cards = append((*cards)[:i], (*cards)[i+1:]...)
			return true
		}
	}
//	for _,card := range &cards{
//		if card == outCard{
//			cards
//			return true
//		}
//	}
	return false
}

func checkHu(players *[]Player,card Card) bool{
	return false
}
func checkPeng(players *[]Player,card Card) (bool,*Player){
	for i := 0; i<len(*players);i++{
		handCards := (*players)[i].HandCards
		num := 0
		for _,c := range handCards{
			if c == card {num++}
		}
		if num>=2 {return true,&((*players)[i])}
	}
	return false,nil
}
func checkGang(players *[]Player,card Card) bool{
	return false
}

func checkAction(needHu bool,needPeng bool,needGang bool,players *[]Player,card Card){
	if needHu{
		end := checkHu(players,card)
		if end{return}
	}else if needPeng{
		end,p := checkPeng(players,card)
		if end{
			fmt.Println(p)
			return
		}
	}else if needGang{
		end := checkGang(players,card)
		if end{return}
	}
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
	for _, p := range ps {
		sort.Sort(cardArray(p.HandCards))
		fmt.Println(p)
	}
	//	开始打牌
//	isOver := false
//	Contains()
	var input int
	for end == true{
		canPlayCardPlayer = nextPlayer(&ps, canPlayCardPlayer)
		fmt.Printf("玩家 ： %v 请出牌", canPlayCardPlayer.UserName)
		a :  _,err := fmt.Scan(&input)
		if err != nil{
			fmt.Println(err)
			return
		}else{
			outCard := intToCard(input)
			if checkOutCard(&(canPlayCardPlayer.HandCards), outCard) == false{
				fmt.Println("该outcard不存在于手牌中")
				goto a
			}
			fmt.Printf("玩家：%v 打出 %v\n手牌为：%v\n", canPlayCardPlayer.UserName,outCard,canPlayCardPlayer.HandCards)
			checkAction(true, true, true,&ps,outCard)
		}
	}
}
