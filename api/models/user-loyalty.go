package models

type UserLoyalty struct {
	CustomerID        uint            `json:"user_id"`
	LoyaltyCampaingID uint            `json:"loyalty_card_id"`
	User              User            `gorm:"foreignKey:CustomerID"`
	LoyaltyCampaing   LoyaltyCampaing `gorm:"foreignkry:LoyaltyCampaingID"`
}
