package TrackingRecord

import "fmt"

type TrackingRecord struct {
	statusCode rune
	transferCode int
	Range
}

func copyTrackingRecord(record TrackingRecord) TrackingRecord{
	var copyRecord = TrackingRecord{record.statusCode,record.transferCode,copy(record.Range)} 
	return copyRecord
}

