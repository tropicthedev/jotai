package config

type Applications struct {
	Timeout   int      `koanf:"timeout"`
	Channel   string   `koanf:"channel"`
	Questions []string `koanf:"questions"`
}

type Interviews struct {
	Notification string `koanf:"notification"`
	Role         string `koanf:"role"`
	Channel      string `koanf:"channel"`
	Private      bool   `koanf:"private"`
}

type Inactivity struct {
	Message                       string `koanf:"message"`
	VacationRole                  string `koanf:"vacation_role"`
	RemoveInactivePlayerAfterDays int    `koanf:"remove_inactive_player_after_days"`
	GracePeriodDays               int    `koanf:"grace_period_days"`
	Timezone                      string `koanf:"timezone"`
	Cron                          string `koanf:"cron"`
}

type WhitelistManager struct {
	Enabled    bool       `koanf:"enabled"`
	Inactivity Inactivity `koanf:"inactivity"`
}

type Discord struct {
	Prefix           string           `koanf:"prefix"`
	Name             string           `koanf:"name"`
	ConsoleChannel   string           `koanf:"console_channel"`
	Token            string           `koanf:"token"`
	GuildID          string           `koanf:"guild_id"`
	ClientID         string           `koanf:"client_id"`
	Status           string           `koanf:"status"`
	JoinChannel      string           `koanf:"join_channel"`
	JoinMessage      string           `koanf:"join_message"`
	AcceptChannel    string           `koanf:"accept_channel"`
	AcceptMessage    string           `koanf:"accept_message"`
	MemberRole       string           `koanf:"member_role"`
	AdminRole        string           `koanf:"admin_role"`
	Owners           []string         `koanf:"owners"`
	Server           Server           `koanf:"server"`
	Applications     Applications     `koanf:"applications"`
	Interviews       Interviews       `koanf:"interviews"`
	WhitelistManager WhitelistManager `koanf:"whitelist_manager"`
}

type Server struct {
	Port  int    `koanf:"port"`
	Token string `koanf:"token"`
}

type Config struct {
	Discord Discord `koanf:"discord"`
}

func New(discord Discord) Config {
	return Config{
		Discord: discord,
	}
}
