package global

import "github.com/jinzhu/gorm"

var (
	GlobalDb      *gorm.DB
	SecretSignKey = []byte("pulsar") //JWT秘钥
)
