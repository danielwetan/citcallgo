# Citcallgo
We can consume Citcall API using this library. You need Citcall Apikey and your IP has been whitelisted. Please refer to official [Citcall documentation](https://docs.citcall.com/) for more detail.

---

## Table Of Content
* [Features](https://github.com/danielwetan/citcallgo#features)
* [Installation](https://github.com/danielwetan/citcallgo#installation)
* [Usage](https://github.com/danielwetan/citcallgo#usage)
* [Examples](https://github.com/danielwetan/citcallgo#examples)
    * [Misscall OTP](https://github.com/danielwetan/citcallgo#misscall-otp)
    * [SMS](https://github.com/danielwetan/citcallgo#sms)
    * [SMS OTP](https://github.com/danielwetan/citcallgo#sms-otp)
* [Contribute](https://github.com/danielwetan/citcallgo#contribute)

## Features
- [x] [Misscall OTP](https://docs.citcall.com/?j#assynchronous-miscall)
- [x] [SMS](https://docs.citcall.com/?j#sms)
- [x] [SMS OTP](https://docs.citcall.com/?j#sms-otp)

## Installation

```bash=
go get github.com/danielwetan/citcallgo
```


## Usage
Import the library in project file
```go=
import "github.com/danielwetan/citcallgo"
```

Create a new instance with your Citcall Apikey
```go=
citcall := citcallgo.New("API_KEY")
```

## Examples

#### Misscall OTP
```go=
package main

import (
	"context"
	"fmt"

	"github.com/danielwetan/citcallgo"
)

func main() {
	citcall := citcallgo.New("API_KEY")

	requestBody := citcallgo.MisscallOtpRequest{
		Msisdn: "6281234567890",
		Gateway: 0,
	}

	res, err := citcall.SendMisscall(context.Background(), &requestBody)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	fmt.Println(res)
}
```

#### SMS
```go=
package main

import (
	"context"
	"fmt"

	"github.com/danielwetan/citcallgo"
)

func main() {
	citcall := citcallgo.New("API_KEY")

	requestBody := citcallgo.SMSRequest{
		Msisdn:   "6281234567890",
		SenderId: "xxxxx",
		Text:     "hello",
	}

	res, err := citcall.SendSMS(context.Background(), &requestBody)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	fmt.Println(res)
}
```

#### SMS OTP
```go=
package main

import (
	"context"
	"fmt"

	"github.com/danielwetan/citcallgo"
)

func main() {
	citcall := citcallgo.New("API_KEY")

	requestBody := citcallgo.SMSOTPRequest{
		Msisdn:   "6281234567890",
		SenderId: "xxxxx",
		Text:     "hello",
	}

	res, err := citcall.SendSMSOTP(context.Background(), &requestBody)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	fmt.Println(res)
}
```


## Contribute
1. You can find a bug or issue in this page [Github Issues](http://github.com/danielwetan/citcallgo/issues)
2. Pull this repository
3. Make a change
4. Send a pull request with clear commit message