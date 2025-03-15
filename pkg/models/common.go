package models

import "time"

type CommonModel struct {
	ID         string    `json:"id" gorm:"primary_key;column:id;type:varchar(32);comment:ID"`
	CreateTime time.Time `json:"createTime,omitempty" gorm:"column:create_time;type:datetime;comment:创建时间"`
	UpdateTime time.Time `json:"updateTime,omitempty" gorm:"column:update_time;type:datetime;comment:更新时间"`
	EndTime    time.Time `json:"endTime,omitempty" gorm:"column:end_time;type:datetime;default:'9999-12-31 00:00:00';comment:逻辑删除时间"`
	UpdateBy   string    `json:"updateBy,omitempty" gorm:"column:update_by;type:varchar(32);comment:更新人"`
	CreateBy   string    `json:"createBy,omitempty" gorm:"column:create_by;type:varchar(32);comment:创建人"`
	DelFlag    int       `json:"delFlag,omitempty" gorm:"column:del_flag;type:int;size:32;default:0;comment:逻辑删除标志【0 ： 未删除 1： 已删除】"`
}
