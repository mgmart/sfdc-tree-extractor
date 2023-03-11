package main

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

func pseudomyse(obj *sObject) {

	// log.Debug("Obj before: ", obj)

	switch obj.Type {
	case "Account":
		for key := range obj.Body {
			switch key {
			case "Name":
				obj.Body["Name"] = gofakeit.Company()
			case "Phone":
				obj.Body["Phone"] = gofakeit.PhoneFormatted()
			case "BillingCity":
				obj.Body["BillingCity"] = gofakeit.City()
			case "BillingState":
				obj.Body["BillingState"] = gofakeit.State()
			case "BillingStreet":
				obj.Body["BillingStreet"] = gofakeit.Street()
			case "Fax":
				obj.Body["Fax"] = gofakeit.PhoneFormatted()
			case "Website":
				obj.Body["Website"] = gofakeit.URL()
			case "AccountNumber":
				obj.Body["AccountNumber"] = strconv.Itoa(gofakeit.Number(1111111, 9999999))
			}
		}
	case "Contact":
		for key := range obj.Body {
			switch key {
			case "FirstName":
				obj.Body["FirstName"] = gofakeit.FirstName()
			case "Phone":
				obj.Body["Phone"] = gofakeit.PhoneFormatted()
			case "LastName":
				obj.Body["LastName"] = gofakeit.LastName()
			case "Fax":
				obj.Body["Fax"] = gofakeit.PhoneFormatted()
			case "Email":
				obj.Body["Email"] = gofakeit.Email()

			}
		}
	case "Campaign":
		for key := range obj.Body {
			switch key {
			case "Name":
				obj.Body["Name"] = gofakeit.SentenceSimple()
			}
		}

	case "Lead":
		for key := range obj.Body {
			switch key {
			case "Company":
				obj.Body["Company"] = gofakeit.Company()
			case "FirstName":
				obj.Body["FirstName"] = gofakeit.FirstName()
			case "City":
				obj.Body["City"] = gofakeit.City()
			case "LastName":
				obj.Body["LastName"] = gofakeit.LastName()
			case "State":
				obj.Body["State"] = gofakeit.StateAbr()
			case "Email":
				obj.Body["Email"] = gofakeit.Email()
			case "Phone":
				obj.Body["Phone"] = gofakeit.PhoneFormatted()
			case "Country":
				obj.Body["Country"] = gofakeit.Country()
			case "Street":
				obj.Body["Street"] = gofakeit.Street()
			case "Fax":
				obj.Body["Fax"] = gofakeit.PhoneFormatted()

			}
		}
	}
	// log.Debug("Obj after: ", obj)

}
