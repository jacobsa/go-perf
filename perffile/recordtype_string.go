// Code generated by "stringer -type=RecordType"; DO NOT EDIT

package perffile

import "fmt"

const (
	_RecordType_name_0 = "RecordTypeMmapRecordTypeLostRecordTypeCommRecordTypeExitRecordTypeThrottleRecordTypeUnthrottleRecordTypeForkRecordTypeReadRecordTypeSamplerecordTypeMmap2RecordTypeAux"
	_RecordType_name_1 = "recordTypeUserStartrecordTypeHeaderEventTyperecordTypeHeaderTracingDatarecordTypeHeaderBuildIDrecordTypeHeaderFinishedRoundrecordTypeHeaderIDIndex"
)

var (
	_RecordType_index_0 = [...]uint8{0, 14, 28, 42, 56, 74, 94, 108, 122, 138, 153, 166}
	_RecordType_index_1 = [...]uint8{0, 19, 44, 71, 94, 123, 146}
)

func (i RecordType) String() string {
	switch {
	case 1 <= i && i <= 11:
		i -= 1
		return _RecordType_name_0[_RecordType_index_0[i]:_RecordType_index_0[i+1]]
	case 64 <= i && i <= 69:
		i -= 64
		return _RecordType_name_1[_RecordType_index_1[i]:_RecordType_index_1[i+1]]
	default:
		return fmt.Sprintf("RecordType(%d)", i)
	}
}
