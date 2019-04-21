package sdk

import (
	"github.com/tmconsulting/amadeus-golang-sdk/logger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/nilLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/session/v03.0"
)

var (
	soapUrl = "http://webservices.amadeus.com/WSAP/"
	amaUrl  = "http://webservices.amadeus.com/"
)

func CreateAmadeusClient(url, originator, passwordRaw, officeId string, lw logger.LogWriter) *AmadeusClient {
	if lw == nil {
		lw = nilLogger.Init()
	}
	return &AmadeusClient{
		Session: CreateSession(),
		service: &WebServicePT{
			Client: &SOAP4Client{
				Url:       url,
				User:      originator,
				Pass:      passwordRaw,
				Agent:     officeId,
				TLSEnable: true,
				Logger:    logger.NewLogger(lw),
			},
		},
	}
}

type AmadeusClient struct {
	//service *soap4_0.WebServicesPTSOAP4Header
	//Session *soap4_0.Session_v3
	// messageIds		[]string
	Session *Session_v03_0.Session
	service *WebServicePT
}

func (client *AmadeusClient) Close() error {
	if client == nil || client.service == nil {
		return nil
	}
	if client.Session != nil && client.Session.TransactionStatusCode != Session_v03_0.TransactionStatusCode[Session_v03_0.End] {
		client.Session.TransactionStatusCode = Session_v03_0.TransactionStatusCode[Session_v03_0.End]
		if _, _, err := client.SecuritySignOutV041(); err != nil {
			return err
		}
		client.Session = nil
	}
	return nil
}
