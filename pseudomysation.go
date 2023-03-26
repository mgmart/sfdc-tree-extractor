package sfdcTreeExtractor

import (
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/cloudflare/cfssl/log"
	"github.com/go-faker/faker/v4"
)

func newPseudo() {
	var r string
	for _, v := range Config.Pseudo {
		for k, v := range v {
			switch v {
			case "PhoneNumber":
				r = faker.Phonenumber()
			case "Company":
				r = fakeCompany()
			case "FirstName":
				r = faker.FirstName()
			case "LastName":
				r = faker.LastName()
			}
			log.Debug("KV: ", k, " - ", r)
		}
	}
	log.Debug("Pseudo: ", Config.Pseudo["Account"]["Name"])

	log.Debug("Address", faker.GetRealAddress())

	//	log.Debug(f.Call(nil))
}

func pseudomyse(obj *sObject) {

	switch obj.Type {
	case "Account":
		for key := range obj.Body {
			switch key {
			case "Name":
				obj.Body["Name"] = fakeCompany()
			case "Phone":
				obj.Body["Phone"] = faker.PhoneNumber
			case "BillingCity":
				obj.Body["BillingCity"] = faker.GetRealAddress().City
			case "BillingState":
				obj.Body["BillingState"] = faker.GetRealAddress().State
			case "BillingStreet":
				obj.Body["BillingStreet"] = faker.GetRealAddress().Address
			case "Fax":
				obj.Body["Fax"] = faker.PhoneNumber
			case "Website":
				obj.Body["Website"] = faker.URL()
			case "AccountNumber":
				obj.Body["AccountNumber"] = faker.CCNumber()
			}
		}
	case "Contact":
		for key := range obj.Body {
			switch key {
			case "FirstName":
				obj.Body["FirstName"] = faker.FirstName()
			case "Phone":
				obj.Body["Phone"] = faker.Phonenumber()
			case "LastName":
				obj.Body["LastName"] = faker.LastName()
			case "Fax":
				obj.Body["Fax"] = faker.Phonenumber()
			case "MobilePhone":
				obj.Body["MobilePhone"] = faker.Phonenumber()
			case "Email":
				obj.Body["Email"] = faker.Email()
			case "MailingStreet":
				addr := faker.GetRealAddress()
				obj.Body["MailingStreet"] = addr.Address + "\n" + addr.City + ", " + addr.State + " " + addr.PostalCode + "\nUSA"
			}
		}
	case "Campaign":
		for key := range obj.Body {
			switch key {
			case "Name":
				obj.Body["Name"] = faker.Sentence()
			}
		}

	case "Lead":
		for key := range obj.Body {
			switch key {
			case "Company":
				obj.Body["Company"] = faker.Name()
			case "FirstName":
				obj.Body["FirstName"] = faker.FirstName()
			case "City":
				obj.Body["City"] = faker.GetRealAddress().City
			case "LastName":
				obj.Body["LastName"] = faker.LastName()
			case "State":
				obj.Body["State"] = faker.GetRealAddress().State
			case "Email":
				obj.Body["Email"] = faker.Email()
			case "Phone":
				obj.Body["Phone"] = faker.PhoneNumber
			case "Country":
				obj.Body["Country"] = faker.GetRealAddress().State
			case "Street":
				obj.Body["Street"] = faker.GetRealAddress().Address
			case "Fax":
				obj.Body["Fax"] = faker.PhoneNumber

			}
		}
	}
	// log.Debug("Obj after: ", obj)

}

type Person struct {
	FirstName string `faker:"firstName"`
	LastName  string `faker:"lastName"`
}

func (pers Person) eMailPart() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	ext := emailExtension[r1.Intn(len(emailExtension)-1)]

	return pers.FirstName + "." + pers.LastName + "@" + ext
}

type Company struct {
	Name string `faker:"company"`
}

// CustomGenerator ...
func myCustomGenerator() {
	_ = faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
		return int64(43), nil
	})
	_ = faker.AddProvider("firstName", func(v reflect.Value) (interface{}, error) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		return strings.Title(women[r1.Intn(50)]), nil
	})

	_ = faker.AddProvider("lastName", func(v reflect.Value) (interface{}, error) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		return strings.Title(lastName[r1.Intn(200)]), nil
	})

	_ = faker.AddProvider("company", func(v reflect.Value) (interface{}, error) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		ret := strings.Title(adjectives[r1.Intn(len(adjectives)-1)]) + "e " + substantives[r1.Intn(len(substantives)-1)] + " " + companyForm[r1.Intn(len(companyForm)-1)]
		return ret, nil
	})

}

// You can also add your own generator function to your own defined tags.
func fakeCompany() string {
	log.Level = log.LevelDebug
	myCustomGenerator()
	var company Company
	_ = faker.FakeData(&company)
	log.Debug("fc: ", company)
	return company.Name
}
