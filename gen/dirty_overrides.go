//
// Author: Remi Ferrand <remi.ferrand_at_cc<dot>in2p3<dot>fr>
//
package main

import (
	"encoding/json"
	"io/ioutil"
)

type DirtyOverrides struct {
	Classes map[string]ClassOverrides `json:"classes"`
}

type ClassOverrides struct {
	Params map[string]ClassParamsOverrides `json:"params"`
}

type ClassParamsOverrides struct {
	Required *bool `json:"required,omitempty"`
}

func (c ClassParamsOverrides) OverrideParams(p *Param) {
	if c.Required != nil {
		p.RequiredRaw = c.Required
	}
}

func loadDirtyOverrides() (DirtyOverrides, error) {
	var overrides DirtyOverrides
	in, e := ioutil.ReadFile("../data/dirty_overrides.json")
	if e != nil {
		return overrides, e
	}
	if e = json.Unmarshal(in, &overrides); e != nil {
		return overrides, e
	}
	return overrides, nil
}
