package main

import "fmt"

// Step 1: Consider user is Observer
type Observer interface {
	update(interface{})
}

// Step 2: Consider youtube channel is Subject
type Subject interface {
	registerObserver(obs Observer)
	removeObserver(obs Observer)
	notifyObservers()
}

type Video struct {
	title string
}

// YoutubeChannel is a concrete implementation of Subject interface
type YoutubeChannel struct {
	Observers []Observer
	NewVideo  *Video
}

func (yt *YoutubeChannel) registerObserver(obs Observer) {
	yt.Observers = append(yt.Observers, obs)
}

func (yt *YoutubeChannel) removeObserver(obs Observer) {
	//
}

// notify to all observers when a new video is released
func (yt *YoutubeChannel) notifyObservers() {
	for _, obs := range yt.Observers {
		obs.update(yt.NewVideo)
	}
}

func (yt *YoutubeChannel) ReleaseNewVideo(video *Video) {
	yt.NewVideo = video
	yt.notifyObservers()
}

// UserInterface is a concrete implementation of Observer interface
type UserInterface struct {
	UserName string
	Videos   []*Video
}

func (ui *UserInterface) update(video interface{}) {
	ui.Videos = append(ui.Videos, video.(*Video))
	for video := range ui.Videos {
		View.AddChildNode(NewVideoComponent(video))
	}
	fmt.Printf("UI %s - Video: '%s' has just been released\n", ui.UserName, video.(*Video).title)
}

func NewUserInterface(username string) Observer {
	return &UserInterface{UserName: username, Videos: make([]*Video, 0)}
}

// usage

func main() {
	var ytChannel Subject = &YoutubeChannel{}
	ui1 := NewUserInterface("Bob")
	ui2 := NewUserInterface("Peter")
	ytChannel.registerObserver(ui1)
	ytChannel.registerObserver(ui2)
	ytChannel.(*YoutubeChannel).ReleaseNewVideo(&Video{title: "Avatar 2 trailer"})
	ytChannel.(*YoutubeChannel).ReleaseNewVideo(&Video{title: "Avengers Endgame trailer"})
}
