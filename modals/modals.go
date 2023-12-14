package modals

type User struct {
	Id              uint64          `json:"id" gorm:"primaryKey"`
	Username        string          `json:"username"`
	PendingRequests []FriendRequest `gorm:"foreignKey:ToUser"`
	SentRequests    []FriendRequest `gorm:"foreignKey:FromUser"`
	Friends         []Friend        `gorm:"foreignKey:MainUser"`
}

type FriendRequest struct {
	Id       uint64 `gorm:"primaryKey"`
	FromUser uint64
	ToUser   uint64
}

type Friend struct {
	Id         uint64 `gorm:"primaryKey"`
	MainUser   uint64
	FriendUser uint64
}
