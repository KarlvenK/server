// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameSubjectTopic = "chii_subject_topics"

// SubjectTopic mapped from table <chii_subject_topics>
type SubjectTopic struct {
	ID          uint32 `gorm:"column:sbj_tpc_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true"`
	SubjectID   uint32 `gorm:"column:sbj_tpc_subject_id;type:mediumint(8) unsigned;not null"`
	UID         uint32 `gorm:"column:sbj_tpc_uid;type:mediumint(8) unsigned;not null"`
	Title       string `gorm:"column:sbj_tpc_title;type:varchar(80);not null"`
	CreatedTime uint32 `gorm:"column:sbj_tpc_dateline;type:int(10) unsigned;not null"`
	UpdatedTime uint32 `gorm:"column:sbj_tpc_lastpost;type:int(10) unsigned;not null"`
	Replies     uint32 `gorm:"column:sbj_tpc_replies;type:mediumint(8) unsigned;not null"`
	State       uint8  `gorm:"column:sbj_tpc_state;type:tinyint(1) unsigned;not null"`
	Status      uint8  `gorm:"column:sbj_tpc_display;type:tinyint(1) unsigned;not null;default:1"`
}

// TableName SubjectTopic's table name
func (*SubjectTopic) TableName() string {
	return TableNameSubjectTopic
}
