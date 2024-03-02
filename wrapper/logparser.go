package wrapper

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pawelk1337/mcsrv/wrapper/events"
)

// LogParser is an interface func to decode any server log line
// to its respective event type. The returned events must be either:
// - Cmd: event holds data to be returned to a user command.
// - State: event affects the state of the wrapper.
// - Game: event related to in-game events, like a player died...
// - Nil: event that hold no value and usually ignored/
type LogParser func(string, int) (events.Event, events.EventType)

type logLine struct {
	timestamp  string
	threadName string
	level      string
	output     string
}

// var logRegex = regexp.MustCompile(`(\[[0-9:]*\]) \[([A-z(-| )#0-9]*)\/([A-z #]*)\]: (.*)`)
var logRegex = regexp.MustCompile(`\[(\d{2}:\d{2}:\d{2}) (INFO|WARN|ERROR|FATAL)\]: (.*)`)

func ParseToLogLine(line string) *logLine {
	matches := logRegex.FindAllStringSubmatch(line, 4)
	if len(matches) == 0 {
		return &logLine{}
	}
	return &logLine{
		timestamp: matches[0][1],
		level:     matches[0][2],
		output:    matches[0][3],
	}
}

var stateEventToRegexp = map[string]*regexp.Regexp{
	events.Started:  regexp.MustCompile(`^Done (?s)(.*)! For help`),
	events.Starting: regexp.MustCompile(`^Starting Minecraft server on (.*)`),
	events.Stopping: regexp.MustCompile(`^Stopping (.*) server`),
	events.Saving:   regexp.MustCompile(`^Saving the game`),
	events.Saved:    regexp.MustCompile(`^Saved (?s)(.*)`),
}

var gameEventToRegex = map[string]*regexp.Regexp{
	events.Banned:          regexp.MustCompile(`^Banned (?s)(.*): (?s)(.*)`),
	events.BanList:         regexp.MustCompile(`^There are (no|\d+) bans(:|\z)`),
	events.BanListEntry:    regexp.MustCompile(`(?s)(.*) was banned by Server: (.*)`),
	events.DataGet:         regexp.MustCompile(`(?s)(.*) has the following (entity|block|storage) data: (.*)`),
	events.DataGetNoEntity: regexp.MustCompile(`^No (entity|block|storage) was found`),
	events.DefaultGameMode: regexp.MustCompile(`^The default game mode is now (Survival|Creative|Adventure|Spectator) Mode`),
	events.Difficulty:      regexp.MustCompile(`^The difficulty (?s)(.*)`),
	events.ExperienceAdd:   regexp.MustCompile(`^Gave ([0-9]+) experience (levels|points) to (?s)(.*)`),
	events.ExperienceQuery: regexp.MustCompile(`(?s)(.*) has ([0-9]+) experience (levels|points)`),
	events.Give:            regexp.MustCompile(`^Gave ([0-9]+) \[(?s)(.*) (?s)(.*)\] to (?s)(.*)`),
	events.NoPlayerFound:   regexp.MustCompile(`^No player was found`),
	// TODO: There is an insane amount of death messages: https://minecraft.gamepedia.com/Death_messages, support all?
	events.PlayerDied:       regexp.MustCompile(`(?s)(.*) (was shot|was pummeled|drowned|blew up|was blown up|was killed by|hit the ground|fell|was slain|suffocated)(.*)`),
	events.PlayerJoined:     regexp.MustCompile(`(?s)(.*) joined the game`),
	events.PlayerLeft:       regexp.MustCompile(`(?s)(.*) left the game`),
	events.PlayerUUID:       regexp.MustCompile(`^UUID of player (?s)(.*) is (?s)(.*)`),
	events.PlayerSay:        regexp.MustCompile(`<(?s)(.*)> (?s)(.*)`),
	events.Kicked:           regexp.MustCompile(`^Kicked (?s)(.*): (.*)`),
	events.Seed:             regexp.MustCompile(`^Seed: (.*)`),
	events.ServerOverloaded: regexp.MustCompile(`^Can't keep up! Is the server overloaded\? Running ([0-9]+)ms or ([0-9]+) ticks behind`),
	events.TimeIs:           regexp.MustCompile(`^The time is (?s)(.*)`),
	events.UnknownItem:      regexp.MustCompile(`^Unknown item (.*)`),
	events.Version:          regexp.MustCompile(`^Starting minecraft server version (.*)`),
	events.WhisperTo:        regexp.MustCompile(`^You whisper to (?s)(.*): (.*)`),
}

