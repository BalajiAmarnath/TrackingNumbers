package TrackingRecord

import "ftm"

type TrackingRecord struct {
	range Range
	statusCode rune
	transferCode int
}

func copy(record TrackingRecord) TrackingRecord {
	var copyRecord := TrackingRecord{record.range.copy(),record.range.transferCode,record.range.statusCode} 
}

