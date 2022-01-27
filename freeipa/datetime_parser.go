package freeipa

import (
	"fmt"
	"time"
)

const (
	// 	RFC3339     = "2006-01-02T15:04:05Z07:00"
	LDAPGeneralizedTimeFormat = "20060102150405Z"
)

// parse LDAP generalized time format as present
// in some "__datetime__" responses
// See https://github.com/freeipa/freeipa/blob/master/ipalib/constants.py#L280
// LDAP_GENERALIZED_TIME_FORMAT = "%Y%m%d%H%M%SZ"
func parseFreeIPADateTimeStr(str string) (time.Time, error) {
	return time.Parse(LDAPGeneralizedTimeFormat, str)
}

// tryParseFreeIPADatetimeMap tries to solve https://github.com/ccin2p3/go-freeipa/issues/1
// Krbprincipalexpiration is returned as a []interface {}
// that is [map[__datetime__:20220428000000Z]]
func tryParseFreeIPADatetimeMap(m map[string]interface{}) (time.Time, error) {
	var tt time.Time
	dV, ok := m["__datetime__"]
	if !ok {
		return tt, fmt.Errorf("no __datetime__ key")
	}

	dsV, ok := dV.(string)
	if !ok {
		return tt, fmt.Errorf("__datetime__ key not a string")
	}

	return parseFreeIPADateTimeStr(dsV)
}
