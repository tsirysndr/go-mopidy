package mopidy

import "github.com/dghubble/sling"

type Client struct {
	base      *sling.Sling
	common    service
	History   *HistoryService
	Library   *LibraryService
	Mixer     *MixerService
	Playback  *PlaybackService
	Playlists *PlaylistsService
	Tracklist *TracklistService
}

type service struct {
	client *Client
}

func NewClient(baseURL string) *Client {
	base := sling.New().Base(baseURL).Set("Accept", "application/json")
	c := &Client{}
	c.base = base
	c.common.client = c
	c.History = (*HistoryService)(&c.common)
	c.Library = (*LibraryService)(&c.common)
	c.Mixer = (*MixerService)(&c.common)
	c.Playback = (*PlaybackService)(&c.common)
	c.Playlists = (*PlaylistsService)(&c.common)
	c.Tracklist = (*TracklistService)(&c.common)
	return c
}
