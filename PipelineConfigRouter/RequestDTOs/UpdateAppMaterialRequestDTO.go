package RequestDTOs

type UpdateAppMaterialRequestDTO struct {
	AppId    int             `json:"appId"`
	Material UpdatedMaterial `json:"material"`
}

type UpdatedMaterial struct {
	Url             string `json:"url"`
	Id              int    `json:"id"`
	GitProviderId   int    `json:"gitProviderId"`
	CheckoutPath    string `json:"checkoutPath"`
	FetchSubmodules bool   `json:"fetchSubmodules"`
}
