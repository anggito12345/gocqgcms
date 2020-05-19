package models

/*CustomerModel customer properties
 */
type CustomerModel struct {
	Name        string `json:"name" bson:"name"`
	LegalType   string `json:"legaltype" bson:"legaltype"`
	BrokerageID string `json:"brokerageid" bson:"brokerageid"`
}

/*NewCustomer intialize customer model
 */
func NewCustomer() *CustomerModel {
	return &CustomerModel{}
}
