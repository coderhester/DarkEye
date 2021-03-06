package fofa

type ipNode struct {
	Ip     string
	Domain string
	Title  string
	Server string
	Finger string
	Port   string
	Alive  bool
}

type Fofa struct {
	Interval    int    `json:"interval"`
	Ip          string `json:"ip"`
	FofaSession string `json:"fofa_session"`
	Pages       int    `json:"pages"`
	FofaComma   bool   `json:"fofa_comma"`

	ipNodes    []ipNode
	ErrChannel chan string `json:"-"`
	Stop       int32       `json:"-"`
}

func NewConfig() Fofa {
	return Fofa{
		Interval: 10,
		Pages:    5,
	}
}
