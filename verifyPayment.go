package irankish

import (
	"errors"
	"fmt"
)

type VerifyPayment struct {
	merchantId      string
	sha1Key         string
	Token           string
	ReferenceNumber string
}

func (vp *VerifyPayment) getTags() (tagsArray []string, err error) {
	tags := map[string]string{
		"token":           "<ns1:token>%s</ns1:token>",
		"merchantId":      "<ns1:merchantId>%s</ns1:merchantId>",
		"referenceNumber": "<ns1:referenceNumber>%s</ns1:referenceNumber>",
		"sha1Key":         "<ns1:sha1Key>%s</ns1:sha1Key>",
	}
	if vp.Token == "" {
		err = errors.New("empty_token")
		return
	}
	if vp.merchantId == "" {
		err = errors.New("empty_merchant_id")
		return
	}
	if vp.ReferenceNumber == "" {
		err = errors.New("empty_reference_id")
		return
	}
	if vp.sha1Key == "" {
		err = errors.New("empty_sha1_key")
		return
	}
	tagsArray = append(tagsArray, fmt.Sprintf(tags["token"], vp.Token))
	tagsArray = append(tagsArray, fmt.Sprintf(tags["merchantId"], vp.merchantId))
	tagsArray = append(tagsArray, fmt.Sprintf(tags["referenceNumber"], vp.ReferenceNumber))
	tagsArray = append(tagsArray, fmt.Sprintf(tags["sha1Key"], vp.sha1Key))

	return
}
