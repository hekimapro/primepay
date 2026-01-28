package primepay

import "encoding/json"

func Collect(payload *CollectionRequestPayload) (*CollectionResponsePayload, error) {

	path := "/api/v1/transact"
	payload.Currency = "TZS"
	payload.Action = "collection"
	responseByte, err := post(path, payload)
	if err != nil {
		return nil, err
	}

	var responseBody CollectionResponsePayload
	if err := json.Unmarshal(responseByte, &responseBody); err != nil {
		return nil, err
	}

	return &responseBody, nil

}

func Disburse(payload *DisbursementRequestPayload) (*DisbursementResponsePayload, error) {

	path := "/api/v1/transact"
	payload.Action = "disbursement"

	if payload.Channel == "" {
		payload.Channel = "all"
	}

	responseByte, err := post(path, payload)
	if err != nil {
		return nil, err
	}

	var responseBody DisbursementResponsePayload
	if err := json.Unmarshal(responseByte, &responseBody); err != nil {
		return nil, err
	}

	return &responseBody, nil

}

func CheckStatus(reference string) (*StatusResponsePayload, error) {

	payload := &statusRequestPayload{Reference: reference}

	path := "/api/v1/query"
	responseByte, err := post(path, payload)
	if err != nil {
		return nil, err
	}

	var responseBody StatusResponsePayload
	if err := json.Unmarshal(responseByte, &responseBody); err != nil {
		return nil, err
	}

	return &responseBody, nil

}
