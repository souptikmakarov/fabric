/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kafka

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChannel(t *testing.T) {
	chn := newChannel(channelNameForTest(t), defaultPartition)

	expectedTopic := channelNameForTest(t)
	actualTopic := chn.topic()
	require.Equal(t, expectedTopic, actualTopic, "Got the wrong topic, expected %s, got %s instead", expectedTopic, actualTopic)

	expectedPartition := int32(defaultPartition)
	actualPartition := chn.partition()
	require.Equal(t, expectedPartition, actualPartition, "Got the wrong partition, expected %d, got %d instead", expectedPartition, actualPartition)
}
