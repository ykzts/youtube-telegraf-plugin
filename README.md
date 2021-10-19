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

  ## List of videos to monitor.
  videos = [
    "gjoHHYnXdqs",
    "OoCsY8odmpM"
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
- youtube_video
  - tags:
    - id - The ID of the video
    - title - The title name
  - fields:
    - comments (int)
    - dislikes (int)
    - favorites (int)
    - likes (int)
    - views (int)

### Example Output

```plain
youtube_channel,id=UCBR8-60-B28hp2BmDPdntcQ,title=YouTube subscribers=32000000i,videos=409i,views=2509852022i 1634655873951240300
youtube_channel,id=UCnrgOD6G0y0_rcubQuICpTQ,title=InfluxData videos=261i,views=489944i,subscribers=4520i 1634655873951240300
youtube_video,id=gjoHHYnXdqs,title=InfluxDB:\ The\ time\ series\ data\ platform\ built\ for\ developers comments=0i,dislikes=1i,favorites=0i,likes=8i,views=377i 1634655874031796300
youtube_video,id=OoCsY8odmpM,title=Intro\ to\ Time\ Series\ Databases\ &\ Data\ |\ Getting\ Started\ [1\ of\ 7] comments=0i,dislikes=20i,favorites=0i,likes=514i,views=74740i 1634655874031796300
```
