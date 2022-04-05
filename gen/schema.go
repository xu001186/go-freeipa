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

package main

type SchemaDump struct {
	Result struct {
		Result Schema `json:"result"`
	} `json:"result"`
}

type Schema struct {
	Topics   []*Topic   `json:"topics"`
	Classes  []*Class   `json:"classes"`
	Commands []*Command `json:"commands"`

	Fingerprint string `json:"fingerprint"`
	TTL         int    `json:"ttl"`
	Version     string `json:"version"`
}

type Topic struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Doc      string `json:"doc"`
	Version  string `json:"version"`
}

type Class struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Version  string `json:"version"`

	Params []*Param `json:"params"`
}

type Command struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Doc      string `json:"doc"`
	Version  string `json:"version"`

	AttrName   string `json:"attr_name"`
	ObjClass   string `json:"obj_class"`
	TopicTopic string `json:"topic_topic"`

	Params []*Param         `json:"params"`
	Output []*CommandOutput `json:"output"`
}

type Param struct {
	Name    string `json:"name"`
	CliName string `json:"cli_name"`
	Label   string `json:"label"`
	Doc     string `json:"doc"`

	Type        string `json:"type"`
	Multivalue  bool   `json:"multivalue"`
	CliMetavar  string `json:"cli_metavar"`
	RequiredRaw *bool  `json:"required"` // use Requried() instead
	Positional  bool   `json:"positional"`

	AlwaysAsk bool     `json:"alwaysask"`
	NoConvert bool     `json:"no_convert"`
	Exclude   []string `json:"exclude"`

	Default          []string `json:"default"`
	DefaultFromParam []string `json:"default_from_param"`
}

type CommandOutput struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`

	Type        string `json:"type"`
	Multivalue  bool   `json:"multivalue"`
	RequiredRaw *bool  `json:"required"` // use Required() instead
}

func (t *Param) Required() bool {
	if t.RequiredRaw == nil {
		return len(t.Default) == 0 && len(t.DefaultFromParam) == 0
	}
	return *t.RequiredRaw
}

func (t *CommandOutput) Required() bool {
	if t.RequiredRaw == nil {
		return true
	}
	return *t.RequiredRaw
}

func (t *Command) PosParams() []*Param {
	var out []*Param
	for _, p := range t.Params {
		if p.Positional {
			out = append(out, p)
		}
	}
	return out
}

func (t *Command) KwParams() []*Param {
	var out []*Param
	for _, p := range t.Params {
		if !p.Positional {
			out = append(out, p)
		}
	}
	return out
}
