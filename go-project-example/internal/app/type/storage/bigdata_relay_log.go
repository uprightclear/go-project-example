package storage

import "time"

type RelayOpLog struct {
	ID        int       `gorm:"column:id;type:int;primary_key;auto_increment"`
	OpType    string    `gorm:"column:op_type;type:varchar(32);not null;default:''"`
	ReqType   string    `gorm:"column:req_type;type:varchar(8);not null;default:''"`
	Target    string    `gorm:"column:target;type:varchar(32);not null;default:''"`
	ReqData   string    `gorm:"column:req_data;type:text"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp;index:idx_created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:current_timestamp;index:idx_updated_at"`
}

func (relayOpLog RelayOpLog) TableName() string {
	return "relay_op_log"
}
