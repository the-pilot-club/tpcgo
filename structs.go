package tpcgo

import "github.com/bwmarrin/discordgo"

type Session struct {
	FCPSession     *FCPSession
	CoreAPISession *CoreAPISession
	VATSIMSession  *VATSIMSession
}

type SessionConfig struct {
	FCPKey       string
	FCPEnv       string
	VATSIMAPIKey string
	CoreApiKey   string
}

type FCPSession struct {
	ApiKeyHeader string
	ApiKey       string
	UserAgent    string
	Environment  string
}

type CoreAPISession struct {
	ApiKeyHeader string
	ApiKey       string
	UserAgent    string
}

type VATSIMSession struct {
	ApiKeyHeader string
	ApiKey       string
	UserAgent    string
}

type NewSuggestion struct {
	StatusID  int    `json:"statusId"`
	DiscordID string `json:"discordId"`
	IdeaText  string `json:"ideaText"`
}

type Suggestions struct {
	ID           string `json:"id,omitempty"`
	StatsID      int    `json:"statsId,omitempty"`
	DiscordId    string `json:"discordId,omitempty"`
	IdeaText     string `json:"ideaText,omitempty"`
	Reason       string `json:"reason,omitempty"`
	MessageID    string `json:"messageId,omitempty"`
	ChannelID    string `json:"channelId,omitempty"`
	ActionUserID string `json:"actionUserId,omitempty"`
}

type QuizQuestion struct {
	ID                string `json:"id,omitempty"`
	Question          string `json:"question,omitempty"`
	OptionA           string `json:"optionA,omitempty"`
	OptionB           string `json:"optionB,omitempty"`
	OptionC           string `json:"optionC,omitempty"`
	CorrectAnswer     string `json:"correctAnswer,omitempty"`
	AnswerDescription string `json:"answerDescription,omitempty"`
	IsActive          bool   `json:"isActive,omitempty"`
	IsOld             bool   `json:"isOld,omitempty"`
	CreatedAt         string `json:"createdAt,omitempty"`
	UpdatedAt         string `json:"updatedAt,omitempty"`
}

type QuizQuestionForResponse struct {
	MessageID  string `json:"messageId,omitempty"`
	QuestionID string `json:"questionId,omitempty"`
}

type QuizUserResponseSet struct {
	MessageID string            `json:"messageId,omitempty"`
	UserID    string            `json:"userId,omitempty"`
	Answer    string            `json:"answer,omitempty"`
	User      *discordgo.Member `json:"user,omitempty"`
}

type QuizUserResponse struct {
	User *discordgo.Member `json:"user,omitempty"`
}

type NotFoundResource struct {
	Detail string `json:"detail,omitempty"`
}

type AuditLogEntry struct {
	UserId      string `json:"user_id,omitempty"`
	SubmittedBy string `json:"submitted_by,omitempty"`
	Text        string `json:"text,omitempty"`
}

type FCPAuditLogAdd struct {
	UserID      string `json:"user_id"`
	SubmittedBy string `json:"submitted_by"`
	Text        string `json:"text"`
}

type AuditLogResponse struct {
	Success string        `json:"success,omitempty"`
	Entry   AuditLogEntry `json:"entry,omitempty"`
}

type SuccessResponse struct {
	Success string `json:"success,omitempty"`
}

type FCPCallsignResponse struct {
	Callsign int `json:"tpcCallsign,omitempty"`
}

type FCPFullUser struct {
	ID              string            `json:"id"`
	DiscordID       string            `json:"discordId,omitempty"`
	DiscordUsername string            `json:"discordUsername,omitempty"`
	FirstName       string            `json:"firstName,omitempty"`
	LastName        string            `json:"lastName,omitempty"`
	Callsign        int               `json:"callsign,omitempty"`
	VATSIMCid       int               `json:"vatsimCid,omitempty"`
	Email           string            `json:"email,omitempty"`
	HomeAirport     string            `json:"homeAirport,omitempty"`
	ChartersCode    string            `json:"chartersCode,omitempty"`
	Bio             string            `json:"bio,omitempty"`
	Status          string            `json:"status,omitempty"`
	AircraftHangar  []FCPUserAircraft `json:"aircraftHangar,omitempty"`
}

type FCPLimitedUser struct {
	ID             string            `json:"id,omitempty"`
	DiscordID      string            `json:"discordId,omitempty"`
	Callsign       int               `json:"callsign,omitempty"`
	VATSIMCid      int               `json:"vatsimCid,omitempty"`
	HomeAirport    string            `json:"homeAirport,omitempty"`
	ChartersCode   string            `json:"chartersCode,omitempty"`
	Bio            string            `json:"bio,omitempty"`
	AircraftHangar []FCPUserAircraft `json:"aircraftHangar,omitempty"`
}

type FCPUserAdd struct {
	UserID string `json:"id"`
}

type FCPUserAircraft struct {
	ICAO string `json:"icao,omitempty"`
	Name string `json:"aircraftName,omitempty"`
}

type FCPFBOs struct {
	FBO []FBO `json:"fbos,omitempty"`
}
type FCPSectors struct {
	Sectors []Sectors `json:"sectors,omitempty"`
}
type FBO struct {
	ICAO        string    `json:"icao,omitempty"`
	Name        string    `json:"name,omitempty"`
	Region      string    `json:"region,omitempty"`
	Fuel        string    `json:"fuel,omitempty"`
	Focus       string    `json:"focus,omitempty"`
	PublicNotes string    `json:"publicNotes,omitempty"`
	Sectors     []Sectors `json:"sectors,omitempty"`
}

