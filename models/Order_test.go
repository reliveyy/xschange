package models

import (
	"testing"
)

func TestChoosePeers(t *testing.T) {
	orders := []Order{
		Order{Selling: true, Quantity: 1},
		Order{Selling: true, Quantity: 1},
		Order{Selling: false, Quantity: 1},
		Order{Selling: false, Quantity: 1},
		Order{Selling: false, Quantity: 1},
		Order{Selling: false, Quantity: 0},
	}

	peers := choosePeers(orders, false)
	if want := 3; len(peers) != want {
		t.Errorf("len(peers) == %v, want %v", len(peers), want)
	}
}

func TestFilterByPrice(t *testing.T) {
	peers := []Order{
		Order{Price: 1},
		Order{Price: 2},
		Order{Price: 3},
		Order{Price: 4},
		Order{Price: 5},
		Order{Price: 6},
	}

	var filtered []Order

	filtered = filterByPrice(peers, true, 4) // for buyer
	if want := 4; len(filtered) != want {
		t.Errorf("len(filtered) == %v, want %v", len(filtered), want)
	}
	if want := 1; filtered[0].Price != want {
		t.Errorf("filtered[0].Price == %v, want %v", filtered[0].Price, want)
	}
	if want := 2; filtered[1].Price != want {
		t.Errorf("filtered[1].Price == %v, want %v", filtered[1].Price, want)
	}
	if want := 3; filtered[2].Price != want {
		t.Errorf("filtered[2].Price == %v, want %v", filtered[2].Price, want)
	}
	if want := 4; filtered[3].Price != want {
		t.Errorf("filtered[3].Price == %v, want %v", filtered[3].Price, want)
	}

	filtered = filterByPrice(peers, false, 4) // for seller
	if want := 3; len(filtered) != want {
		t.Errorf("len(filtered) == %v, want %v", len(filtered), want)
	}
	if want := 4; filtered[0].Price != want {
		t.Errorf("filtered[0].Price == %v, want %v", filtered[0].Price, want)
	}
	if want := 5; filtered[1].Price != want {
		t.Errorf("filtered[1].Price == %v, want %v", filtered[1].Price, want)
	}
	if want := 6; filtered[2].Price != want {
		t.Errorf("filtered[2].Price == %v, want %v", filtered[2].Price, want)
	}
}

func TestSortPeers(t *testing.T) {
	peers := []Order{
		Order{Price: 10},
		Order{Price: 11},
		Order{Price: 12},
	}

	sortPeers(peers, false)
	if want := 12; peers[0].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
	if want := 11; peers[1].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
	if want := 10; peers[2].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}

	sortPeers(peers, true)

	if want := 10; peers[0].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
	if want := 11; peers[1].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
	if want := 12; peers[2].Price != want {
		t.Errorf("peers[0].Price == %v, want %v", peers[0].Price, want)
	}
}
