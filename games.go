package SteamCommon

type SteamUserGame struct {
	AppID           uint64 `json:"appid"`
	PlaytimeForever uint64 `json:"playtime_forever"`

	Name           string `json:"name"`
	ImgIconUrl     string `json:"img_icon_url"`
	ImgLogoUrl     string `json:"img_logo_url"`
	CommunityStats bool   `json:"has_community_visible_stats"`
}
