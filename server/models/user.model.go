package models

import (
	"time"

	"gorm.io/gorm"
)



type User struct {
	
	UserId    		int64     			`gorm:"column:userId;primary_key;auto_increment:false"`
	UserName  		string     			`gorm:"column:userName;not null"`
	Password  		[]byte     			`gorm:"column:password;not null"`
	Email     		string     			`gorm:"column:email;not null"`
	Role      		UserRole   			`gorm:"column:role;not null;DEFAULT:'user'"`
	Coin      		int        			`gorm:"column:coin;not null;DEFAULT:50"`
	SellCoin  		int        			`gorm:"column:sellCoin;not null;DEFAULT:0"`
	UserImage 		string     			`gorm:"column:userImage;not null;DEFAULT:'https://cdn4.vectorstock.com/i/thumb-large/94/53/avatar-icon-person-man-vector-38549453.jpg'"`
	Gender    		GenderEnum 			`gorm:"column:gender;not null;DEFAULT:'unknow'"`
	CreatedAt 		time.Time  			`gorm:"column:createdAt;type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt 		time.Time  			`gorm:"column:updatedAt;type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
    // Deleted 	gorm.DeletedAt
	DeletedAt 		gorm.DeletedAt 		`gorm:"index"`
	
}


type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser UserRole = "user"

)

type GenderEnum string

const (
	GenderEnumMail   GenderEnum = "male"
	GenderEnumFemail GenderEnum = "female"
	GenderEnumUnknow GenderEnum = "unknow"
)

