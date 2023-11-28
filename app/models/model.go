package models

import (
	"database/sql/driver"
	"fmt"
	"github.com/spf13/cast"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type LocalTime time.Time

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt *LocalTime `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	//CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	//UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// 获取 ID 的字符串格式
func (b BaseModel) GetStringID() string {
	return cast.ToString(b.ID)
}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
