package cards

import "math/rand"
import "time"

var ranks = []string {"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suits = []string {"hearts", "spades", "diamonds", "clubs"}

type Deck struct {
    Pile
    discard Pile
}

func NewDeck(num int) *Deck {
    var deck *Deck = new(Deck)

    //if num==nil {
    //    num=1;
    //}

    for i:=0; i<num;i++ {
        for rank := 0; rank < len(ranks); rank++ {
            for suit := 0; suit < len(suits); suit++ {
                deck.PutDown(NewCard(ranks[rank], suits[suit]))
            }
        }
    }
    

    deck.Shuffle()
    return deck
}

func (deck *Deck) DealCard() *Card {
    if len(deck.Cards) == 0 {
        if len(deck.discard.Cards) == 0 {
            return nil
        }

        for card := deck.discard.PickUp(); card != nil; card = deck.discard.PickUp() {
            deck.PutDown(card)
        }

        deck.Shuffle()
    }

    card := deck.Cards[len(deck.Cards) - 1]
    deck.Cards = deck.Cards[:len(deck.Cards) - 1]
    return card
}

func (deck *Deck) Discard(card *Card) {
    deck.discard.PutDown(card)
}

func (deck *Deck) Shuffle() {
    maxIndex := len(deck.Cards) - 1
    rand.Seed(int64(time.Now().UnixNano()))
    for i := 0; i < maxIndex; i++ {
        save         := deck.Cards[i]
        j            := rand.Intn(maxIndex - i) + i + 1
        deck.Cards[i] = deck.Cards[j]
        deck.Cards[j] = save
    }
}

func (deck *Deck) Stack(card *Card) *Deck {
    deck.PutDown(card)
    return deck
}
