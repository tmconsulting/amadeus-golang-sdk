package queue_placepnr

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/formats"
)

type QueuePlacePNR struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QUQPCQ_03_1_1A Queue_PlacePNR"`

	// specify the type of placement to be done queue place on queue  queue place on another queue  queue place off queue
	PlacementOption *SelectionDetailsTypeI `xml:"placementOption"`

	TargetDetails []*TargetDetails `xml:"targetDetails"`  // maxOccurs="10"

	// contains the record locator to be queue placed
	RecordLocator *ReservationControlInformationTypeI `xml:"recordLocator"`
}

type TargetDetails struct {

	// used to specify the target office for which the queue count is to be displayed
	TargetOffice *AdditionalBusinessSourceInformationType `xml:"targetOffice"`

	// used to specify the queue if required
	QueueNumber *QueueInformationTypeI `xml:"queueNumber,omitempty"`  // minOccurs="0"

	// used to select the category
	CategoryDetails *SubQueueInformationTypeI `xml:"categoryDetails,omitempty"`  // minOccurs="0"

	// used to pass the target date/time if not default
	PlacementDate *StructuredDateTimeInformationType `xml:"placementDate,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type AdditionalBusinessSourceInformationType struct {

	// SOURCE TYPE
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	// ORIGINATOR DETAILS
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails,omitempty"`  // minOccurs="0"
}

type OriginatorIdentificationDetailsTypeI struct {

	// the office that is being targetted
	InHouseIdentification1 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification1"`
}

type QueueInformationDetailsTypeI struct {

	// queue number
	Number formats.NumericInteger_Length1To2 `xml:"number"`
}

type QueueInformationTypeI struct {

	// queue identification
	QueueDetails *QueueInformationDetailsTypeI `xml:"queueDetails"`
}

type ReservationControlInformationDetailsTypeI struct {

	// contains the record locator to be queue placed
	ControlNumber formats.AlphaNumericString_Length1To8 `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {

	// contains the record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation"`
}

type SelectionDetailsInformationTypeI struct {

	// to specify if noramal or delay queue
	Option formats.AlphaNumericString_Length1To3 `xml:"option"`
}

type SelectionDetailsTypeI struct {

	// specify queue type
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails"`
}

type SourceTypeDetailsTypeI struct {

	// to define if own office or different office being targetted
	SourceQualifier1 formats.AlphaNumericString_Length1To3 `xml:"sourceQualifier1"`
}

type StructuredDateTimeInformationType struct {

	// used to select the date range
	TimeMode *formats.NumericInteger_Length1To1 `xml:"timeMode,omitempty"`  // minOccurs="0"

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year formats.Year_YYYY `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day"`

	// Hour between 0 and 23
	Hour formats.Hour_hH `xml:"hour,omitempty"`  // minOccurs="0"
}

type SubQueueInformationDetailsTypeI struct {

	// E for every category    A for cats with items to be worked C for category number N for nickname CN for both category number and nickname numeric for date range
	IdentificationType formats.AlphaNumericString_Length1To3 `xml:"identificationType"`

	// category number
	ItemNumber formats.AlphaNumericString_Length1To3 `xml:"itemNumber,omitempty"`  // minOccurs="0"

	// used for nickname on inbound used for category name on outbound
	ItemDescription formats.AlphaNumericString_Length1To35 `xml:"itemDescription,omitempty"`  // minOccurs="0"
}

type SubQueueInformationTypeI struct {

	// identifies the category or categories.
	SubQueueInfoDetails *SubQueueInformationDetailsTypeI `xml:"subQueueInfoDetails"`
}
