package irankish

import (
	"errors"
	"fmt"
)

type Payment struct {
	Amount           string
	merchantId       string
	InvoiceId        string
	PaymentId        string
	SpecialPaymentId string
	CallbackUrl      string
	Description      string
	ExtraParameter1  string
	ExtraParameter2  string
	ExtraParameter3  string
	ExtraParameter4  string
}

func (p *Payment) getTags() (tagsArray []string, err error) {
	tags := map[string]string{
		"amount":           "<ns1:amount>%s</ns1:amount>",
		"merchantId":       "<ns1:merchantId>%s</ns1:merchantId>",
		"invoiceId":        "<ns1:invoiceNo>%s</ns1:invoiceNo>",
		"revertUrl":        "<ns1:revertURL>%s</ns1:revertURL>",
		"paymentId":        "<ns1:paymentId>%s</ns1:paymentId>",
		"specialPaymentId": "<ns1:specialPaymentId>%s</ns1:specialPaymentId>",
		"description":      "<ns1:description>%s</ns1:description>",
	}
	if p.Amount == "" {
		err = errors.New("empty_amount")
		return
	}
	if p.merchantId == "" || p.InvoiceId == "" || p.CallbackUrl == "" {
		err = errors.New("empty_merchant_id")
		return
	}
	if p.InvoiceId == "" || p.CallbackUrl == "" {
		err = errors.New("empty_invoice_id")
		return
	}
	if p.CallbackUrl == "" {
		err = errors.New("empty_callback_url")
		return
	}
	tagsArray = append(tagsArray, fmt.Sprintf(tags["amount"], p.Amount))
	tagsArray = append(tagsArray, fmt.Sprintf(tags["merchantId"], p.merchantId))
	tagsArray = append(tagsArray, fmt.Sprintf(tags["invoiceId"], p.InvoiceId))
	tagsArray = append(tagsArray, fmt.Sprintf(tags["revertUrl"], p.CallbackUrl))
	if p.PaymentId != "" {
		tagsArray = append(tagsArray, fmt.Sprintf(tags["paymentId"], p.PaymentId))
	}
	if p.SpecialPaymentId != "" {
		tagsArray = append(tagsArray, fmt.Sprintf(tags["specialPaymentId"], p.SpecialPaymentId))
	}
	if p.Description != "" {
		tagsArray = append(tagsArray, fmt.Sprintf(tags["description"], p.Description))
	}
	return
}