type Sectors struct {
	StartICAO    string `json:"startIcao,omitempty"`
	EndICAO      string `json:"endIcao,omitempty"`
	Size         string `json:"size,omitempty"`
	SectorNumber int    `json:"sectorNumber,omitempty"`
}

type DataFeed struct {
	General     *GeneralInfo `json:"general,omitempty"`
	Pilots      []Pilot      `json:"pilots,omitempty"`
	Controllers []Controller `json:"controllers,omitempty"`
	ATISs       []ATIS       `json:"atis,omitempty"`
}

type GeneralInfo struct {
	Version          int    `json:"version,omitempty"`
	Reload           int    `json:"reload,omitempty"`
	Update           string `json:"update,omitempty"`
	UpdateTimestamp  string `json:"update_timestamp,omitempty"`
	ConnectedClients int    `json:"connected_clients,omitempty"`
	UniqueUsers      int    `json:"unique_users,omitempty"`
}

type Pilot struct {
	CID            int         `json:"cid,omitempty"`
	Name           string      `json:"name,omitempty"`
	Callsign       string      `json:"callsign,omitempty"`
	Server         string      `json:"server,omitempty"`
	PilotRating    int         `json:"pilot_rating,omitempty"`
	MilitaryRating int         `json:"military_rating,omitempty"`
	Latitude       float64     `json:"latitude,omitempty"`
	Longitude      float64     `json:"longitude,omitempty"`
	Altitude       int         `json:"altitude,omitempty"`
	Groundspeed    int         `json:"groundspeed,omitempty"`
	Transponder    string      `json:"transponder,omitempty"`
	Heading        int         `json:"heading,omitempty"`
	QnhIHg         float64     `json:"qnh_i_hg,omitempty"`
	QnhMb          int         `json:"qnh_mb,omitempty"`
	FlightPlan     *FlightPlan `json:"flight_plan,omitempty"`
	LogonTime      string      `json:"logon_time,omitempty"`
	LastUpdated    string      `json:"last_updated,omitempty"`
}

type FlightPlan struct {
	FlightRules         string `json:"flight_rules,omitempty"`
	Aircraft            string `json:"aircraft,omitempty"`
	AircraftFAA         string `json:"aircraft_faa,omitempty"`
	AircraftShort       string `json:"aircraft_short,omitempty"`
	Departure           string `json:"departure,omitempty"`
	Arrival             string `json:"arrival,omitempty"`
	Alternate           string `json:"alternate,omitempty"`
	CruiseTAS           string `json:"cruise_tas,omitempty"`
	Altitude            string `json:"altitude,omitempty"`
	DepTime             string `json:"deptime,omitempty"`
	EnrouteTime         string `json:"enroute_time,omitempty"`
	FuelTime            string `json:"fuel_time,omitempty"`
	Remarks             string `json:"remarks,omitempty"`
	Route               string `json:"route,omitempty"`
	RevisionID          int    `json:"revision_id,omitempty"`
	AssignedTransponder string `json:"assigned_transponder,omitempty"`
}

type Controller struct {
	CID         int      `json:"cid,omitempty"`
	Name        string   `json:"name,omitempty"`
	Callsign    string   `json:"callsign,omitempty"`
	Frequency   string   `json:"frequency,omitempty"`
	Facility    int      `json:"facility,omitempty"`
	Rating      int      `json:"rating,omitempty"`
	Server      string   `json:"server,omitempty"`
	VisualRange int      `json:"visual_range,omitempty"`
	TextAtis    []string `json:"text_atis,omitempty"`
	LogonTime   string   `json:"logon_time,omitempty"`
}

type ATIS struct {
	CID         int      `json:"cid,omitempty"`
	Name        string   `json:"name,omitempty"`
	Callsign    string   `json:"callsign,omitempty"`
	Frequency   string   `json:"frequency,omitempty"`
	Facility    int      `json:"facility,omitempty"`
	Rating      int      `json:"rating,omitempty"`
	ATISCode    string   `json:"atis_code,omitempty"`
	Server      string   `json:"server,omitempty"`
	VisualRange int      `json:"visual_range,omitempty"`
	TextAtis    []string `json:"text_atis,omitempty"`
	LogonTime   string   `json:"logon_time,omitempty"`
}

type PreFile struct {
	CID        int         `json:"cid,omitempty"`
	Name       string      `json:"name,omitempty"`
	Callsign   string      `json:"callsign,omitempty"`
	FlightPlan *FlightPlan `json:"flight_plan,omitempty"`
	LastUpdate string      `json:"last_updated,omitempty"`
}

type MemberDiscord struct {
	ID        string `json:"id,omitempty"`
	DiscordID string `json:"user_id,omitempty"`
}

type VATSIMUserHours struct {
	ID    int     `json:"id,omitempty"`
	ATC   float64 `json:"atc,omitempty"`
	Pilot float64 `json:"pilot,omitempty"`
	S1    float64 `json:"s1,omitempty"`
	S2    float64 `json:"s2,omitempty"`
	S3    float64 `json:"s3,omitempty"`
	C1    float64 `json:"c1,omitempty"`
	C2    float64 `json:"c2,omitempty"`
	C3    float64 `json:"c3,omitempty"`
	I1    float64 `json:"i1,omitempty"`
	I2    float64 `json:"i2,omitempty"`
	I3    float64 `json:"i3,omitempty"`
	SUP   float64 `json:"sup,omitempty"`
	ADM   float64 `json:"adm,omitempty"`
}
