package models

// Language is a type describing an interface language.
type Language struct {
	Name      string
	RawName   string
	RawObject rawLanguage
}

type rawLanguage struct {
	BibleBot string `json:"biblebot"`
	Credit   string `json:"credit"`

	CommandList string `json:"commandlist"`
	Usage       string `json:"usage"`

	CommandListName string `json:"commandlistName"`
	Links           string `json:"links"`

	Website    string `json:"website"`
	Copyrights string `json:"copyrights"`
	Code       string `json:"code"`
	Server     string `json:"server"`
	Terms      string `json:"terms"`

	CreedsText string `json:"creeds_text"`

	ApostlesName  string `json:"apostles_name"`
	Nicene325Name string `json:"nicene325_name"`
	NiceneName    string `json:"nicene_name"`
	ChalcedonName string `json:"chalcedon_name"`

	CatechismsText string `json:"catechisms_text"`

	Catechisms  string `json:"catechisms"`
	Confessions string `json:"confessions"`

	Protestant       string `json:"protestant"`
	Lutheran         string `json:"lutheran"`
	Reformed         string `json:"reformed"`
	Baptist          string `json:"baptist"`
	Catholic         string `json:"catholic"`
	EasternOrthodox  string `json:"eastern_orthodox"`
	OrientalOrthodox string `json:"oriental_orthodox"`

	PubDomain string `json:"pubdomain"`

	InvalidRange   string `json:"invalidrange"`
	Error          string `json:"error"`
	PassageTooLong string `json:"passagetoolong"`

	Subcommands string `json:"subcommands"`

	SetVersionSuccess     string `json:"setversionsuccess"`
	SetVersionFail        string `json:"setversionfail"`
	VersionUsed           string `json:"versionused"`
	GuildVersionUsed      string `json:"guildversionused"`
	SetVersionUsage       string `json:"setversionusage"`
	SetServerVersionUsage string `json:"setserverversionusage"`
	InfoUsage             string `json:"infousage"`
	ListVersionUsage      string `json:"listversionusage"`

	// TODO finish
}
