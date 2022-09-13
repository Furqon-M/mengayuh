package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	camapaignFormatter := CampaignFormatter{}
	camapaignFormatter.ID = campaign.ID
	camapaignFormatter.UserID = campaign.UserID
	camapaignFormatter.Name = campaign.Name
	camapaignFormatter.ShortDescription = campaign.ShortDescription
	camapaignFormatter.GoalAmount = campaign.GoalAmount
	camapaignFormatter.CurrentAmount = campaign.CurrentAmount
	camapaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		camapaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return camapaignFormatter

}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)

	}

	return campaignsFormatter
}
