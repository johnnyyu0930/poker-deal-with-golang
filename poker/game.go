package poker

type game struct {
	players []player
	deck    deck
}

func Game(args ...string) game {
	var game game
	for _, name := range args {
		game.players = append(game.players, Player(name))
	}
	var deck deck = Deck()
	game.deck = deck
	game.deal()
	game.show()
	return game
}

func (game game) deal() {
	for i := 0; i < len(game.deck.cards); i++ {
		var card card = game.deck.cards[i]
		var player int = (i % 4)
		game.players[player].addCard(card)
	}
}

func (game game) show() {
	for _, player := range game.players {
		player.showCards()
	}
}
