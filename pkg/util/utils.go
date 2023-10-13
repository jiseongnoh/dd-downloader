package util

import "github.com/girishg4t/dd-downloader/pkg/model"
func ReadHeader(m []model.InnerFieldMapping, headers *[]string) {
	// add the host header in CSV file 
	*headers = append(*headers, "host")

	for _, val := range m {
		if val.Field == "-" && val.InnerField != nil || len(val.InnerField) > 0 {
			ReadHeader(val.InnerField, headers)
			continue
		}
		*headers = append(*headers, val.Field)
	}
}
