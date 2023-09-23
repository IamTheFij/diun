package model

import (
	"time"

	"github.com/crazy-max/diun/v4/pkg/utl"
)

// Watch holds data necessary for watch configuration
type Watch struct {
	Workers         int            `yaml:"workers,omitempty" json:"workers,omitempty" validate:"required,min=1"`
	Schedule        string         `yaml:"schedule,omitempty" json:"schedule,omitempty"`
	Jitter          *time.Duration `yaml:"jitter,omitempty" json:"jitter,omitempty" validate:"required"`
	FirstCheckNotif *bool          `yaml:"firstCheckNotif,omitempty" json:"firstCheckNotif,omitempty" validate:"required"`
	RunOnStartup    *bool          `yaml:"runOnStartup,omitempty" json:"runOnStartup,omitempty" validate:"required"`
	CompareDigest   *bool          `yaml:"compareDigest,omitempty" json:"compareDigest,omitempty" validate:"required"`
	Healthchecks    *Healthchecks  `yaml:"healthchecks,omitempty" json:"healthchecks,omitempty"`
	ImageDefaults   *ImageDefaults `yaml:"imageDefaults,omitempty" json:"imageDefaults,omitempty"`
}

// GetDefaults gets the default values
func (s *Watch) GetDefaults() *Watch {
	n := &Watch{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *Watch) SetDefaults() {
	s.Workers = 10
	s.Jitter = utl.NewDuration(30 * time.Second)
	s.FirstCheckNotif = utl.NewFalse()
	s.RunOnStartup = utl.NewTrue()
	s.CompareDigest = utl.NewTrue()
}
