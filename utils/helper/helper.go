package helper

import (
	"csat-report-webhook/entities"
	"csat-report-webhook/viewmodels"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func MakeMetabaseRequest(request viewmodels.JSONRequest, multichannel entities.Multichannel) (*viewmodels.MetabaseDataRequest, error) {
	parameters := []viewmodels.MetabaseParameter{
		viewmodels.MetabaseParameter{
			Type: "category",
			Target: []interface{}{
				"variable",
				[]string{"template-tag", "app_code"},
			},
			Value: multichannel.GetAppID(),
		},
		viewmodels.MetabaseParameter{
			Type: "category",
			Target: []interface{}{
				"variable",
				[]string{"template-tag", "row_limit"},
			},
			Value: strconv.Itoa(request.Limit),
		},
	}

	if len(request.On) > 0 {
		target := []interface{}{
			"dimension",
			[]string{"template-tag", "created_at"},
		}
		parameters = append(parameters, viewmodels.MetabaseParameter{
			Type:   "date/all-options",
			Target: target,
			Value:  request.On,
		})
	}

	if len(request.StartDate) > 0 && len(request.EndDate) > 0 {
		layout := "2006-01-02"
		parsedStartDate, err := time.Parse(layout, request.StartDate)
		if err != nil {
			return &viewmodels.MetabaseDataRequest{}, err
		}

		parsedEndDate, err := time.Parse(layout, request.EndDate)
		if err != nil {
			return &viewmodels.MetabaseDataRequest{}, err
		}
		target := []interface{}{
			"dimension",
			[]string{"template-tag", "created_at"},
		}
		if !parsedEndDate.After(parsedStartDate) {
			return &viewmodels.MetabaseDataRequest{}, errors.New("end_date must greater than start_date")
		}
		parameters = append(parameters, viewmodels.MetabaseParameter{
			Type:   "date/all-options",
			Target: target,
			Value:  fmt.Sprintf("%s~%s", request.StartDate, request.EndDate),
		})
	}

	return &viewmodels.MetabaseDataRequest{
		IgnoreCache: true,
		Parameters:  parameters,
	}, nil
}

func MakeMetabaseSerializedParams(request viewmodels.JSONRequest, multichannel entities.Multichannel) (string, error) {
	parameters := []viewmodels.MetabaseParameter{
		viewmodels.MetabaseParameter{
			Type: "category",
			Target: []interface{}{
				"variable",
				[]string{"template-tag", "app_code"},
			},
			Value: multichannel.GetAppID(),
		},
		viewmodels.MetabaseParameter{
			Type: "category",
			Target: []interface{}{
				"variable",
				[]string{"template-tag", "row_limit"},
			},
			Value: strconv.Itoa(request.Limit),
		},
	}

	if len(request.On) > 0 {
		target := []interface{}{
			"dimension",
			[]string{"template-tag", "created_at"},
		}
		parameters = append(parameters, viewmodels.MetabaseParameter{
			Type:   "date/all-options",
			Target: target,
			Value:  request.On,
		})
	}

	if len(request.StartDate) > 0 && len(request.EndDate) > 0 {
		layout := "2006-01-02"
		parsedStartDate, err := time.Parse(layout, request.StartDate)
		if err != nil {
			return "", err
		}

		parsedEndDate, err := time.Parse(layout, request.EndDate)
		if err != nil {
			return "", err
		}
		target := []interface{}{
			"dimension",
			[]string{"template-tag", "created_at"},
		}
		if !parsedEndDate.After(parsedStartDate) {
			return "", errors.New("end_date must greater than start_date")
		}
		parameters = append(parameters, viewmodels.MetabaseParameter{
			Type:   "date/all-options",
			Target: target,
			Value:  fmt.Sprintf("%s~%s", request.StartDate, request.EndDate),
		})
	}

	jsonParams, _ := json.Marshal(parameters)

	values := url.Values{"parameters": {string(jsonParams[:])}}

	return values.Encode(), nil
}
