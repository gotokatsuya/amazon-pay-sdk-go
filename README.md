# amazon-pay-sdk-go

amazon-pay-sdk-go is a Go client library for accessing the [Amazon Pay API](https://amazonpaycheckoutintegrationguide.s3.amazonaws.com/amazon-pay-checkout/introduction.html).

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
