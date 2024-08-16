package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func cutFixes(line string) (string, bool) {
	nl, foundPre := strings.CutPrefix(line, "$")
	nl = strings.TrimSuffix(nl, "\n")
	nl = strings.Trim(nl, "\r")
	if !foundPre {
		return line, false
	}
	return nl, true
}

func removeQuotes(val string) string {
	val = strings.TrimSuffix(val, `"`)
	val = strings.TrimPrefix(val, `"`)
	return val
}

func parseDuration(hhmmss string) (time.Duration, error) {
	hhmmss = strings.TrimSuffix(hhmmss, `"`)
	hhmmss = strings.TrimPrefix(hhmmss, `"`)
	parts := strings.Split(hhmmss, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid duration format")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hours: %v", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %v", err)
	}

	seconds, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid seconds: %v", err)
	}

	duration := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds*float64(time.Second))
	return duration, nil
}

func parseTimeWithCurrentDate(hhmmss string) (time.Time, error) {
	// Get the current date
	hhmmss = strings.TrimSuffix(hhmmss, `"`)
	hhmmss = strings.TrimPrefix(hhmmss, `"`)
	currentDate := time.Now().Format("2006-01-02")
	// Combine the current date with the provided time
	fullTimeStr := currentDate + " " + hhmmss

	// Define the layout corresponding to the full date-time string
	layout := "2006-01-02 15:04:05"

	// Parse the combined date-time string to a time.Time object
	parsedTime, err := time.Parse(layout, fullTimeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %v", err)
	}

	return parsedTime, nil
}

func parseHeartbeat(dat []string) (Heartbeat, error) {
	ltg, er := strconv.Atoi(removeQuotes(dat[1]))
	if er != nil {
		ltg = 0
	}
	ttg, err := parseDuration(dat[2])
	if err != nil {
		ttg = 0
	}
	tod, err := parseTimeWithCurrentDate(dat[3])
	if err != nil {
		tod = time.Time{}
	}
	rt, err := parseDuration(dat[4])
	if err != nil {
		rt = 0
	}
	fs := removeQuotes(dat[5])
	return Heartbeat{LapsToGo: ltg, TimeToGo: ttg.String(), TimeOfDay: tod, RaceTime: rt.String(), FlagStatus: fs}, nil
}

func parseCompetitorInfo(dat []string) (CompetitorInfo, error) {
	rn := removeQuotes(dat[1])
	num := removeQuotes(dat[2])
	tn, err := strconv.Atoi(removeQuotes(dat[3]))
	if err != nil {
		tn = 0
	}
	fn := removeQuotes(dat[4])
	ln := removeQuotes(dat[5])
	nat := removeQuotes(dat[6])
	cn, err := strconv.Atoi(removeQuotes(dat[7]))
	if err != nil {
		cn = 0
	}
	return CompetitorInfo{RegistrationNumber: rn, Number: num, TransponderNumber: tn, FirstName: fn, LastName: ln, Nationality: nat, ClassNumber: cn}, nil
}

func parseCompInfo(dat []string) (CompInfo, error) {
	rn := removeQuotes(dat[1])
	num := removeQuotes(dat[2])
	cn, err := strconv.Atoi(removeQuotes(dat[3]))
	if err != nil {
		cn = 0
	}
	fn := removeQuotes(dat[4])
	ln := removeQuotes(dat[5])
	nat := removeQuotes(dat[6])
	return CompInfo{RegistrationNumber: rn, Number: num, ClassNumber: cn, FirstName: fn, LastName: ln, Nationality: nat}, nil
}

func ParseRunInfo(dat []string) (RunInfo, error) {
	un, err := strconv.Atoi(removeQuotes(dat[1]))
	if err != nil {
		un = 0
	}
	des := removeQuotes(dat[2])
	return RunInfo{UniqueNumber: un, Description: des}, nil
}

func paseClassInfo(dat []string) (ClassInfo, error) {
	un, err := strconv.Atoi(removeQuotes(dat[1]))
	if err != nil {
		un = 0
	}
	des := removeQuotes(dat[2])
	return ClassInfo{UniqueNumber: un, Description: des}, nil
}

func parseSettingInfo(dat []string) SettingInfo {
	return SettingInfo{Description: removeQuotes(dat[1]), Value: removeQuotes(dat[2])}
}

func parseRaceInfo(dat []string) (RaceInfo, error) {
	pos, err := strconv.Atoi(removeQuotes(dat[1]))
	if err != nil {
		pos = 0
	}
	rn := removeQuotes(dat[2])
	laps, err := strconv.Atoi(removeQuotes(dat[3]))
	if err != nil {
		laps = 0
	}
	tt, err := parseDuration(removeQuotes(dat[4]))
	if err != nil {
		tt = 0
	}
	return RaceInfo{Position: pos, RegistrationNumber: rn, Laps: laps, TotalTime: tt.String()}, nil
}

func ParsePQInfo(dat []string) (PracticeQualifyInfo, error) {
	pos, err := strconv.Atoi(removeQuotes(dat[1]))
	if err != nil {
		pos = 0
	}
	rn := removeQuotes(dat[2])
	bl, err := strconv.Atoi(removeQuotes(dat[3]))
	if err != nil {
		bl = 0
	}
	blt, err := parseDuration(removeQuotes(dat[4]))
	if err != nil {
		blt = 0
	}
	return PracticeQualifyInfo{Position: pos, RegistrationNumber: rn, BestLap: bl, BestLaptime: blt.String()}, nil
}

func ParseInitRecord(dat []string) (InitRecord, error) {
	tod, err := parseTimeWithCurrentDate(removeQuotes(dat[1]))
	if err != nil {
		tod = time.Time{}
	}
	date, err := parseTimeWithCurrentDate(removeQuotes(dat[2]))
	if err != nil {
		date = time.Time{}
	}
	return InitRecord{TimeOfDay: tod, Date: date}, nil
}

func parsePassingInfo(dat []string) (PassingInfo, error) {
	rn := removeQuotes(dat[1])
	lt, err := parseDuration(removeQuotes(dat[2]))
	if err != nil {
		lt = 0
	}
	tt, err := parseDuration(removeQuotes(dat[3]))
	if err != nil {
		tt = 0
	}
	return PassingInfo{RegistrationNumber: rn, LapTime: lt.String(), TotalTime: tt.String()}, nil
}

func ParseCorrectedFinish(dat []string) (CorrectedFinish, error) {
	rn := removeQuotes(dat[1])
	num := removeQuotes(dat[2])
	laps, err := strconv.Atoi(removeQuotes(dat[3]))
	if err != nil {
		laps = 0
	}
	tt, err := parseDuration(removeQuotes(dat[4]))
	if err != nil {
		tt = 0
	}
	cor, err := parseDuration(removeQuotes(dat[5]))
	if err != nil {
		cor = 0
	}
	return CorrectedFinish{RegistrationNumber: rn, Number: num, Laps: laps, TotalTime: tt.String(), CorrectionTime: cor.String()}, nil
}
