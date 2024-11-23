package github

type Tools struct {
	Rust   string
	Go     string
	TinyGo string
	Node   string
	Python string
	Spin   string
}

func DefaultTools() Tools {
	return Tools{
		Rust:   "1.80.1",
		Go:     "1.23.2",
		TinyGo: "v0.33.0",
		Python: "3.13.0",
		Node:   "22",
		Spin:   "",
	}
}
