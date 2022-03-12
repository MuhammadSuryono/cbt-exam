package session

import "time"

type SessionAccess struct {
	CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`                       //
	CreatedBy            int       `gorm:"column:created_by" json:"created_by"`                       //
	DeletedAt            time.Time `gorm:"column:deleted_at" json:"deleted_at"`                       //
	EndTime              string    `gorm:"column:end_time" json:"end_time"`                           //
	ID                   int64     `gorm:"column:id;primary_key" json:"id"`                           //
	IsActive             int       `gorm:"column:is_active" json:"is_active"`                         //
	ScheduleAnnouncement string    `gorm:"column:schedule_announcement" json:"schedule_announcement"` //
	SessionDay           string    `gorm:"column:session_day" json:"session_day"`                     //
	SessionKey           string    `gorm:"column:session_key" json:"session_key"`                     //
	SessionName          string    `gorm:"column:session_name" json:"session_name"`                   //
	StartTime            string    `gorm:"column:start_time" json:"start_time"`                       //
	UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`                       //
	UpdatedBy            int       `gorm:"column:updated_by" json:"updated_by"`                       //
	WithResult           int       `gorm:"column:with_result" json:"with_result"`                     //
}

// TableName sets the insert table name for this struct type
func (s *SessionAccess) TableName() string {
	return "session_access"
}
