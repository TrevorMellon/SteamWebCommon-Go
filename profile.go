package SteamCommon

import (
	"time"
)

type SteamID uint64

type SteamUserProfile struct {
	SteamID         SteamID
	PersonaName     string
	Realname        string
	DisplayPersona  string
	DisplayRealname string
	SummaryParsed   time.Time
	DateAdded       time.Time
	Friends         []SteamUserProfile
	FriendSince     time.Time
	FriendSinceStr  string
	Avatar          SteamAvatar
	Url             string
	Status          OnlineStatus
	ProfileType     SteamProfileType
	StatusColor     string
}

type SteamAvatar struct {
	Small  string
	Medium string
	Large  string
}

type OnlineStatus int

const (
	StatusUnknown OnlineStatus = iota
	StatusOnline
	StatusOffline
	StatusBusy
	StatusAway
	StatusLookingToTrade
	StatusLookingToGame
)

type SteamProfileType int

/*
	An enumeration of profile types
*/
const (
	ProfileUnknown SteamProfileType = iota
	ProfilePublic
	ProfilePrivate
	ProfileSemi
)
