package models

type User struct {
	Name           string   `bson:"name"`
	Phone          string   `bson:"phone"`
	Email          string   `bson:"email"`
	Address        string   `bson:"address"`
	AddressDetail  string   `bson:"addressDetail"`
	Message        string   `bson:"message"`
	ReferralSource []string `bson:"referralSource"`
	CreatedAt      string   `bson:"createdAt,omitempty"`
}
