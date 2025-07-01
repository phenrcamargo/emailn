package enum

type CampaignStatus int

const (
	Pending    CampaignStatus = 0
	InProgress CampaignStatus = 1
	Done       CampaignStatus = 2
)

func (c CampaignStatus) GetDescribe() string {
	switch c {
	case Pending:
		return "Pending"

	case InProgress:
		return "In Progress"

	case Done:
		return "Done"

	default:
		return "Unknown campaign status"
	}
}
