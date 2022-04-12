package metadata1

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

type Metrics struct {
	El           int `json:"el"`
	Script       int `json:"script"`
	H            int `json:"h"`
	Batt         int `json:"batt"`
	Perf         int `json:"perf"`
	Auto         int `json:"auto"`
	Tz           int `json:"tz"`
	Fp2          int `json:"fp2"`
	Lsubid       int `json:"lsubid"`
	Browser      int `json:"browser"`
	Capabilities int `json:"capabilities"`
	Gpu          int `json:"gpu"`
	Dnt          int `json:"dnt"`
	Math         int `json:"math"`
	Tts          int `json:"tts"`
	Input        int `json:"input"`
	Canvas       int `json:"canvas"`
	Captchainput int `json:"captchainput"`
	Pow          int `json:"pow"`
}

type Interaction struct {
	Clicks                int           `json:"clicks"`
	Touches               int           `json:"touches"`
	KeyPresses            int           `json:"keyPresses"`
	Cuts                  int           `json:"cuts"`
	Copies                int           `json:"copies"`
	Pastes                int           `json:"pastes"`
	KeyPressTimeIntervals []interface{} `json:"keyPressTimeIntervals"`
	MouseClickPositions   []interface{} `json:"mouseClickPositions"`
	KeyCycles             []interface{} `json:"keyCycles"`
	MouseCycles           []interface{} `json:"mouseCycles"`
	TouchCycles           []interface{} `json:"touchCycles"`
}

type Scripts struct {
	DynamicUrls       []string `json:"dynamicUrls"`
	InlineHashes      []int    `json:"inlineHashes"`
	Elapsed           int      `json:"elapsed"`
	DynamicURLCount   int      `json:"dynamicUrlCount"`
	InlineHashesCount int      `json:"inlineHashesCount"`
}

type History struct {
	Length int `json:"length"`
}

type Performance struct {
	Timing Timing `json:"timing"`
}

type Timing struct {
	ConnectStart               int `json:"connectStart"`
	NavigationStart            int `json:"navigationStart"`
	LoadEventEnd               int `json:"loadEventEnd"`
	DomLoading                 int `json:"domLoading"`
	SecureConnectionStart      int `json:"secureConnectionStart"`
	FetchStart                 int `json:"fetchStart"`
	DomContentLoadedEventStart int `json:"domContentLoadedEventStart"`
	ResponseStart              int `json:"responseStart"`
	ResponseEnd                int `json:"responseEnd"`
	DomInteractive             int `json:"domInteractive"`
	DomainLookupEnd            int `json:"domainLookupEnd"`
	RedirectStart              int `json:"redirectStart"`
	RequestStart               int `json:"requestStart"`
	UnloadEventEnd             int `json:"unloadEventEnd"`
	UnloadEventStart           int `json:"unloadEventStart"`
	DomComplete                int `json:"domComplete"`
	DomainLookupStart          int `json:"domainLookupStart"`
	LoadEventStart             int `json:"loadEventStart"`
	DomContentLoadedEventEnd   int `json:"domContentLoadedEventEnd"`
	RedirectEnd                int `json:"redirectEnd"`
	ConnectEnd                 int `json:"connectEnd"`
}

type Automation struct {
	Phantom Phantom `json:"phantom"`
	Wd      Wd      `json:"wd"`
}

type Phantom struct {
	Properties PhantomProperties `json:"properties"`
}

type PhantomProperties struct {
	Window []interface{} `json:"window"`
}

type Wd struct {
	Properties WdProperties `json:"properties"`
}

type WdProperties struct {
	Document  []interface{} `json:"document"`
	Navigator []interface{} `json:"navigator"`
	Window    []interface{} `json:"window"`
}

type Capabilities struct {
	CSS     CSS `json:"css"`
	Js      Js  `json:"js"`
	Elapsed int `json:"elapsed"`
}

type CSS struct {
	TextShadow       int `json:"textShadow"`
	WebkitTextStroke int `json:"WebkitTextStroke"`
	BoxShadow        int `json:"boxShadow"`
	BorderRadius     int `json:"borderRadius"`
	BorderImage      int `json:"borderImage"`
	Opacity          int `json:"opacity"`
	Transform        int `json:"transform"`
	Transition       int `json:"transition"`
}

type Js struct {
	Audio        bool   `json:"audio"`
	Geolocation  bool   `json:"geolocation"`
	LocalStorage string `json:"localStorage"`
	Touch        bool   `json:"touch"`
	Video        bool   `json:"video"`
	WebWorker    bool   `json:"webWorker"`
}

type Gpu struct {
	Vendor     string   `json:"vendor"`
	Model      string   `json:"model"`
	Extensions []string `json:"extensions"`
}

