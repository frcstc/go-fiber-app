/**
 * @Author: AF
 * @Date: 2021/8/9 18:28
 */

package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type BaseModel struct {
	CreatedTime LocalTime  `gorm:"autoCreateTime" json:"createdTime"`
	UpdatedTime LocalTime  `gorm:"autoCreateTime" json:"updatedTime"`
	DeletedTime *LocalTime `gorm:"index" json:"deletedTime"`
}

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
