/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"bytes"
	"testing"
	"time"
)

func TestComputeSHA256(t *testing.T) {
	if !bytes.Equal(ComputeSHA256([]byte("foobar")), ComputeSHA256([]byte("foobar"))) {
		t.Fatalf("Expected hashes to match, but they did not match")
	}
	if bytes.Equal(ComputeSHA256([]byte("foobar1")), ComputeSHA256([]byte("foobar2"))) {
		t.Fatalf("Expected hashes to be different, but they match")
	}
}

func TestComputeSHA3256(t *testing.T) {
	if !bytes.Equal(ComputeSHA3256([]byte("foobar")), ComputeSHA3256([]byte("foobar"))) {
		t.Fatalf("Expected hashes to match, but they did not match")
	}
	if bytes.Equal(ComputeSHA3256([]byte("foobar1")), ComputeSHA3256([]byte("foobar2"))) {
		t.Fatalf("Expected hashed to be different, but they match")
	}
}

func TestUUIDGeneration(t *testing.T) {
	uuid := GenerateUUID()
	if len(uuid) != 36 {
		t.Fatalf("UUID length is not correct. Expected = 36, Got = %d", len(uuid))
	}
	uuid2 := GenerateUUID()
	if uuid == uuid2 {
		t.Fatalf("Two UUIDs are equal. This should never occur")
	}
}

func TestTimestamp(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("timestamp now: %v", CreateUtcTimestamp())
		time.Sleep(200 * time.Millisecond)
	}
}

func TestToChaincodeArgs(t *testing.T) {
	expected := [][]byte{[]byte("foo"), []byte("bar")}
	actual := ToChaincodeArgs("foo", "bar")
	if len(expected) != len(actual) {
		t.Fatalf("Got %v, expected %v", actual, expected)
	}
	for i := range expected {
		if !bytes.Equal(expected[i], actual[i]) {
			t.Fatalf("Got %v, expected %v", actual, expected)
		}
	}
}

func TestConcatenateBytesNormal(t *testing.T) {
	first := []byte("first")
	second := []byte("second")
	third := []byte("third")

	result := ConcatenateBytes(first, second, third)
	expected := []byte("firstsecondthird")
	if !bytes.Equal(result, expected) {
		t.Errorf("Did not concatenate bytes correctly, expected %s, got %s", expected, result)
	}
}

func TestConcatenateBytesNil(t *testing.T) {
	first := []byte("first")
	second := []byte(nil)
	third := []byte("third")

	result := ConcatenateBytes(first, second, third)
	expected := []byte("firstthird")
	if !bytes.Equal(result, expected) {
		t.Errorf("Did not concatenate bytes correctly, expected %s, got %s", expected, result)
	}
}
