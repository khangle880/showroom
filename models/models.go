package models

type Room struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	URLKey           string `json:"url_key"`
	ImageURL         string `json:"image_url"`
	Description      string `json:"description"`
	FollowerNum      int    `json:"follower_num"`
	IsLive           bool   `json:"is_live"`
	IsParty          bool   `json:"is_party"`
	NextLiveSchedule int    `json:"next_live_schedule"`
}

type StreamingURL struct {
	IsDefault bool   `json:"is_default"`
	URL       string `json:"url"`
	Label     string `json:"label"`
	Type      string `json:"type"`
	ID        int    `json:"id"`
	Quality   int    `json:"quality"`
}

type LiveRoom struct {
	RoomURLKey       string         `json:"room_url_key"`
	OfficialLv       int            `json:"official_lv"`
	FollowerNum      int            `json:"follower_num"`
	StartedAt        int            `json:"started_at"`
	LiveID           int            `json:"live_id"`
	ImageSquare      string         `json:"image_square"`
	IsLive           bool           `json:"is_live"`
	IsFollow         bool           `json:"is_follow"`
	StreamingURLList []StreamingURL `json:"streaming_url_list"`
	LiveType         int            `json:"live_type"`
	Tags             []any          `json:"tags"`
	Image            string         `json:"image"`
	ViewNum          int            `json:"view_num"`
	GenreID          int            `json:"genre_id"`
	MainName         string         `json:"main_name"`
	PremiumRoomType  int            `json:"premium_room_type"`
	CellType         int            `json:"cell_type"`
	BcsvrKey         string         `json:"bcsvr_key"`
	RoomID           int            `json:"room_id"`
}
type Onlives struct {
	GenreID     int        `json:"genre_id"`
	Banners     []any      `json:"banners"`
	HasUpcoming bool       `json:"has_upcoming"`
	GenreName   string     `json:"genre_name"`
	Lives       []LiveRoom `json:"lives"`
}

type CustomRoom struct {
	Nick   string
	RoomId int
}
