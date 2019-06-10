package poker

type card struct {
	suit  string
	point string
}

func newCard(s string, p string) card {
	var Card card
	Card.setSuit(s)
	Card.setPoint(p)
	return Card
}

func (card *card) setPoint(point string) {
	card.point = point
}

func (card *card) setSuit(suit string) {
	card.suit = suit
}

func (card card) getCard() card {
	return card
}

func (card card) show() string {
	return card.suit + card.point
}
