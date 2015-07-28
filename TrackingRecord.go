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

func merge(trackingRecords []TrackingRecord )[]TrackingRecord{
	var mergedtrackingRecords[]TrackingRecord
	var outputPointer int = 0
	mergedtrackingRecords = append(mergedtrackingRecords,trackingRecords[0]);
	for i := 1; i < len(trackingRecords); i++ {
      if(trackingRecords[i].lo-mergedtrackingRecords[outputPointer].hi==1) &&
	  (trackingRecords[i].statusCode == mergedtrackingRecords[outputPointer].statusCode) &&
	  (trackingRecords[i].transferCode == mergedtrackingRecords[outputPointer].transferCode){
	  	  
		//t  := TrackingRecord{trackingRecords[i].statusCode,trackingRecords[i].transferCode,Range{mergedtrackingRecords[outputPointer].lo,trackingRecords[i].hi}}
			mergedtrackingRecords[outputPointer].hi = trackingRecords[i].hi
			} else{
			t  := TrackingRecord{trackingRecords[i].statusCode,trackingRecords[i].transferCode,Range{trackingRecords[i].lo,trackingRecords[i].hi}}
			mergedtrackingRecords = append(mergedtrackingRecords,t)
			outputPointer+=1
		}
    }
	return mergedtrackingRecords
	}
