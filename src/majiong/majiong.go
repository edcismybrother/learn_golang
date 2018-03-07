package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	//	"container/list"
)

type Player struct {
	UserId    int
	UserName  string
	HandCards []Card
	Zhuang    bool
}

var (
	nameList          = []string{"刘德华", "张学友", "成龙", "周星驰", "张国荣", "李小龙", "皮智文"}
	cards             = []Card{}
	color             = []string{"万", "条", "筒", "字"}
	value             = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	canPlayCardPlayer *Player
	end               bool = true
	ps = []Player{}
)

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

func Init(num int,players *[]Player) {
	l := len(nameList)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	lens := r.Perm(l)
	for i := 0; i < num; i++ {
		if i == 0 {
			p := Player{i + 1, nameList[lens[i]], []Card{}, true}
			*players = append(*players, p)
//			canPlayCardPlayer = &(*players)[0]
		} else {
			p := Player{i + 1, nameList[lens[i]], []Card{}, false}
			*players = append(*players, p)
		}
	}
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

func FaPai(players *[]Player) {
	for i := 0; i < len(*players); i++ {
//		p := &players[i]
		isZ := (*players)[i].Zhuang
		if isZ {
			(*players)[i].HandCards = append((*players)[i].HandCards, cards[0:14]...)
			cards = cards[14:]
		} else {
			(*players)[i].HandCards = append((*players)[i].HandCards, cards[0:13]...)
			cards = cards[13:]
		}
	}
}

func ColorToInt(s string) int {
	var in int
	if s == "万" {
		in = 1
	} else if s == "条" {
		in = 2
	} else if s == "筒" {
		in = 3
	} else if s == "字" {
		in = 4
	}
	return in
}
func IntToColor(in int) string {
	var s string
	if in == 1 {
		s = "万"
	} else if in == 2 {
		s = "条"
	} else if in == 3 {
		s = "筒"
	} else if in == 4 {
		s = "字"
	}
	return s
}

type cardArray []Card

func (card cardArray) Len() int {
	return len(card)
}
func (card cardArray) Less(i, j int) bool {
	ci := ColorToInt(card[i].color)
	cj := ColorToInt(card[j].color)
	cii := ci*10 + int(card[i].value)
	cjj := cj*10 + int(card[j].value)
	return cii < cjj
}
func (card cardArray) Swap(i, j int) { card[i], card[j] = card[j], card[i] }

/**打牌移除手牌**/
func Remove(slice []Card, elems ...Card) []Card {
	isInElems := make(map[Card]bool)
	for _, elem := range elems {
		isInElems[elem] = true
	}
	w := 0
	for _, elem := range slice {
		if !isInElems[elem] {
			slice[w] = elem
			w++
		}
	}
	return slice[:w]
}

/**下一个出牌的玩家**/
func nextPlayer(players []Player, actionPlayer *Player) *Player {
	if actionPlayer == nil {
		//		pp :=  *players
		return &players[0]
	} else {
		for k, p := range players {
			if p.UserId == actionPlayer.UserId {
				index := (k + 1) % len(players)
				return &(players)[index]
			}
		}
	}
	return nil
}

func intToCard(number int) Card {
	value := byte(number % 100)
	color := IntToColor(number / 100)
	return Card{value, color}
}

/**检查出牌是否存在于手牌中**/
func checkOutCard(cards []Card, outCard Card) bool {
	for i := 0; i < len(cards); i++ {

		if cards[i] == outCard {
			cards = append(cards[:i], cards[i+1:]...)
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

/**查胡**/
func checkHu(players []Player, card Card) bool {
	return false
}

/**查碰**/
func checkPeng(players []Player, card Card) (bool, *Player) {
	for i := 0; i < len(players); i++ {
		handCards := players[i].HandCards
		num := 0
		for _, c := range handCards {
			if c == card {
				num++
			}
		}
		if num == 2 {
			return true, &players[i]
		}
	}
	return false, nil
}

/**查杠**/
func checkGang(players []Player, card Card) bool {
	return false
}

/**检查是否有人可以进行特殊操作**/
func checkAction(needHu bool, needPeng bool, needGang bool, players []Player, card Card) ([]int, *Player) {
	var canHu bool
	var canPeng bool
	var canGang bool
	var p *Player
	var actions []int

	if needHu {
		canHu = checkHu(players, card)
	}
	if needPeng {
		canPeng, p = checkPeng(players, card)
	}
	if needGang {
		canGang = checkGang(players, card)
	}

	if canHu {
		actions = append(actions, 3)
		fmt.Println("可以胡牌")
		return []int{},nil
	} else if canPeng || canGang {
		if canPeng {actions = append(actions, 1)}
		if canGang {actions = append(actions, 2)}
		fmt.Printf("玩家 ： %v canPeng：%v canGang：%v\n", p.UserName, canPeng, canGang)
	return actions,p
	}
	return []int{},nil
}

/**处理碰**/
func handlePeng(player *Player) {

}

/**处理杠**/
func handleGang(player *Player) {

}

/**处理胡**/
func handleHu(player *Player) {

}

/**输入需要进行的操作**/
func handleAction(action int, player *Player) {
	fmt.Println("请输入操作对应的序号，0不进行操作，1碰2杠3胡")
	if action == 0 {
		fmt.Printf("玩家%v放弃特殊操作\n", player.UserName)
	} else if action == 1 {
		handlePeng(player)
	} else if action == 2 {
		handleGang(player)
	} else if action == 3 {
		handleHu(player)
	}
}

func moCard(cards []Card,p *Player){
	moC := cards[0]
	cards = append(cards[1:])
	p.HandCards = append(p.HandCards, moC)
	sort.Sort(cardArray(p.HandCards))
}

func main() {

	var member int
	member = 3
	fmt.Printf("初始化玩家人数，人数：%v \n", member)
	Init(member,&ps)
	fmt.Println(ps)
	fmt.Println("初始玩家成功，开始洗牌")
	XiPai()
	fmt.Println("发牌")
	FaPai(&ps)
	for i := 0;i < len(ps);i++ {
		sort.Sort(cardArray(ps[i].HandCards))
		fmt.Println(ps[i])
	}
	var input int
	for end {
		canPlayCardPlayer = nextPlayer(ps, canPlayCardPlayer)
		if len(canPlayCardPlayer.HandCards)%3 != 2 {
			moCard(cards,canPlayCardPlayer)
		}
		fmt.Printf("玩家 ： %v 请出牌", canPlayCardPlayer.UserName)
	a:
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			outCard := intToCard(input)
			if checkOutCard(canPlayCardPlayer.HandCards, outCard) == false {
				fmt.Println("该outcard不存在于手牌中")
				goto a
			}
			fmt.Printf("玩家：%v 打出 %v\n手牌为：%v\n", canPlayCardPlayer.UserName, outCard, canPlayCardPlayer.HandCards)
			var actions []int
			var action int
			actions,canPlayCardPlayer = checkAction(true, true, true, ps, outCard)
			
			if len(actions)>0{
				fmt.Printf("可以进行的操作有：%v",actions)
				fmt.Scan(&action)
				handleAction(action, canPlayCardPlayer)
			}
		}
	}
}
