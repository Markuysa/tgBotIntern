package models

type DaimyoRights interface {
	GetCardsList() error
	CreateCardIncreasementRequest() error
	GetSamuraiList() error
	SetCardsBalances() error
	BindShogun() error
}
