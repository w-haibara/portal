package portal

type listItem struct {
	title, desc string
}

func (i listItem) Title() string {
	return i.title
}

func (i listItem) Description() string {
	return i.desc
}

func (i listItem) FilterValue() string {
	return i.title
}
