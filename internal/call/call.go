package call

import (
	"fmt"
	"log/slog"

	"soap/internal/countryInfo"

	"github.com/hooklift/gowsdl/soap"
	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()

	countryClient := soap.NewClient(
		"http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso",
	)

	countryService := countryInfo.NewCountryInfoServiceSoapType(countryClient)

	fullCountryReq := &countryInfo.CountriesUsingCurrency{SISOCurrencyCode: "EUR"}
	fullCountryResp, err := countryService.CountriesUsingCurrency(fullCountryReq)
	if err != nil {
		slog.ErrorContext(ctx, "could not complete request", "error", err)

		return err
	}

	for i := range fullCountryResp.CountriesUsingCurrencyResult.TCountryCodeAndName {
		fmt.Printf("Counntry Code: %s, Country Name: %s\n",
			fullCountryResp.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SISOCode,
			fullCountryResp.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SName,
		)
	}

	return nil
}