type Math struct {
	Tan string `json:"tan"`
	Sin string `json:"sin"`
	Cos string `json:"cos"`
}

type Form struct {
	ApCredentialAutofillHint ApCredentialAutofillHint `json:"ap-credential-autofill-hint"`
	Password                 Password                 `json:"password"`
}

type ApCredentialAutofillHint struct {
	Clicks                int           `json:"clicks"`
	Touches               int           `json:"touches"`
	KeyPresses            int           `json:"keyPresses"`
	Cuts                  int           `json:"cuts"`
	Copies                int           `json:"copies"`
	Pastes                int           `json:"pastes"`
	KeyPressTimeIntervals []interface{} `json:"keyPressTimeIntervals"`
	MouseClickPositions   []interface{} `json:"mouseClickPositions"`
	KeyCycles             []interface{} `json:"keyCycles"`
	MouseCycles           []interface{} `json:"mouseCycles"`
	TouchCycles           []interface{} `json:"touchCycles"`
	Width                 int           `json:"width"`
	Height                int           `json:"height"`
	TotalFocusTime        int           `json:"totalFocusTime"`
	Checksum              string        `json:"checksum"`
	Autocomplete          bool          `json:"autocomplete"`
	Prefilled             bool          `json:"prefilled"`
}

type Password struct {
	Clicks                int           `json:"clicks"`
	Touches               int           `json:"touches"`
	KeyPresses            int           `json:"keyPresses"`
	Cuts                  int           `json:"cuts"`
	Copies                int           `json:"copies"`
	Pastes                int           `json:"pastes"`
	KeyPressTimeIntervals []interface{} `json:"keyPressTimeIntervals"`
	MouseClickPositions   []interface{} `json:"mouseClickPositions"`
	KeyCycles             []interface{} `json:"keyCycles"`
	MouseCycles           []interface{} `json:"mouseCycles"`
	TouchCycles           []interface{} `json:"touchCycles"`
	Width                 int           `json:"width"`
	Height                int           `json:"height"`
	TotalFocusTime        int           `json:"totalFocusTime"`
	Autocomplete          bool          `json:"autocomplete"`
	Prefilled             bool          `json:"prefilled"`
}

type Canvas struct {
	Hash          int           `json:"hash"`
	EmailHash     interface{}   `json:"emailHash"`
	HistogramBins []interface{} `json:"histogramBins"`
}

type Token struct {
	IsCompatible   bool `json:"isCompatible"`
	PageHasCaptcha int  `json:"pageHasCaptcha"`
}

type Auth struct {
	Form FormAuth `json:"form"`
}

type FormAuth struct {
	Method string `json:"method"`
}

type Payload struct {
	Metrics      Metrics       `json:"metrics"`
	Start        int           `json:"start"`
	Interaction  Interaction   `json:"interaction"`
	Scripts      Scripts       `json:"scripts"`
	History      History       `json:"history"`
	Battery      map[int]int   `json:"battery"`
	Performance  Performance   `json:"performance"`
	Automation   Automation    `json:"automation"`
	End          int           `json:"end"`
	TimeZone     int           `json:"timeZone"`
	FlashVersion interface{}   `json:"flashVersion"`
	Plugins      string        `json:"plugins"`
	DupedPlugins string        `json:"dupedPlugins"`
	ScreenInfo   string        `json:"screenInfo"`
	Referrer     string        `json:"referrer"`
	UserAgent    string        `json:"userAgent"`
	Location     string        `json:"location"`
	WebDriver    bool          `json:"webDriver"`
	Capabilities Capabilities  `json:"capabilities"`
	Gpu          Gpu           `json:"gpu"`
	Dnt          interface{}   `json:"dnt"`
	Math         Math          `json:"math"`
	TimeToSubmit int           `json:"timeToSubmit"`
	Form         Form          `json:"form"`
	Canvas       Canvas        `json:"canvas"`
	Token        Token         `json:"token"`
	Auth         Auth          `json:"auth"`
	Errors       []interface{} `json:"errors"`
	Version      string        `json:"version"`
	LsUbid       string        `json:"lsUbid"`
}

type BasePayload struct {
	Location  string
	Referrer  string
	Start     int
	UserAgent string
}

func randomInt(max int, min ...int) int {
	if len(min) == 0 {
		min = []int{0}
	}
	rint := rand.Intn(max - min[0])
	return min[0] + rint
}

func nowDate() int {
	return int(time.Now().UnixNano() / 1e6)
}

