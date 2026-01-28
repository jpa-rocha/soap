//nolint:staticcheck
package countryInfo

import (
	"encoding/xml"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// against "unused imports"
var (
	_ time.Time
	_ xml.Name
)

type ListOfContinentsByName struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfContinentsByName"`
}

type ListOfContinentsByNameResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfContinentsByNameResponse"`

	ListOfContinentsByNameResult *ArrayOftContinent `xml:"ListOfContinentsByNameResult,omitempty"`
}

type ListOfContinentsByCode struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfContinentsByCode"`
}

type ListOfContinentsByCodeResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfContinentsByCodeResponse"`

	ListOfContinentsByCodeResult *ArrayOftContinent `xml:"ListOfContinentsByCodeResult,omitempty"`
}

type ListOfCurrenciesByName struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCurrenciesByName"`
}

type ListOfCurrenciesByNameResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCurrenciesByNameResponse"`

	ListOfCurrenciesByNameResult *ArrayOftCurrency `xml:"ListOfCurrenciesByNameResult,omitempty"`
}

type ListOfCurrenciesByCode struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCurrenciesByCode"`
}

type ListOfCurrenciesByCodeResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCurrenciesByCodeResponse"`

	ListOfCurrenciesByCodeResult *ArrayOftCurrency `xml:"ListOfCurrenciesByCodeResult,omitempty"`
}

type CurrencyName struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CurrencyName"`

	SCurrencyISOCode string `xml:"sCurrencyISOCode,omitempty"`
}

type CurrencyNameResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CurrencyNameResponse"`

	CurrencyNameResult string `xml:"CurrencyNameResult,omitempty"`
}

type ListOfCountryNamesByCode struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCountryNamesByCode"`
}

type ListOfCountryNamesByCodeResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCountryNamesByCodeResponse"`

	ListOfCountryNamesByCodeResult *ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByCodeResult,omitempty"`
}

type ListOfCountryNamesByName struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCountryNamesByName"`
}

type ListOfCountryNamesByNameResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCountryNamesByNameResponse"`

	ListOfCountryNamesByNameResult *ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByNameResult,omitempty"`
}

type ListOfCountryNamesGroupedByContinent struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCountryNamesGroupedByContinent"`
}

type ListOfCountryNamesGroupedByContinentResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfCountryNamesGroupedByContinentResponse"`

	ListOfCountryNamesGroupedByContinentResult *ArrayOftCountryCodeAndNameGroupedByContinent `xml:"ListOfCountryNamesGroupedByContinentResult,omitempty"`
}

type CountryName struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryName"`

	SCountryISOCode string `xml:"sCountryISOCode,omitempty"`
}

type CountryNameResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryNameResponse"`

	CountryNameResult string `xml:"CountryNameResult,omitempty"`
}

type CountryISOCode struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryISOCode"`

	SCountryName string `xml:"sCountryName,omitempty"`
}

type CountryISOCodeResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryISOCodeResponse"`

	CountryISOCodeResult string `xml:"CountryISOCodeResult,omitempty"`
}

type CapitalCity struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CapitalCity"`

	SCountryISOCode string `xml:"sCountryISOCode,omitempty"`
}

type CapitalCityResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CapitalCityResponse"`

	CapitalCityResult string `xml:"CapitalCityResult,omitempty"`
}

type CountryCurrency struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryCurrency"`

	SCountryISOCode string `xml:"sCountryISOCode,omitempty"`
}

type CountryCurrencyResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryCurrencyResponse"`

	CountryCurrencyResult *TCurrency `xml:"CountryCurrencyResult,omitempty"`
}

type CountryFlag struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryFlag"`

	SCountryISOCode string `xml:"sCountryISOCode,omitempty"`
}

type CountryFlagResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryFlagResponse"`

	CountryFlagResult string `xml:"CountryFlagResult,omitempty"`
}

type CountryIntPhoneCode struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryIntPhoneCode"`

	SCountryISOCode string `xml:"sCountryISOCode,omitempty"`
}

type CountryIntPhoneCodeResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountryIntPhoneCodeResponse"`

	CountryIntPhoneCodeResult string `xml:"CountryIntPhoneCodeResult,omitempty"`
}

type FullCountryInfo struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo FullCountryInfo"`

	SCountryISOCode string `xml:"sCountryISOCode,omitempty"`
}

type FullCountryInfoResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo FullCountryInfoResponse"`

	FullCountryInfoResult *TCountryInfo `xml:"FullCountryInfoResult,omitempty"`
}

type FullCountryInfoAllCountries struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo FullCountryInfoAllCountries"`
}

type FullCountryInfoAllCountriesResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo FullCountryInfoAllCountriesResponse"`

	FullCountryInfoAllCountriesResult *ArrayOftCountryInfo `xml:"FullCountryInfoAllCountriesResult,omitempty"`
}

type CountriesUsingCurrency struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountriesUsingCurrency"`

	SISOCurrencyCode string `xml:"sISOCurrencyCode,omitempty"`
}

type CountriesUsingCurrencyResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo CountriesUsingCurrencyResponse"`

	CountriesUsingCurrencyResult *ArrayOftCountryCodeAndName `xml:"CountriesUsingCurrencyResult,omitempty"`
}

type ListOfLanguagesByName struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfLanguagesByName"`
}

type ListOfLanguagesByNameResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfLanguagesByNameResponse"`

	ListOfLanguagesByNameResult *ArrayOftLanguage `xml:"ListOfLanguagesByNameResult,omitempty"`
}

type ListOfLanguagesByCode struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfLanguagesByCode"`
}

type ListOfLanguagesByCodeResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfLanguagesByCodeResponse"`

	ListOfLanguagesByCodeResult *ArrayOftLanguage `xml:"ListOfLanguagesByCodeResult,omitempty"`
}

type LanguageName struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo LanguageName"`

	SISOCode string `xml:"sISOCode,omitempty"`
}

type LanguageNameResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo LanguageNameResponse"`

	LanguageNameResult string `xml:"LanguageNameResult,omitempty"`
}

type LanguageISOCode struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo LanguageISOCode"`

	SLanguageName string `xml:"sLanguageName,omitempty"`
}

type LanguageISOCodeResponse struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo LanguageISOCodeResponse"`

	LanguageISOCodeResult string `xml:"LanguageISOCodeResult,omitempty"`
}

type TContinent struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo tContinent"`

	SCode string `xml:"sCode,omitempty"`

	SName string `xml:"sName,omitempty"`
}

type TCurrency struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo tCurrency"`

	SISOCode string `xml:"sISOCode,omitempty"`

	SName string `xml:"sName,omitempty"`
}

type TCountryCodeAndName struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo tCountryCodeAndName"`

	SISOCode string `xml:"sISOCode,omitempty"`

	SName string `xml:"sName,omitempty"`
}

type TCountryCodeAndNameGroupedByContinent struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo tCountryCodeAndNameGroupedByContinent"`

	Continent *TContinent `xml:"Continent,omitempty"`

	CountryCodeAndNames *ArrayOftCountryCodeAndName `xml:"CountryCodeAndNames,omitempty"`
}

type TCountryInfo struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo tCountryInfo"`

	SISOCode string `xml:"sISOCode,omitempty"`

	SName string `xml:"sName,omitempty"`

	SCapitalCity string `xml:"sCapitalCity,omitempty"`

	SPhoneCode string `xml:"sPhoneCode,omitempty"`

	SContinentCode string `xml:"sContinentCode,omitempty"`

	SCurrencyISOCode string `xml:"sCurrencyISOCode,omitempty"`

	SCountryFlag string `xml:"sCountryFlag,omitempty"`

	Languages *ArrayOftLanguage `xml:"Languages,omitempty"`
}

type TLanguage struct {
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo tLanguage"`

	SISOCode string `xml:"sISOCode,omitempty"`

	SName string `xml:"sName,omitempty"`
}

