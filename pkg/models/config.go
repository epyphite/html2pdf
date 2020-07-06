package models

//Config is the basic configuration structure for the service
type Config struct {
	DestinationFolder string `json:"destination_folder"`
	SourceFolder      string `json:"source_folder"`
	SourceType        string `json:"source_type"`
	SourceName        string `json:"source_name"`
	TabName           string `json:"tab_name"`
	ColumnNumber      string `json:"column_number"`
	EnableCloud       bool   `json:"enablecloud"`
	CloudProvider     string `json:"cloudprovider"`
	CloudRegion       string `json:"cloudregion"`
	CloudStorage      string `json:"cloudstorage"`     //Bucket or Root Key folder
	CloudDestination  string `json:"clouddestination"` //Folder or Path

}
