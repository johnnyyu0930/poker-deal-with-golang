package poker

type deck struct {
	cards [52]card
}

func Deck() deck {
	var deck deck
	suits := [4]string{"H", "S", "D", "C"}
	points := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	i := 0
	for _, suit := range suits {
		for _, point := range points {
			deck.cards[i] = newCard(suit, point)
			i++
		}
	}
	return deck
}
