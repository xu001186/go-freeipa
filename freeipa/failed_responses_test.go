// Copyright © 2022 IN2P3 Computing Centre, IN2P3, CNRS
// Copyright © 2018 Philippe Voinov
//
// Contributor(s): Remi Ferrand <remi.ferrand_at_cc.in2p3.fr>, 2021
//
// This software is governed by the CeCILL license under French law and
// abiding by the rules of distribution of free software.  You can  use,
// modify and/ or redistribute the software under the terms of the CeCILL
// license as circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and  rights to copy,
// modify and redistribute granted by the license, users are provided only
// with a limited warranty  and the software's author,  the holder of the
// economic rights,  and the successive licensors  have only  limited
// liability.
//
// In this respect, the user's attention is drawn to the risks associated
// with loading,  using,  modifying and/or developing or reproducing the
// software by the user in light of its specific status of free software,
// that may mean  that it is complicated to manipulate,  and  that  also
// therefore means  that it is reserved for developers  and  experienced
// professionals having in-depth computer knowledge. Users are therefore
// encouraged to load and test the software's suitability as regards their
// requirements in conditions enabling the security of their systems and/or
// data to be ensured and,  more generally, to use and operate it in the
// same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

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