var activeGameEvents = map[string]*regexp.Regexp{
	events.PlayerDied:       gameEventToRegex[events.PlayerDied],
	events.PlayerJoined:     gameEventToRegex[events.PlayerJoined],
	events.PlayerLeft:       gameEventToRegex[events.PlayerLeft],
	events.PlayerUUID:       gameEventToRegex[events.PlayerUUID],
	events.PlayerSay:        gameEventToRegex[events.PlayerSay],
	events.ServerOverloaded: gameEventToRegex[events.ServerOverloaded],
	events.TimeIs:           gameEventToRegex[events.TimeIs],
	events.Version:          gameEventToRegex[events.Version],
}

func RegisterGameEvent(ev string) {
	_, ok := activeGameEvents[ev]
	if !ok {
		activeGameEvents[ev] = gameEventToRegex[ev]
	}
}

func LogParserFunc(line string, tick int) (events.Event, events.EventType) {
	remove_ansi := regexp.MustCompile(`\x1B(?:[@-Z\\-_]|\[[0-?]*[ -/]*[@-~])`)
	line = remove_ansi.ReplaceAllString(line, "")

	ll := ParseToLogLine(line)
	if ll.output == "" {
		return events.NilEvent, events.TypeNil
	}

	for e, reg := range stateEventToRegexp {
		if reg.MatchString(ll.output) {
			return events.NewStateEvent(e), events.TypeState
		}
	}
	for e, reg := range activeGameEvents {
		matches := reg.FindStringSubmatch(ll.output)
		if matches == nil {
			continue
		}
		switch e {
		case events.BanList:
			return HandleBanList(matches)
		case events.BanListEntry:
			return HandleBanListEntry(matches)
		case events.Difficulty:
			return HandleDifficulty(matches)
		case events.ExperienceQuery:
			return HandleExperienceQuery(matches)
		case events.PlayerJoined:
			return HandlePlayerJoined(matches, tick)
		case events.PlayerLeft:
			return HandlePlayerLeft(matches, tick)
		case events.PlayerDied:
			return HandlePlayerDied(matches, tick)
		case events.PlayerUUID:
			return HandlePlayerUUIDEvent(matches, tick)
		case events.PlayerSay:
			return HandlePlayerSayEvent(matches, tick)
		case events.Version:
			return HandleVersionEvent(matches)
		case events.TimeIs:
			return HandleTimeEvent(matches)
		case events.DataGet:
			return HandleDataGet(matches)
		case events.DataGetNoEntity:
			return HandleDataGetNoEntity(matches)
		case events.Seed:
			return HandleSeed(matches)
		case events.ServerOverloaded:
			return HandleServerOverloaded(matches, tick)
		case events.DefaultGameMode:
			return HandleDefaultGameMode(matches)
		case events.Banned:
			return HandleBanned(matches)
		case events.WhisperTo, events.ExperienceAdd, events.Give, events.NoPlayerFound,
			events.Kicked, events.UnknownItem:
			return events.NewGameEvent(e), events.TypeCmd
		default:
			gameEvent := events.NewGameEvent(e)
			gameEvent.Tick = tick
			return gameEvent, events.TypeGame
		}
	}
	return events.NilEvent, events.TypeNil
}

func HandleBanList(matches []string) (events.GameEvent, events.EventType) {
	blEvent := events.NewGameEvent(events.BanList)
	blEvent.Data = map[string]string{
		"entry_type": "header",
	}
	if matches[1] != "no" {
		// This indicates that there are entries to report back...
		blEvent.Data["entry_count"] = matches[1]
	}
	return blEvent, events.TypeCmd
}

func HandleBanListEntry(matches []string) (events.GameEvent, events.EventType) {
	bleEvent := events.NewGameEvent(events.BanList)
	bleEvent.Data = map[string]string{
		"entry_type": "item",
		"entry_name": matches[1],
		"reason":     matches[2],
	}
	return bleEvent, events.TypeCmd
}

func HandleDifficulty(matches []string) (events.GameEvent, events.EventType) {
	dfEvent := events.NewGameEvent(events.Difficulty)
	dfEvent.Data = map[string]string{}
	if strings.Contains(matches[1], "did not change") {
		dfEvent.Data["error_message"] = matches[0]
	}
	return dfEvent, events.TypeCmd
}

