package freeipa

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_failedOperations_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    failedOperations
		wantErr bool
	}{
		{
			name:    "no entries",
			jsonStr: `[]`,
			want:    failedOperations{},
			wantErr: false,
		},
		{
			name:    "single entry",
			jsonStr: `[["admin", "no such entry"]]`,
			want: failedOperations{
				failedOperation{
					Name:   "admin",
					Reason: "no such entry",
				},
			},
			wantErr: false,
		},
		{
			name:    "multiple entries",
			jsonStr: `[["admin", "no such entry"], ["other", "This entry is already a member"]]`,
			want: failedOperations{
				failedOperation{
					Name:   "admin",
					Reason: "no such entry",
				},
				failedOperation{
					Name:   "other",
					Reason: "This entry is already a member",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var failedOp failedOperations
			if err := json.Unmarshal([]byte(tt.jsonStr), &failedOp); err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected error %v", err)
				}
			} else {
				if tt.wantErr {
					t.Error("no error caught")
				}
			}

			if !reflect.DeepEqual(failedOp, tt.want) {
				t.Errorf("got %+v, want %+v", failedOp, tt.want)
			}
		})
	}
}
