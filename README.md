# amazon-pay-sdk-go

amazon-pay-sdk-go is a Go client library for accessing the [Amazon Pay API](https://developer.amazon.com/ja/docs/amazon-pay/intro.html).

## Usage

```go
import "github.com/gotokatsuya/amazon-pay-sdk-go/amazonpay"

func main() {
    pay, err := amazonpay.New("..")
    ...
}
```

## License

This library is distributed under the MIT license.