func NewPayload(base BasePayload) *Payload {
	rand.Seed(time.Now().UnixNano())

	return &Payload{
		Metrics: Metrics{
			El:           0,
			Script:       0,
			H:            0,
			Batt:         0,
			Perf:         0,
			Auto:         0,
			Tz:           0,
			Fp2:          0,
			Lsubid:       0,
			Browser:      0,
			Capabilities: 0,
			Gpu:          0,
			Dnt:          0,
			Math:         0,
			Tts:          0,
			Input:        0,
			Canvas:       0,
			Captchainput: 0,
			Pow:          0,
		},
		Start: base.Start,
		Interaction: Interaction{
			Clicks:                0,
			Touches:               0,
			KeyPresses:            0,
			Cuts:                  0,
			Copies:                0,
			Pastes:                0,
			KeyPressTimeIntervals: []interface{}{},
			MouseClickPositions:   []interface{}{},
			KeyCycles:             []interface{}{},
			MouseCycles:           []interface{}{},
			TouchCycles:           []interface{}{},
		},
		Scripts: Scripts{
			DynamicUrls: []string{
				"https://c.amazon-adsystem.com/bao-csm/forensics/a9-tq-forensics-incremental.min.js",
				"https://images-na.ssl-images-amazon.com/images/I/31YXrY93hfL.js",
				"https://images-na.ssl-images-amazon.com/images/I/61XKxrBtDVL._RC|11Y+5x+kkTL.js,51KMV3Cz2XL.js,31x4ENTlVIL.js,31f4+QIEeqL.js,01N6xzIJxbL.js,518BI433aLL.js,01rpauTep4L.js,31QZSjMuoeL.js,61ofwvddDeL.js,01KsMxlPtzL.js_.js?AUIClients/AmazonUI",
				"https://images-na.ssl-images-amazon.com/images/I/21G215oqvfL._RC|21OJDARBhQL.js,218GJg15I8L.js,31lucpmF4CL.js,2119M3Ks9rL.js,51AjK+mcqiL.js_.js?AUIClients/AuthenticationPortalAssets",
				"https://images-na.ssl-images-amazon.com/images/I/01wGDSlxwdL.js?AUIClients/AuthenticationPortalInlineAssets",
				"https://images-na.ssl-images-amazon.com/images/I/31TyXYXSSZL.js?AUIClients/CVFAssets",
				"https://images-na.ssl-images-amazon.com/images/I/81gLkT0N6tL.js?AUIClients/SiegeClientSideEncryptionAUI",
				"https://images-na.ssl-images-amazon.com/images/I/31jdfgcsPAL.js?AUIClients/AmazonUIFormControlsJS",
				"https://images-na.ssl-images-amazon.com/images/I/71QCGnDPPDL.js?AUIClients/FWCIMAssets",
				"https://images-na.ssl-images-amazon.com/images/I/71SLENeYBFL.js?AUIClients/ACICAssets",
				"https://static.siege-amazon.com/prod/profiles/AuthenticationPortalSigninNA.js",
			},
			InlineHashes: []int{
				-1746719145, -1110312912, -815800953, -314038750, -211036372, 216868775, 1424856663, 288533800, 318224283, -362024285, 585973559, 4606827, -1611905557, 1800521327, 2118020403, 1532181211, -1957835502, -222105570, -1616636591,
			},
			Elapsed:           8,
			DynamicURLCount:   11,
			InlineHashesCount: 19,
		},
		History: History{
			Length: randomInt(15, 5),
		},
		Battery: map[int]int{},
		Performance: Performance{
			Timing: Timing{
				ConnectStart:               base.Start - 1882,
				NavigationStart:            base.Start - 1885,
				LoadEventEnd:               base.Start - 414,
				DomLoading:                 base.Start - 864,
				SecureConnectionStart:      base.Start - 1800,
				FetchStart:                 base.Start - 1872,
				DomContentLoadedEventStart: base.Start - 548,
				ResponseStart:              base.Start - 892,
				ResponseEnd:                base.Start - 835,
				DomInteractive:             base.Start - 548,
				DomainLookupEnd:            base.Start - 1822,
				RedirectStart:              0,
				RequestStart:               base.Start - 1767,
				UnloadEventEnd:             base.Start - 880,
				UnloadEventStart:           base.Start - 884,
				DomComplete:                base.Start - 422,
				DomainLookupStart:          base.Start - 1882,
				LoadEventStart:             base.Start - 422,
				DomContentLoadedEventEnd:   base.Start - 548,
				RedirectEnd:                0,
				ConnectEnd:                 base.Start - 1767,
			},
		},
		Automation: Automation{
			Phantom: Phantom{
				PhantomProperties{
					Window: []interface{}{},
				},
			},
			Wd: Wd{
				WdProperties{
					Document:  []interface{}{},
					Navigator: []interface{}{},
					Window:    []interface{}{},
				},
			},
		},
		End:          randomInt(30000, 20000),
		TimeZone:     -8,
		FlashVersion: nil,
		Plugins:      "dOHjxYMm 4816958.fv379. 583583271Portable Document Format extension Brave document  ||1920-1080-984-24-*-*-*",
		DupedPlugins: "dOHjxYMm 4816958.fv379. 583583271Portable Document Format extension Brave document  ||1920-1080-984-24-*-*-*",
		ScreenInfo:   "1920-1080-984-24-*-*-*",
		Referrer:     base.Referrer,
		UserAgent:    base.UserAgent,
		Location:     base.Location,
		WebDriver:    false,
		Capabilities: Capabilities{
			CSS: CSS{
				TextShadow:       1,
				WebkitTextStroke: 1,
				BoxShadow:        1,
				BorderRadius:     1,
				BorderImage:      1,
				Opacity:          1,
				Transform:        1,
				Transition:       1,
			},
			Js: Js{
				Audio:        true,
				Geolocation:  true,
				LocalStorage: "supported",
				Touch:        false,
				Video:        true,
				WebWorker:    true,
			},
			Elapsed: 0,
		},
		Gpu: Gpu{
			Vendor: "Apple",
			Model:  "Apple M1",
			Extensions: []string{
				"ANGLE_instanced_arrays",
				"EXT_blend_minmax",
				"EXT_color_buffer_half_float",
				"EXT_disjoint_timer_query",
				"EXT_float_blend",
				"EXT_frag_depth",
				"EXT_shader_texture_lod",
				"EXT_texture_compression_rgtc",
				"EXT_texture_filter_anisotropic",
				"WEBKIT_EXT_texture_filter_anisotropic",
				"EXT_sRGB",
				"OES_element_index_uint",
				"OES_fbo_render_mipmap",
				"OES_standard_derivatives",
				"OES_texture_float",
				"OES_texture_float_linear",
				"OES_texture_half_float",
				"OES_texture_half_float_linear",
				"OES_vertex_array_object",
				"WEBGL_color_buffer_float",
				"WEBGL_compressed_texture_s3tc",
				"WEBKIT_WEBGL_compressed_texture_s3tc",
				"WEBGL_compressed_texture_s3tc_srgb",
				"WEBGL_debug_renderer_info",
				"WEBGL_debug_shaders",
				"WEBGL_depth_texture",
				"WEBKIT_WEBGL_depth_texture",
				"WEBGL_draw_buffers",
				"WEBGL_lose_context",
				"WEBKIT_WEBGL_lose_context",
				"WEBGL_multi_draw",
			},
		},
		Dnt: nil,
		Math: Math{
			Tan: "-1.4214488238747245",
			Sin: "0.8178819121159085",
			Cos: "-0.5753861119575491",
		},
		TimeToSubmit: randomInt(12500, 11000),
		Form: Form{
			ApCredentialAutofillHint: ApCredentialAutofillHint{
				Clicks:                0,
				Touches:               0,
				KeyPresses:            0,
				Cuts:                  0,
				Copies:                0,
				Pastes:                0,
				KeyPressTimeIntervals: []interface{}{},
				MouseClickPositions:   []interface{}{},
				KeyCycles:             []interface{}{},
				MouseCycles:           []interface{}{},
				TouchCycles:           []interface{}{},
				Width:                 0,
				Height:                0,
				TotalFocusTime:        0,
				Checksum:              "",
				Autocomplete:          false,
				Prefilled:             true,
			},
			Password: Password{
				Clicks:                0,
				Touches:               0,
				KeyPresses:            0,
				Cuts:                  0,
				Copies:                0,
				Pastes:                0,
				KeyPressTimeIntervals: []interface{}{},
				MouseClickPositions:   []interface{}{},
				KeyCycles:             []interface{}{},
				MouseCycles:           []interface{}{},
				TouchCycles:           []interface{}{},
				Width:                 296,
				Height:                31,
				TotalFocusTime:        0,
				Autocomplete:          false,
				Prefilled:             true,
			},
		},
		Canvas: Canvas{
			Hash:          0,
			EmailHash:     nil,
			HistogramBins: []interface{}{},
		},
		Token: Token{
			IsCompatible:   true,
			PageHasCaptcha: 0,
		},
		Auth: Auth{
			FormAuth{
				Method: "POST",
			},
		},
		Errors:  []interface{}{},
		Version: "4.0.0",
		LsUbid:  "X59-9642719-1673432:" + strconv.Itoa(nowDate()/1e3),
	}
}

func (p *Payload) ToString() string {
	j, _ := json.Marshal(p)
	return string(j)
}
