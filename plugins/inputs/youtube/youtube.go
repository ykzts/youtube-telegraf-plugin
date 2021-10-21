package youtube

import (
	"context"
	"strings"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// YouTube - plugin main structure
type YouTube struct {
	Channels       []string `toml:"channels"`
	Videos         []string `toml:"videos"`
	APIKey         string   `toml:"api_key"`
	youtubeService *youtube.Service
}

const sampleConfig = `
  ## List of channels to monitor.
  channels = [
    "UCBR8-60-B28hp2BmDPdntcQ",
    "UCnrgOD6G0y0_rcubQuICpTQ"
  ]

  ## List of videos to monitor.
  videos = [
    "gjoHHYnXdqs",
    "OoCsY8odmpM"
  ]

  ## Google API key.
  # api_key = ""
`

// SampleConfig returns sample configuration for this plugin.
func (y *YouTube) SampleConfig() string {
	return sampleConfig
}

// Description returns the plugin description.
func (y *YouTube) Description() string {
	return "Gather channel information from YouTube channels."
}

// Create YouTube Service
func (y *YouTube) createYouTubeService(ctx context.Context) (*youtube.Service, error) {
	return youtube.NewService(ctx, option.WithAPIKey(y.APIKey))
}

// Gather YouTube Metrics
func (y *YouTube) Gather(acc telegraf.Accumulator) error {
	ctx := context.Background()

	if y.youtubeService == nil {
		service, err := y.createYouTubeService(ctx)
		if err != nil {
			return err
		}

		y.youtubeService = service
	}

	if len(y.Channels) > 0 {
		call := y.youtubeService.Channels.
			List([]string{"snippet", "statistics"}).
			Id(strings.Join(y.Channels, ",")).
			MaxResults(50)

		resp, err := call.Do()
		if err != nil {
			return err
		}

		now := time.Now()

		for _, item := range resp.Items {
			tags := getTags(item)
			fields := getFields(item)

			acc.AddFields("youtube_channel", fields, tags, now)
		}
	}

	if len(y.Videos) > 0 {
		call := y.youtubeService.Videos.
			List([]string{"snippet", "statistics"}).
			Id(strings.Join(y.Videos, ",")).
			MaxResults(50)

		resp, err := call.Do()
		if err != nil {
			return err
		}

		now2 := time.Now()

		for _, item := range resp.Items {
			tags := getVideoTags(item)
			fields := getVideoFields(item)

			acc.AddFields("youtube_video", fields, tags, now2)
		}
	}

	return nil
}

func getTags(channelInfo *youtube.Channel) map[string]string {
	return map[string]string{
		"id":    channelInfo.Id,
		"title": channelInfo.Snippet.Title,
	}
}

func getFields(channelInfo *youtube.Channel) map[string]interface{} {
	return map[string]interface{}{
		"subscribers": channelInfo.Statistics.SubscriberCount,
		"videos":      channelInfo.Statistics.VideoCount,
		"views":       channelInfo.Statistics.ViewCount,
	}
}

func getVideoTags(videoInfo *youtube.Video) map[string]string {
	return map[string]string{
		"id":    videoInfo.Id,
		"title": videoInfo.Snippet.Title,
	}
}

func getVideoFields(videoInfo *youtube.Video) map[string]interface{} {
	return map[string]interface{}{
		"comments":  videoInfo.Statistics.CommentCount,
		"dislikes":  videoInfo.Statistics.DislikeCount,
		"favorites": videoInfo.Statistics.FavoriteCount,
		"likes":     videoInfo.Statistics.LikeCount,
		"views":     videoInfo.Statistics.ViewCount,
	}
}

func init() {
	inputs.Add("youtube", func() telegraf.Input {
		return &YouTube{}
	})
}
