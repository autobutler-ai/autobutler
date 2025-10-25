package types

type PageName string

const (
	PageCalendar PageName = "Calendar"
	PageFiles    PageName = "Files"
	PageHome     PageName = "Home"
	PagePhotos   PageName = "Photos"
)

type Page struct {
	Name PageName
	Href string
}

func newPage(name PageName, href string) Page {
	return Page{
		Name: name,
		Href: href,
	}
}
