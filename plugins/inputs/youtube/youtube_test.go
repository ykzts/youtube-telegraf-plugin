package youtube

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/api/youtube/v3"
)

func TestGetTags(t *testing.T) {
	id := "UCnrgOD6G0y0_rcubQuICpTQ"
	title := "InfluxData"

	channel := &youtube.Channel{
		Id: id,
		Snippet: &youtube.ChannelSnippet{
			Title: title,
		},
	}

	getTagsReturn := getTags(channel)

	correctTagsReturn := map[string]string{
		"id":    id,
		"title": title,
	}

	require.Equal(t, true, reflect.DeepEqual(getTagsReturn, correctTagsReturn))
}

func TestGetFields(t *testing.T) {
	subscribers := uint64(1)
	videos := uint64(2)
	views := uint64(3)

	channel := &youtube.Channel{
		Statistics: &youtube.ChannelStatistics{
			SubscriberCount: subscribers,
			VideoCount:      videos,
			ViewCount:       views,
		},
	}

	getFieldsReturn := getFields(channel)

	correctFieldReturn := map[string]interface{}{
		"subscribers": uint64(1),
		"videos":      uint64(2),
		"views":       uint64(3),
	}

	require.Equal(t, true, reflect.DeepEqual(getFieldsReturn, correctFieldReturn))
}
