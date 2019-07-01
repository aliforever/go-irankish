package irankish

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/aliforever/go-irankish/file"

	"github.com/aliforever/go-irankish/translation"

	"github.com/go-errors/errors"
)

type IranKish struct {
	MerchantId string
	Sha1Key    string
	Payment    *Payment
	Verify     *VerifyPayment
}

type MakeTokenResult struct {
	Message string
	Result  bool
	Token   string
}

type VerifyPaymentResult struct {
	Result string
}

func (ik *IranKish) makeTokenXML() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://tempuri.org/">
    <SOAP-ENV:Body>
        <ns1:MakeToken>
            %tags%
        </ns1:MakeToken>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
}

func (ik *IranKish) verifyPaymentXML() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://tempuri.org/">
    <SOAP-ENV:Body>
        <ns1:KicccPaymentsVerification>
            %tags%
        </ns1:KicccPaymentsVerification>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
}

func (ik *IranKish) MakeToken() (mtr *MakeTokenResult, err error) {
	if ik.MerchantId == "" {
		err = errors.New("empty_merchant_id")
		return
	}
	if ik.Payment == nil {
		err = errors.New("nil_payment")
		return
	}
	ik.Payment.merchantId = ik.MerchantId
	tags, err := ik.Payment.getTags()
	if err != nil {
		return
	}
	joinTags := strings.Join(tags, "\n")

	stringXML := ik.makeTokenXML()
	finalXML := strings.Replace(stringXML, "%tags%", joinTags, -1)
	client := http.Client{}
	request, err := http.NewRequest("POST", makeTokenUrl, strings.NewReader(finalXML))
	request.Header.Add("Host", "ikc.shaparak.ir")
	request.Header.Add("Connection", "Keep-Alive")
	request.Header.Add("User-Agent", "PHP-SOAP/5.6.30")
	request.Header.Add("Content-Type", "text/xml; charset=utf-8")
	request.Header.Add("SOAPAction", `"http://tempuri.org/ITokens/MakeToken"`)

	req, err := client.Do(request)
	if err != nil {
		return
	}
	defer req.Body.Close()
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	mtr, err = ik.parseMakeTokenResult(string(res))
	if err != nil {
		return
	}
	return
}

func (ik *IranKish) parseMakeTokenResult(response string) (mtr *MakeTokenResult, err error) {
	if !strings.Contains(response, "<MakeTokenResult ") {
		file.PutContents("make_token_response.html", []byte(response))
		err = errors.New("wrong_response")
		return
	}
	fmt.Println(response)
	mtr = &MakeTokenResult{}
	messageRegex := regexp.MustCompile(`<a:message.+"/>(.+)</a:message>`)
	message := messageRegex.FindStringSubmatch(response)
	if len(message) > 0 {
		mtr.Message = strings.TrimSpace(message[0])
	}
	resultRegex := regexp.MustCompile(`<a:result>(.+)</a:result>`)
	result := resultRegex.FindStringSubmatch(response)
	if len(result) > 0 {
		if result[0] == "true" {
			mtr.Result = true
		} else {
			mtr.Result = false
		}
	}
	tokenRegex := regexp.MustCompile(`<a:token>(.+)</a:token>`)
	token := tokenRegex.FindStringSubmatch(response)
	if len(token) > 0 {
		mtr.Token = token[0]
	}
	//file.PutContents("make_token_response.html", []byte(response))
	return
}

func (ik *IranKish) VerifyPayment() (vpr *VerifyPaymentResult, err error) {
	if ik.MerchantId == "" {
		err = errors.New("empty_merchant_id")
		return
	}
	if ik.Sha1Key == "" {
		err = errors.New("empty_sha1key")
		return
	}
	if ik.Verify == nil {
		err = errors.New("nil_payment")
		return
	}

	ik.Verify.merchantId = ik.MerchantId
	ik.Verify.sha1Key = ik.Sha1Key
	tags, err := ik.Verify.getTags()
	if err != nil {
		return
	}
	joinTags := strings.Join(tags, "\n")
	/*gopath := os.Getenv("GOPATH")
	verifyPaymentXML, err := file.GetContents(gopath + "/src/github.com/aliforever/go-irankish/xml/verifyPayment.xml")
	if err != nil {
		return
	}
	stringXML := string(verifyPaymentXML)*/
	stringXML := ik.verifyPaymentXML()
	finalXML := strings.Replace(stringXML, "%tags%", joinTags, -1)
	client := http.Client{}
	request, err := http.NewRequest("POST", verifyPaymentUrl, strings.NewReader(finalXML))
	request.Header.Add("Host", "ikc.shaparak.ir")
	request.Header.Add("Connection", "Keep-Alive")
	request.Header.Add("User-Agent", "PHP-SOAP/5.6.30")
	request.Header.Add("Content-Type", "text/xml; charset=utf-8")
	request.Header.Add("SOAPAction", `"http://tempuri.org/IVerify/KicccPaymentsVerification"`)

	req, err := client.Do(request)
	if err != nil {
		return
	}
	defer req.Body.Close()
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	vpr, err = ik.parseVerifyPaymentResult(string(res))
	if err != nil {
		return
	}
	return
}

func (ik *IranKish) parseVerifyPaymentResult(response string) (vpr *VerifyPaymentResult, err error) {
	if !strings.Contains(response, "<KicccPaymentsVerificationResult>") {
		file.PutContents("verify_payment_response.html", []byte(response))
		err = errors.New("wrong_response")
		return
	}
	vpr = &VerifyPaymentResult{}
	verificationResultTag := []string{"<KicccPaymentsVerificationResult>", "</KicccPaymentsVerificationResult>"}
	verificationResultStartTagIndex := strings.Index(response, verificationResultTag[0])
	verificationResult := response[verificationResultStartTagIndex+len(verificationResultTag[0]) : strings.Index(response, verificationResultTag[1])]
	vpr.Result = verificationResult
	return
}

func (ik *IranKish) ResultMessage(code string) string {
	if _, ok := translation.CallBackCodes[code]; !ok {
		return "unknown_error"
	}
	return translation.CallBackCodes[code]
}
