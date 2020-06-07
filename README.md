# YouTube Telegraf Plugin

Gather channel information from [YouTube](https://www.youtube.com/) channels.

### Configuration

```toml
[[inputs.youtube]]
  ## List of channels to monitor.
  channels = [
    "UCBR8-60-B28hp2BmDPdntcQ",
    "UCnrgOD6G0y0_rcubQuICpTQ"
  ]

  ## YouTube API key.
  # api_key = ""
```

### Metrics

- youtube_channel
  - tags:
    - id - The ID of the channel
    - title - The title name
  - fields:
    - subscribers (int)
    - videos (int)
    - views (int)

### Example Output

```plain
youtube_channel,id=UCBR8-60-B28hp2BmDPdntcQ,title=YouTube subscribers=30400000i,videos=318i,views=2256168113i 1591499999079340600
youtube_channel,id=UCnrgOD6G0y0_rcubQuICpTQ,title=InfluxData subscribers=2740i,videos=150i,views=252193i 1591499999079340600
```
