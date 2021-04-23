package models

import (
	"reflect"
	"regexp"
	"strings"
)

// Language is a type describing an interface language.
type Language struct {
	Name      string
	RawName   string
	RawObject RawLanguage
	Prefix    string
}

// RawLanguage is the raw JSON file as a Go struct.
type RawLanguage struct {
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
	ServerVersionUsed     string `json:"serverversionused"`
	SetVersionUsage       string `json:"setversionusage"`
	SetServerVersionUsage string `json:"setserverversionusage"`
	InfoUsage             string `json:"infousage"`
	ListVersionUsage      string `json:"listversionusage"`

	SetLanguageSuccess     string `json:"setlanguagesuccess"`
	SetLanguageFail        string `json:"setlanguagefail"`
	LanguageUsed           string `json:"languageused"`
	SetLanguageUsage       string `json:"setlanguageusage"`
	SetServerLanguageUsage string `json:"setserverlanguageusage"`
	ListLanguageUsage      string `json:"listlanguageusage"`

	PrefixOneChar      string `json:"prefixonechar"`
	PrefixSuccess      string `json:"prefixsuccess"`
	ServerPrefixUsed   string `json:"serverprefixused"`
	ServerBracketsUsed string `json:"serverbracketsused"`
	SetPrefixUsage     string `json:"setprefixsusage"`
	SetBracketUsage    string `json:"setbracketsusage"`

	SetVerseNumbersUsage string `json:"setversenumbersusage"`
	SetHeadingsUsage     string `json:"setheadingsusage"`
	SetDisplayUsage      string `json:"setdisplayusage"`
	VerseNumbers         string `json:"versenumbers"`
	VerseNumbersSuccess  string `json:"versenumberssuccess"`
	Headings             string `json:"headings"`
	HeadingsSuccess      string `json:"headingssuccess"`
	OtherFormatFail      string `json:"otherformatfail"`
	Formatting           string `json:"formatting"`
	FormattingSuccess    string `json:"formattingsuccess"`
	FormattingFail       string `json:"formattingfail"`

	FailedSearch     string `json:"failedsearch"`
	FailedPreference string `json:"failedpreference"`

	NoServerPerm             string `json:"noserverperm"`
	SetServerVersionSuccess  string `json:"setserverversionsuccess"`
	SetServerVersionFail     string `json:"setserverversionfail"`
	SetServerLanguageSuccess string `json:"setserverlanguagesuccess"`
	SetServerLanguageFail    string `json:"setserverlanguagefail"`
	ServerLanguageUsed       string `json:"serverlanguageused"`
	SetServerBracketsSuccess string `json:"setserverbracketssuccess"`
	ServerBracketsFail       string `json:"serverbracketsfail"`

	SetVOTDTimeSuccess   string `json:"setvotdtimesuccess"`
	VOTDTimeUsed         string `json:"votdtimeused"`
	NoVOTDTimeUsed       string `json:"novotdtimeused"`
	ClearVOTDTimeSuccess string `json:"clearvotdtimesuccess"`
	VOTD                 string `json:"votd"`
	VOTDCantProcess      string `json:"votdcantprocess"`

	VersionInfo       string `json:"versioninfo"`
	VersionInfoFailed string `json:"versioninfofailed"`

	ExpectedParameter string `json:"expectedparameter"`
	DonutSpam         string `json:"donutspam"`

	ShardCount      string `json:"shardcount"`
	CachedServers   string `json:"cachedservers"`
	CachedChannels  string `json:"cachedchannels"`
	CachedUsers     string `json:"cachedusers"`
	PreferenceCount string `json:"preferencecount"`
	ServerPrefCount string `json:"serverprefcount"`
	VersionCount    string `json:"versioncount"`
	LanguageCount   string `json:"languagecount"`
	RunningOn       string `json:"runningon"`

	QueryTooShort      string `json:"queryTooShort"`
	SearchResults      string `json:"searchResults"`
	NothingFound       string `json:"nothingFound"`
	PageOf             string `json:"pageOf"`
	SearchNotSupported string `json:"searchNotSupported"`

	PlsWait         string `json:"plswait"`
	Supporters      string `json:"supporters"`
	AnonymousDonors string `json:"anonymousDonors"`
	DonorsNotListed string `json:"donorsNotListed"`

	AddVersionSuccess string `json:"addversionsuccess"`

	VerseError     string `json:"verseerror"`
	InvalidSection string `json:"invalidsection"`

	Enabled  string `json:"enabled"`
	Disabled string `json:"disabled"`

	Author     string `json:"author"`
	SinglePage string `json:"singlePage"`
	Pages      string `json:"pages"`
	Category   string `json:"category"`
	Section    string `json:"section"`
	Sections   string `json:"sections"`

	Commands  cmds `json:"commands"`
	Arguments args `json:"arguments"`
}

