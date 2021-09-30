/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSortedKeys(t *testing.T) {
	mapKeyValue := make(map[string]int)
	mapKeyValue["blue"] = 10
	mapKeyValue["apple"] = 15
	mapKeyValue["red"] = 12
	mapKeyValue["123"] = 22
	mapKeyValue["a"] = 33
	mapKeyValue[""] = 30
	require.Equal(t, []string{"", "123", "a", "apple", "blue", "red"}, GetSortedKeys(mapKeyValue))
}

func TestGetValuesBySortedKeys(t *testing.T) {
	type name struct {
		fName string
		lName string
	}
	mapKeyValue := make(map[string]*name)
	mapKeyValue["2"] = &name{"Two", "two"}
	mapKeyValue["3"] = &name{"Three", "three"}
	mapKeyValue["5"] = &name{"Five", "five"}
	mapKeyValue[""] = &name{"None", "none"}

	sortedRes := []*name{}
	GetValuesBySortedKeys(&mapKeyValue, &sortedRes)
	require.Equal(
		t,
		[]*name{{"None", "none"}, {"Two", "two"}, {"Three", "three"}, {"Five", "five"}},
		sortedRes,
	)
}
