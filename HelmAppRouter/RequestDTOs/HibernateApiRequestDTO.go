package RequestDTOs

type HibernateApiRequestDTO struct {
	AppId     string     `json:"appId"`
	Resources []Resource `json:"resources"`
}
type Resource struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Version   string `json:"version"`
	Namespace string `json:"namespace"`
}
