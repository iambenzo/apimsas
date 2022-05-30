[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/iambenzo/apimsas?tab=doc)
[![go report card](https://goreportcard.com/badge/github.com/iambenzo/apimsas "go report")](https://goreportcard.com/report/github.com/iambenzo/apimsas)
[![Test](https://github.com/iambenzo/apimsas/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/iambenzo/apimsas/actions/workflows/test.yml)


This module exposes a Provider struct which can be used to generate SAS tokens for authenticating with Azure's APIM Management API.

## Setup

If you haven't already, enable APIM's Management API by navigating to your APIM resource in the Azure Portal and using the pane's navigation bar on the left to find and select "Management API". From there, toggle "Enable Management REST API" to `Yes`.

Microsoft's documentation for the above steps can be found [here](https://docs.microsoft.com/en-gb/rest/api/apimanagement/apimanagementrest/api-management-rest?WT.mc_id=Portal-fx).

Once activated, on the same page, locate the "Credentials" section and take note of your `Identifier` and either your `Primary key`, or `Secondary key`. These will be used when instantiating the Provider struct as `id` and `key` respectively.

## Usage

```go
package main

import (
	"github.com/iambenzo/apimsas"
	"fmt"
)

func main() {

    // Instantiate token provider
	sas := apimsas.NewApimSasProvider("identifier", "primary/secondary key")

    // Generate token
	token, err := sas.GetSasToken()

	if err != nil {
		fmt.Printf("%v\n", err)
	}

    // Show off our generated token
	fmt.Printf("%s\n", token)

    // Make API requests from here...
}
```

Microsoft's documentation on the API Management REST API can be found [here](https://docs.microsoft.com/en-us/rest/api/apimanagement/).

