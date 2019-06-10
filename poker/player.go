package poker

import "fmt"

type player struct {
	name string
	handCards []card
}

func Player(name string)(player) {
	var player player
	player.name = name
	return player
}

func (player *player) addCard(card card) {
	player.handCards = append(player.handCards, card)
}

func (player player) showCards() {
	fmt.Print(player.name + "\t")
	for _, card := range player.handCards {
		fmt.Print(card.show() + "\t")
	}
	fmt.Println("")
}