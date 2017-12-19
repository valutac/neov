package authentication

type domain struct {
	Name string `json:"name"`
}

type identity struct {
	Methods  []string `json:"methods"`
	Password struct {
		User struct {
			Domain   *domain `json:"domain,omitempty"`
			Name     string  `json:"name"`
			Password string  `json:"password"`
		} `json:"user"`
	} `json:"password"`
}

type project struct {
	Domain *domain `json:"domain,omitempty"`
	ID     string  `json:"id"`
	Name   string  `json:"name"`
}

type scope struct {
	Project project `json:"project"`
}

type data struct {
	Auth struct {
		Identity identity `json:"identity"`
		Scope    *scope   `json:"scope,omitempty"`
	} `json:"auth"`
}
