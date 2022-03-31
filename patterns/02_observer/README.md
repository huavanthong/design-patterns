# Introduction
This tutorial help you understand about Observe pattern in Behavioral Patterns

# Question
* [Why we consider Youtube channels is a subject in ex1_youtube_sub.go?](#youtube-channel)
* [You recognize that observer is similar with service interface and service implmentation?](#)


# Answer Question
### Youtube channel
Youtube channels là subject là vì nó là điều gì đó đang được quan sát.  
Để hiểu rõ hơn về phân tích design này: [here](https://www.meisternote.com/app/note/xlyX-PxbCi8j/observer)
### Explain about Server Interface and Service Impl 
Chúng ta đã có kinh nghiệm thông qua QROrder.  
Refer: [here](https://github.com/huavanthong/MasterJava/tree/master/QROrder/src/main/java/com/qrorder/demo/service)  

Vậy qua đó, chúng ta thấy được vài điều sau đây:
1. ServiceImpl là nơi để suy nghĩ các methods trừu tượng cho việc sử dụng service của ta. Về cơ bản, ta thấy rằng ServiceImpl có nét tương đồng với Observer Design Pattern với điểm sau đây:
    - Đều cung cấp interface 
    - Sau đó, phải implement các interface đó.

2. Nhưng đồng thời ta thấy có điểm khác sau đay:  
    - Ta cung cấp ServiceImpl, nhưng Service chỉ implement một interface đó, không còn chỗ nào khác implement Interface đó nữa.
    - Ở Observer patterns, chúng ta thấy có nhiều chỗ implement interface, dẫn đến nó được dùng cho việc register các object mà ta cần.
    - Ở Observer patterns, ta cũng thấy được rằng, các Subject, phải có các method đang tính trạng thái như là: getState(), setState()

# Experience
### Intent
* Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
* Encapsulate the core (or common or engine) components in a Subject abstraction, and the variable (or optional or user interface) components in an Observer hierarchy.
* The "View" part of Model-View-Controller.

## Go programming
Các bước để implement Observer Pattern như sau:  
**Step 1:** Sử dụng Interface để implement Observer.  
**Step 2:** Phân tích và xác đinh các methods cần thiết cho Subject, và sử dụng Interface để provide các methods đó.  
**Step 3:** Khi xác định được methods rồi, ta cần phải suy nghĩ đến nội dung implement các methods đó.  
**Step 4:** Đồng thời cũng suy nghĩ xem, User cần gì, cần data gì cho các User, ==> Tạo interface cho User.  
**Step 5:** Có inteface user, thì phải có người dùng ==> Chúng ta cần tạo new User.  
```go 
// Step 1:
type Observer interface {
	update(interface{})
}

// Step 2:
type Subject interface {
	registerObserver(obs Observer)
	removeObserver(obs Observer)
	notifyObservers()
}

// Step 3:
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

// Step 4:
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

// Step 5: 
func NewUserInterface(username string) Observer {
	return &UserInterface{UserName: username, Videos: make([]*Video, 0)}
}
```
## Java programming
Các bước để implement Observer Pattern như sau:  
**Step 1:** Tạo các common function cho mấy chức năng khác implement nó.  
**Step 2:** Tạo table constant để ta có thể refer đến methods mà ta muốn xử lý.  
**Step 3:** Tạo các structure để handle cho từng methods.  
**Step 4:** Implement tất cả methods trong từng handler của nó.  
**Step 5:** Implement switch case to select the optional method.  

-----------------------------------------------------------------------
