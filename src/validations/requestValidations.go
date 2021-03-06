package validations

import "github.com/laithrafid/bookstore_utils-go/errors_utils"

func RequestQParams(qp map[string][]string, rType string) (bool, error) {
	if rType == "source" {
		if qp["country"] != nil && qp["language"] != nil && qp["category"] != nil {
			// return false,
			e := errors_utils.NewBadRequestError("Invalid params. Either country and language or category")
			return false, e
		}
	} else if rType == "all" {
		if qp["top"] == nil || qp["sources"] == nil {
			// return false,
			e := errors_utils.NewBadRequestError("Invalid params. Either top or language is missing")
			return false, e
		}
	}
	return true, nil
}
