package main

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

func pseudomyse(obj *sObject) {

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

	}
}
