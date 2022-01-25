package freeipa

import (
	"encoding/json"
	"fmt"
)

const (
	FailedReasonNoSuchEntry    = "no such entry"
	FailedReasonAlreadyAMember = "This entry is already a member"
)

type FailedOperations map[string]map[string]failedOperations

type fromRootFailedOperations map[string]failedOperations

func (f fromRootFailedOperations) String() string {
	userFriendlyFailures := make(map[string]string)
	for rootFailName, fOperations := range f {
		for _, fOperation := range fOperations {
			fromRootName := fmt.Sprintf("%s/%s", rootFailName, fOperation.Name)
			userFriendlyFailures[fromRootName] = fOperation.Reason
		}
	}

	return fmt.Sprintf("%+v", userFriendlyFailures)
}

func (f FailedOperations) GetFailures() fromRootFailedOperations {
	failures := make(fromRootFailedOperations)
	for rootFailureCategoryName, v := range f {
		for subFailureCategoryName, failedOp := range v {
			if len(failedOp) > 0 {
				fromRootFailureName := fmt.Sprintf("%s/%s", rootFailureCategoryName, subFailureCategoryName)
				failures[fromRootFailureName] = append(failures[fromRootFailureName], failedOp...)
			}
		}
	}

	return failures
}

type failedOperation struct {
	Name   string
	Reason string
}

type failedOperations []failedOperation

func (f *failedOperations) UnmarshalJSON(b []byte) error {
	var rawFailedStr [][]string
	if err := json.Unmarshal(b, &rawFailedStr); err != nil {
		return err
	}

	*f = failedOperations{}
	for _, failedEntry := range rawFailedStr {
		if len(failedEntry) != 2 {
			return fmt.Errorf("failed entry %v does not have two elements", failedEntry)
		}

		*f = append(*f, failedOperation{
			Name:   failedEntry[0],
			Reason: failedEntry[1],
		})
	}

	return nil
}