type ArrayOftCountryCodeAndName struct {
	TCountryCodeAndName []*TCountryCodeAndName `xml:"tCountryCodeAndName,omitempty"`
}

type ArrayOftLanguage struct {
	TLanguage []*TLanguage `xml:"tLanguage,omitempty"`
}

type ArrayOftContinent struct {
	TContinent []*TContinent `xml:"tContinent,omitempty"`
}

type ArrayOftCurrency struct {
	TCurrency []*TCurrency `xml:"tCurrency,omitempty"`
}

type ArrayOftCountryCodeAndNameGroupedByContinent struct {
	TCountryCodeAndNameGroupedByContinent []*TCountryCodeAndNameGroupedByContinent `xml:"tCountryCodeAndNameGroupedByContinent,omitempty"`
}

type ArrayOftCountryInfo struct {
	TCountryInfo []*TCountryInfo `xml:"tCountryInfo,omitempty"`
}

type CountryInfoServiceSoapType interface {
	/* Returns a list of continents ordered by name. */
	ListOfContinentsByName(request *ListOfContinentsByName) (*ListOfContinentsByNameResponse, error)

	/* Returns a list of continents ordered by code. */
	ListOfContinentsByCode(request *ListOfContinentsByCode) (*ListOfContinentsByCodeResponse, error)

	/* Returns a list of currencies ordered by name. */
	ListOfCurrenciesByName(request *ListOfCurrenciesByName) (*ListOfCurrenciesByNameResponse, error)

	/* Returns a list of currencies ordered by code. */
	ListOfCurrenciesByCode(request *ListOfCurrenciesByCode) (*ListOfCurrenciesByCodeResponse, error)

	/* Returns the name of the currency (if found) */
	CurrencyName(request *CurrencyName) (*CurrencyNameResponse, error)

	/* Returns a list of all stored counties ordered by ISO code */
	ListOfCountryNamesByCode(
		request *ListOfCountryNamesByCode,
	) (*ListOfCountryNamesByCodeResponse, error)

	/* Returns a list of all stored counties ordered by country name */
	ListOfCountryNamesByName(
		request *ListOfCountryNamesByName,
	) (*ListOfCountryNamesByNameResponse, error)

	/* Returns a list of all stored counties grouped per continent */
	ListOfCountryNamesGroupedByContinent(
		request *ListOfCountryNamesGroupedByContinent,
	) (*ListOfCountryNamesGroupedByContinentResponse, error)

	/* Searches the database for a country by the passed ISO country code */
	CountryName(request *CountryName) (*CountryNameResponse, error)

	/* This function tries to found a country based on the passed country name. */
	CountryISOCode(request *CountryISOCode) (*CountryISOCodeResponse, error)

	/* Returns the  name of the captial city for the passed country code */
	CapitalCity(request *CapitalCity) (*CapitalCityResponse, error)

	/* Returns the currency ISO code and name for the passed country ISO code */
	CountryCurrency(request *CountryCurrency) (*CountryCurrencyResponse, error)

	/* Returns a link to a picture of the country flag */
	CountryFlag(request *CountryFlag) (*CountryFlagResponse, error)

	/* Returns the internation phone code for the passed ISO country code */
	CountryIntPhoneCode(request *CountryIntPhoneCode) (*CountryIntPhoneCodeResponse, error)

	/* Returns a struct with all the stored country information. Pass the ISO country code */
	FullCountryInfo(request *FullCountryInfo) (*FullCountryInfoResponse, error)

	/* Returns an array with all countries and all the language information stored */
	FullCountryInfoAllCountries(
		request *FullCountryInfoAllCountries,
	) (*FullCountryInfoAllCountriesResponse, error)

	/* Returns a list of all countries that use the same currency code. Pass a ISO currency code */
	CountriesUsingCurrency(request *CountriesUsingCurrency) (*CountriesUsingCurrencyResponse, error)

	/* Returns an array of languages ordered by name */
	ListOfLanguagesByName(request *ListOfLanguagesByName) (*ListOfLanguagesByNameResponse, error)

	/* Returns an array of languages ordered by code */
	ListOfLanguagesByCode(request *ListOfLanguagesByCode) (*ListOfLanguagesByCodeResponse, error)

	/* Find a language name based on the passed ISO language code */
	LanguageName(request *LanguageName) (*LanguageNameResponse, error)

	/* Find a language ISO code based on the passed language name */
	LanguageISOCode(request *LanguageISOCode) (*LanguageISOCodeResponse, error)
}

