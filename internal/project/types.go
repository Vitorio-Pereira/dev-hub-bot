package project

import "time"

type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "Low"
	case PriorityMedium:
		return "Medium"
	case PriorityHigh:
		return "High"
	default:
		return "Unknown"
	}
}

type Difficulty int

const (
	DifficultyLow Difficulty = iota
	DifficultyMedium
	DifficultyHigh
)

func (d Difficulty) String() string {
	switch d {
	case DifficultyLow:
		return "Low"
	case DifficultyMedium:
		return "Medium"
	case DifficultyHigh:
		return "High"
	default:
		return "Unknown"
	}
}

type Stack struct {
	Languages []string
	Infra     []string
	Other     []string
}

type Project struct {
	ID         string
	GuildID    string
	Name       string
	Repo       string
	Stack      Stack
	StartedAt  time.Time
	Priority   Priority
	Difficulty Difficulty
	Phase      string
	Progress   int
	Done       []string
	NextSteps  []string
	Blockers   []string
	UpdatedAt  time.Time
}
