package contract

type NewCampaign struct {
	Name     string
	Content  string
	Contacts []NewContact
}