type countryInfoServiceSoapType struct {
	client *soap.Client
}

func NewCountryInfoServiceSoapType(client *soap.Client) CountryInfoServiceSoapType {
	return &countryInfoServiceSoapType{
		client: client,
	}
}

func (service *countryInfoServiceSoapType) ListOfContinentsByName(
	request *ListOfContinentsByName,
) (*ListOfContinentsByNameResponse, error) {
	response := new(ListOfContinentsByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) ListOfContinentsByCode(
	request *ListOfContinentsByCode,
) (*ListOfContinentsByCodeResponse, error) {
	response := new(ListOfContinentsByCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) ListOfCurrenciesByName(
	request *ListOfCurrenciesByName,
) (*ListOfCurrenciesByNameResponse, error) {
	response := new(ListOfCurrenciesByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) ListOfCurrenciesByCode(
	request *ListOfCurrenciesByCode,
) (*ListOfCurrenciesByCodeResponse, error) {
	response := new(ListOfCurrenciesByCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) CurrencyName(
	request *CurrencyName,
) (*CurrencyNameResponse, error) {
	response := new(CurrencyNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) ListOfCountryNamesByCode(
	request *ListOfCountryNamesByCode,
) (*ListOfCountryNamesByCodeResponse, error) {
	response := new(ListOfCountryNamesByCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) ListOfCountryNamesByName(
	request *ListOfCountryNamesByName,
) (*ListOfCountryNamesByNameResponse, error) {
	response := new(ListOfCountryNamesByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) ListOfCountryNamesGroupedByContinent(
	request *ListOfCountryNamesGroupedByContinent,
) (*ListOfCountryNamesGroupedByContinentResponse, error) {
	response := new(ListOfCountryNamesGroupedByContinentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) CountryName(
	request *CountryName,
) (*CountryNameResponse, error) {
	response := new(CountryNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) CountryISOCode(
	request *CountryISOCode,
) (*CountryISOCodeResponse, error) {
	response := new(CountryISOCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) CapitalCity(
	request *CapitalCity,
) (*CapitalCityResponse, error) {
	response := new(CapitalCityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) CountryCurrency(
	request *CountryCurrency,
) (*CountryCurrencyResponse, error) {
	response := new(CountryCurrencyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) CountryFlag(
	request *CountryFlag,
) (*CountryFlagResponse, error) {
	response := new(CountryFlagResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) CountryIntPhoneCode(
	request *CountryIntPhoneCode,
) (*CountryIntPhoneCodeResponse, error) {
	response := new(CountryIntPhoneCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) FullCountryInfo(
	request *FullCountryInfo,
) (*FullCountryInfoResponse, error) {
	response := new(FullCountryInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) FullCountryInfoAllCountries(
	request *FullCountryInfoAllCountries,
) (*FullCountryInfoAllCountriesResponse, error) {
	response := new(FullCountryInfoAllCountriesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) CountriesUsingCurrency(
	request *CountriesUsingCurrency,
) (*CountriesUsingCurrencyResponse, error) {
	response := new(CountriesUsingCurrencyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) ListOfLanguagesByName(
	request *ListOfLanguagesByName,
) (*ListOfLanguagesByNameResponse, error) {
	response := new(ListOfLanguagesByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) ListOfLanguagesByCode(
	request *ListOfLanguagesByCode,
) (*ListOfLanguagesByCodeResponse, error) {
	response := new(ListOfLanguagesByCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) LanguageName(
	request *LanguageName,
) (*LanguageNameResponse, error) {
	response := new(LanguageNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *countryInfoServiceSoapType) LanguageISOCode(
	request *LanguageISOCode,
) (*LanguageISOCodeResponse, error) {
	response := new(LanguageISOCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
