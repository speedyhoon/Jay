package embedded

import "time"

// Alarm properties
type Alarm struct {
	SirenOn         time.Duration
	SirenOff        time.Duration
	Repeats         uint8
	CheckEvery      time.Duration
	Types           Cluster
	Attempts        uint8
	AttemptsFunc    func()
	KeyAttempts     uint8
	KeyAttemptsFunc func()
}

type Configuration struct { // J+
	Alarm   Alarm
	DataLog DataLog
	Types
}

// DataLog properties for data logging files.
type DataLog struct {
	Version     uint
	VersionRace uint
	FileSizeMax uint32
	Cluster
	Off bool
}

// Embedded types
type (
	UniMog  uint8
	Cluster []byte
)

type Types struct { // J+
	UniMog
	Cluster
}
