package models

type ShogunRights interface {
	GetSlavesList() error
	CreateCard() (Card, error)
	BindCardToDaimyo(card Card) error
	// add get information about slave method
}
