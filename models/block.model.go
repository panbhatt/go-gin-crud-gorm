package models

import "time"

type Block struct {
	Hash         string    `gorm:"type:varchar(66);primary_key"`
	number       uint32    `gorm:"type:int8;not null"`
	PreviousHash string    `gorm:"type:varchar(66);column:prev_block_h"`
	TxCount      uint16    `gorm:"type:type:int8;column:tx_count"`
	TimeStamp    time.Time `gorm:"type:numeric(10);`
}
