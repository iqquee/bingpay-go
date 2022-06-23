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
1. You can use the below Go command to install Gin
```sh
$ go get -u github.com/hisyntax/bingpay-go
```
2. Import it in your code:
```sh
import "github.com/hisyntax/bingpay-go"
```
# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
```