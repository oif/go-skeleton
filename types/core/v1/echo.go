package v1

type Echo struct {
	Name string
}

func (e Echo) Hey() string {
	return "Hey " + e.Name
}