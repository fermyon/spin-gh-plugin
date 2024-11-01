package github

type ActionTriggers struct {
	ManualDispatch bool
	Schedule       string
	PullRequest    string
	Push           string
}
