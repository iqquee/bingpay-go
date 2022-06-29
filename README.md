# Bingpay-go
This is a Go library that allows you to integrate BINGPAY into your Go project. Bingpay is a FinTech company which allows for Cheap Airtime &amp; Data Topup, Send &amp; Recieve Cash, Pay Utility Bills, Purchase Giftcards, Trade Airtime, Paypal &amp; Giftcards and International Topup.

# Get Started
- In other to use this package, you need to first create an account with bingpay via https://bingpay.ng/register
- After your account have been successfully created, you would need an API Token which is used to authorize users to use the bingpay api.

## How to get the Bingpay API token
- After you signin to your Bingpay dashboard, click on the top right corner which has your profile picture and username
- On the pop up, click on the "developer api" option and you would get a form to fill which would request for two things: your website address and the reason why you need the developer API. After those have been inputed then you can submit your request which would be under review for some day(s) before you get an API token. 

# Installation
To install this bingpay package, you need to install [Go](https://golang.org/) and set your Go workspace first.
1. You can use the below Go command to install BingPay
```sh
$ go get -u github.com/hisyntax/bingpay-go
```
2. Import it in your code:
```sh
import "github.com/hisyntax/bingpay-go"
```
## Note : All methods in this package returns three (3) things:
- The object of the response
- An int (status code) i.e  status 200 or status 400
- An error (if any)

# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```
## Wallet 
- ### Check Balance
Use this to fetch your bingpay wallet balance
```go
package main

import (
	"fmt"
	bingpay "github.com/hisyntax/bingpay-go"
	"github.com/hisyntax/bingpay-go/bingpayBal"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	response, status, err := bingpayBal.CheckBalance()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
## Airtime
- ### All Networks
Use this to fetch the list of all networks supported by bingpay

```go
package main

import (
	"fmt"
	bingpay "github.com/hisyntax/bingpay-go"
	"github.com/hisyntax/bingpay-go/airtime"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	response, status, err := airtime.AllNetworks()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Verify Phone Number
Use this to verify customer's phone number.
Only country ISO codes are allowed. example NG for Nigeria, US for United States of America etc.
You can find list of all ISO codes [here](https://www.nationsonline.org/oneworld/country_code_list.htm)

```go
package main

import (
	"fmt"
	bingpay "github.com/hisyntax/bingpay-go"
	"github.com/hisyntax/bingpay-go/airtime"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	country := "NG" //for nigeria
	number := "08000000000"
	response, status, err := airtime.VerifyPhoneNumber(country, number)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Buy Airtime
Use this to perform airtime purchase.
You get 2% discount on every airtime purchase instantly.

```go
package main

import (
	"fmt"
	bingpay "github.com/hisyntax/bingpay-go"
	"github.com/hisyntax/bingpay-go/airtime"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	phone := "08000000000"
	amount := 100 //for 100 naria topup
	network_id := 1 //for MTN
	response, status, err := airtime.BuyAirtime(phone,amount,network_id)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```

## Data
- ### All data plans
Use this to fetch all data plans

```go
package main

import (
	"fmt"
	bingpay "github.com/hisyntax/bingpay-go"
	"github.com/hisyntax/bingpay-go/data"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	response, status, err := data.AllDataPlans()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Data plans
Fetch Data plans for a specific network.
Network id can be gotten from All networks endpoint.

```go
package main

import (
	"fmt"
	bingpay "github.com/hisyntax/bingpay-go"
	"github.com/hisyntax/bingpay-go/data"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	/* network ID's
			MTN = 1
			Airtel = 2
			9mobile = 3
			Glo = 4
	*/

	network_id := 1 //for MTN
	response, status, err := data.DataPlans(network_id)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```
- ### Buy Data
Use this to perform a data purchase.
You get 2% discount on every data purchase instantly.

```go
package main

import (
	"fmt"
	bingpay "github.com/hisyntax/bingpay-go"
	"github.com/hisyntax/bingpay-go/data"
)

func main() {
	bingpay.Token = "your bingpay api secret key" // to add your secret key for the api requests 

	/* network ID's
			MTN = 1
			Airtel = 2
			9mobile = 3
			Glo = 4
	*/
	phone := "08000000000"
	plan := 
	network_id := 1 //for MTN
	response, status, err := data.BuyData(1)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(status)
	fmt.Println(response)
}
```

```sh
# run example.go 
$ go run example.go
```