func HandleExperienceQuery(matches []string) (events.GameEvent, events.EventType) {
	xqEvent := events.NewGameEvent(events.ExperienceQuery)
	xqEvent.Data = map[string]string{
		"amount": matches[2],
	}
	return xqEvent, events.TypeCmd
}

func HandlePlayerJoined(matches []string, tick int) (events.GameEvent, events.EventType) {
	pjEvent := events.NewGameEvent(events.PlayerJoined)
	pjEvent.Tick = tick
	pjEvent.Data = map[string]string{
		"player_name": matches[1],
	}
	return pjEvent, events.TypeGame
}

func HandlePlayerLeft(matches []string, tick int) (events.GameEvent, events.EventType) {
	plEvent := events.NewGameEvent(events.PlayerLeft)
	plEvent.Tick = tick
	plEvent.Data = map[string]string{
		"player_name": matches[1],
	}
	return plEvent, events.TypeGame
}

func HandlePlayerDied(matches []string, tick int) (events.GameEvent, events.EventType) {
	pdEvent := events.NewGameEvent(events.PlayerDied)
	pdEvent.Tick = tick
	pdEvent.Data = map[string]string{
		"player_name":   matches[1],
		"death_by":      matches[2],
		"death_details": "",
	}
	if len(matches) >= 4 {
		pdEvent.Data["death_details"] = matches[3]
	}
	return pdEvent, events.TypeGame
}

func HandlePlayerUUIDEvent(matches []string, tick int) (events.GameEvent, events.EventType) {
	puEvent := events.NewGameEvent(events.PlayerUUID)
	puEvent.Tick = tick
	puEvent.Data = map[string]string{
		"player_name": matches[1],
		"player_uuid": matches[2],
	}
	return puEvent, events.TypeGame
}

func HandlePlayerSayEvent(matches []string, tick int) (events.GameEvent, events.EventType) {
	psEvent := events.NewGameEvent(events.PlayerSay)
	psEvent.Tick = tick
	psEvent.Data = map[string]string{
		"player_name":    matches[1],
		"player_message": matches[2],
	}
	return psEvent, events.TypeGame
}

func HandleVersionEvent(matches []string) (events.GameEvent, events.EventType) {
	versionEvent := events.NewGameEvent(events.Version)
	versionEvent.Data = map[string]string{
		"version": matches[1],
	}
	return versionEvent, events.TypeCmd
}

func HandleTimeEvent(matches []string) (events.GameEvent, events.EventType) {
	tickStr := matches[1]
	tick, _ := strconv.Atoi(tickStr)
	timeEvent := events.NewGameEvent(events.TimeIs)
	timeEvent.Tick = tick
	return timeEvent, events.TypeCmd
}

func HandleDataGet(matches []string) (events.GameEvent, events.EventType) {
	dgEvent := events.NewGameEvent(events.DataGet)
	dgEvent.Data = map[string]string{
		"player_name": matches[1],
		"data_type":   matches[2],
		"data_raw":    matches[3],
	}
	return dgEvent, events.TypeCmd
}

func HandleDataGetNoEntity(matches []string) (events.GameEvent, events.EventType) {
	dgEvent := events.NewGameEvent(events.DataGet)
	dgEvent.Data = map[string]string{
		"error_message": matches[0],
	}
	return dgEvent, events.TypeCmd
}

func HandleSeed(matches []string) (events.GameEvent, events.EventType) {
	sdEvent := events.NewGameEvent(events.Seed)
	sdEvent.Data = map[string]string{
		"data_raw": matches[1],
	}
	return sdEvent, events.TypeCmd
}

func HandleServerOverloaded(matches []string, tick int) (events.GameEvent, events.EventType) {
	soEvent := events.NewGameEvent(events.ServerOverloaded)
	soEvent.Tick = tick
	soEvent.Data = map[string]string{
		"lag_time": matches[1],
		"lag_tick": matches[2],
	}
	return soEvent, events.TypeGame
}

func HandleDefaultGameMode(matches []string) (events.GameEvent, events.EventType) {
	gmEvent := events.NewGameEvent(events.DefaultGameMode)
	gmEvent.Data = map[string]string{
		"default_game_mode": matches[1],
	}
	return gmEvent, events.TypeGame
}

func HandleBanned(matches []string) (events.GameEvent, events.EventType) {
	bnEvent := events.NewGameEvent(events.Banned)
	bnEvent.Data = map[string]string{
		"player_name": matches[1],
		"reason":      matches[2],
	}
	return bnEvent, events.TypeGame
}
