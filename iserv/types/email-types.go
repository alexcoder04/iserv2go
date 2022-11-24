package types

type EMail struct {
	Subject    string
	From       string
	To         string
	ToDispName string
	CCs        []string
	Body       string
}
