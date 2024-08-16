package main

import "time"

type Heartbeat struct {
	LapsToGo   int       `json:"lapsToGo"`
	TimeToGo   string    `json:"timeToGo"` // Custom type to handle duration
	TimeOfDay  time.Time `json:"timeOfDay"`
	RaceTime   string    `json:"raceTime"` // Custom type to handle duration
	FlagStatus string    `json:"flagStatus"`
}

type CompetitorInfo struct {
	RegistrationNumber string `json:"registrationNumber"`
	Number             string `json:"number"`
	TransponderNumber  int    `json:"transponderNumber"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Nationality        string `json:"nationality"`
	ClassNumber        int    `json:"classNumber"`
}

type CompInfo struct {
	RegistrationNumber string `json:"registrationNumber"`
	Number             string `json:"number"`
	ClassNumber        int    `json:"classNumber"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Nationality        string `json:"nationality"`
}

type RunInfo struct {
	UniqueNumber int    `json:"uniqueNumber"`
	Description  string `json:"description"`
}

type ClassInfo struct {
	UniqueNumber int    `json:"uniqueNumber"`
	Description  string `json:"description"`
}

type SettingInfo struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

type RaceInfo struct {
	Position           int    `json:"position"`
	RegistrationNumber string `json:"registrationNumber"`
	Laps               int    `json:"laps"`
	TotalTime          string `json:"totalTime"` // Custom type to handle duration
}

type PracticeQualifyInfo struct {
	Position           int    `json:"position"`
	RegistrationNumber string `json:"registrationNumber"`
	BestLap            int    `json:"bestLap"`
	BestLaptime        string `json:"bestLaptime"` // Custom type to handle duration
}

type InitRecord struct {
	TimeOfDay time.Time `json:"timeOfDay"`
	Date      time.Time `json:"date"`
}

type PassingInfo struct {
	RegistrationNumber string `json:"registrationNumber"`
	LapTime            string `json:"lapTime"`   // Custom type to handle duration
	TotalTime          string `json:"totalTime"` // Custom type to handle duration
}

type CorrectedFinish struct {
	RegistrationNumber string `json:"registrationNumber"` // max 8 characters
	Number             string `json:"number"`             // max 5 characters
	Laps               int    `json:"laps"`               // 0 - 99999
	TotalTime          string `json:"totalTime"`          // "HH:MM:SS.DDD", Custom type to handle duration
	CorrectionTime     string `json:"correctionTime"`     // "+/-HH:MM:SS.DDD", Custom type to handle duration
}

type TimingMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
