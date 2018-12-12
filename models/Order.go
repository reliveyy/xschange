package models

import (
	"fmt"
	"log"
	"sort"
	"time"
)

// Order created by the user
type Order struct {
	UserID    uint     `json:"userID"`
	Selling   bool     `json:"selling"`
	Quantity  int      `json:"quantity"`
	Remain    int      `json:"remain"`
	Price     int      `json:"price"`
	Matchs    []*Match `json:"matchs"`
	CreatedAt int64    `json:"createAt"`
}

// DoSettlement caculate and set users' balance
func (o *Order) DoSettlement() {
}

// GetRemain return order remain quantity for trade
func (o *Order) GetRemain() int {
	matched := 0
	for _, match := range o.Matchs {
		matched += match.Quantity
	}
	return o.Quantity - matched
}

// Match create matchs for the order
func (o *Order) Match(orders *[]Order) {
	peers := []Order{}
	peers = choosePeers(*orders, !o.Selling)
	peers = filterByPrice(peers, !o.Selling, o.Price)
	sortPeers(peers, !o.Selling)

	fmt.Println("\033[2J")
	for _, peer := range peers {
		// fmt.Println(peer)
		var match Match
		if o.Remain >= peer.Remain {
			match = Match{Order: &peer, Quantity: peer.Remain, Price: peer.Price}
			o.Remain -= peer.Remain
			peer.Remain = 0
		} else {
			match = Match{Order: &peer, Quantity: o.Remain, Price: peer.Price}
			o.Remain = 0
			peer.Remain -= o.Remain
		}

		o.Matchs = append(o.Matchs, &match)
		if o.Remain == 0 {
			break
		}
	}

	for _, m := range o.Matchs {
		log.Println(*m)
	}
}

// Create place a new order
func (o *Order) Create(orders *[]Order) {
	o.Remain = o.Quantity
	o.CreatedAt = time.Now().Unix()
	*orders = append(*orders, *o)
	o.Match(orders)
	o.DoSettlement()
	log.Println(orders)
	log.Println((*orders)[1].Matchs)
}

// private

func choosePeers(orders []Order, forBuyer bool) []Order {
	peers := []Order{}
	for _, o := range orders {
		if o.Selling == forBuyer && o.Remain != 0 {
			peers = append(peers, o)
		}
	}

	return peers
}

func filterByPrice(peers []Order, forBuyer bool, price int) []Order {
	filtered := []Order{}
	for _, p := range peers {
		log.Println(p)
		if forBuyer && p.Price <= price {
			filtered = append(filtered, p)
		}

		if !forBuyer && p.Price >= price {
			filtered = append(filtered, p)
		}
	}

	return filtered
}

func sortPeers(peers []Order, selling bool) {
	sort.Slice(peers, func(i, j int) bool {
		if selling {
			return peers[i].Price < peers[j].Price
		}

		return peers[i].Price > peers[j].Price
	})
}