type cmds struct {
	Search          string `json:"search"`
	Version         string `json:"version"`
	Set             string `json:"set"`
	SetServer       string `json:"setserver"`
	Setup           string `json:"setup"`
	List            string `json:"list"`
	Info            string `json:"info"`
	Status          string `json:"status"`
	DailyVerse      string `json:"dailyverse"`
	Random          string `json:"random"`
	TrueRandom      string `json:"truerandom"`
	Formatting      string `json:"formatting"`
	Stats           string `json:"stats"`
	BibleBot        string `json:"biblebot"`
	AddVersion      string `json:"addversion"`
	RMVersion       string `json:"rmversion"`
	SetVerseNumbers string `json:"setversenumbers"`
	SetHeadings     string `json:"setheadings"`
	SetDisplay      string `json:"setdisplay"`
	Echo            string `json:"echo"`
	Leave           string `json:"leave"`
	Language        string `json:"language"`
	SetPrefix       string `json:"setprefix"`
	SetBrackets     string `json:"setbrackets"`
	Clear           string `json:"clear"`
	Nicene          string `json:"nicene"`
	Nicene325       string `json:"nicene325"`
	Apostles        string `json:"apostles"`
	Chalcedon       string `json:"chalcedon"`
	Creeds          string `json:"creeds"`
	Resources       string `json:"resources"`
	Announcements   string `json:"announcements"`
	Announce        string `json:"announce"`
	Misc            string `json:"misc"`
	Invite          string `json:"invite"`
	Supporters      string `json:"supporters"`
}

type args struct {
	Yes     string `json:"yes"`
	No      string `json:"no"`
	Enable  string `json:"enable"`
	Disable string `json:"disable"`
	True    string `json:"true"`
	False   string `json:"false"`
}

func (lng Language) GetString(str string) string {
	return lng.TranslatePlaceholdersInString(lng.GetRawString(str))
}

func (lng Language) GetRawString(str string) string {
	v := reflect.ValueOf(lng.RawObject)

	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Name == str && !stringInSlice(str, []string{"Commands", "Arguments"}) {
			return v.Field(i).String()
		}
	}

	return str
}

func (lng Language) TranslatePlaceholdersInString(str string) string {
	placeholderRegex, _ := regexp.Compile("<([^<>]+)>")
	placeholders := placeholderRegex.FindAllString(str, -1)

	for _, placeholder := range placeholders {
		tmpPlaceholder := placeholder

		switch placeholder {
		case "<+>":
			placeholder = lng.Prefix
		case "<truerandom>":
			placeholder = lng.GetCommandTranslation("TrueRandom")
		case "<dailyverse>":
			placeholder = lng.GetCommandTranslation("DailyVerse")
		default:
			purifiedQuery := strings.Title(placeholder[1 : len(placeholder)-1])
			possibleCommand := lng.GetCommandTranslation(purifiedQuery)
			possibleArgument := lng.GetArgumentTranslation(purifiedQuery)

			if possibleCommand != purifiedQuery {
				placeholder = possibleCommand
			} else if possibleArgument != purifiedQuery {
				placeholder = possibleArgument
			}
		}

		str = strings.ReplaceAll(str, tmpPlaceholder, placeholder)
	}

	return str
}

func (lng Language) GetCommandByTranslation(str string) string {
	v := reflect.ValueOf(lng.RawObject.Commands)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == str {
			return strings.ToLower(v.Type().Field(i).Name)
		}
	}

	return str
}

func (lng Language) GetCommandTranslation(str string) string {
	v := reflect.ValueOf(lng.RawObject.Commands)

	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Name == str {
			return v.Field(i).String()
		}
	}

	return str
}

func (lng Language) GetArgumentByTranslation(str string) string {
	v := reflect.ValueOf(lng.RawObject.Arguments)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == str {
			return strings.ToLower(v.Type().Field(i).Name)
		}
	}

	return str
}

func (lng Language) GetArgumentTranslation(str string) string {
	v := reflect.ValueOf(lng.RawObject.Arguments)

	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Name == str {
			return v.Field(i).String()
		}
	}

	return str
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
