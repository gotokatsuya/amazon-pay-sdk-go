package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gotokatsuya/amazon-pay-sdk-go/amazonpay"
)

// envs
var (
	publicKeyID    = os.Getenv("AMAZON_PAY_PUBLIC_KEY_ID")
	privateKeyPath = os.Getenv("AMAZON_PAY_PRIVATE_KEY_PATH")
	storeID        = os.Getenv("AMAZON_PAY_STORE_ID")
	merchantID     = os.Getenv("AMAZON_PAY_MERCHANT_ID")
)

// local datastore
var (
	chargePermissionID string
)

func main() {
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}

	amazonpayCli, err := amazonpay.New(publicKeyID, privateKeyData, "jp", true, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req := &amazonpay.CreateCheckoutSessionRequest{
			WebCheckoutDetails: &amazonpay.WebCheckoutDetails{
				CheckoutReviewReturnURL: "http://localhost:8000/review",
			},
			StoreID:              storeID,
			ChargePermissionType: "Recurring",
			RecurringMetadata: &amazonpay.RecurringMetadata{
				Frequency: &amazonpay.Frequency{
					Unit:  "Month",
					Value: "12",
				},
				Amount: &amazonpay.Price{
					Amount:       "100",
					CurrencyCode: "JPY",
				},
			},
		}
		payload, err := req.ToPayload()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		signature, err := amazonpayCli.GenerateButtonSignature(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			AmazonPayPayload     string
			AmazonPaySignature   string
			AmazonPayPublicKeyID string
			AmazonPayMerchantID  string
		}{
			AmazonPayPayload:     payload,
			AmazonPaySignature:   signature,
			AmazonPayPublicKeyID: publicKeyID,
			AmazonPayMerchantID:  merchantID,
		}
		if err := template.Must(template.ParseFiles("index.html")).Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/review", func(w http.ResponseWriter, r *http.Request) {
		checkoutSessionID := r.URL.Query().Get("amazonCheckoutSessionId")
		resp, httpResp, err := amazonpayCli.GetCheckoutSession(r.Context(), checkoutSessionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		switch httpResp.StatusCode {
		case http.StatusOK, http.StatusCreated:
			data := struct {
				CheckoutSessionID string
				PaymentDescriptor string
			}{
				CheckoutSessionID: resp.CheckoutSessionID,
				PaymentDescriptor: resp.PaymentPreferences[0].PaymentDescriptor,
			}
			if err := template.Must(template.ParseFiles("review.html")).Execute(w, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		default:
			http.Error(w, resp.ErrorResponse.ReasonCode+" | "+resp.ErrorResponse.Message, http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/approve", func(w http.ResponseWriter, r *http.Request) {
		checkoutSessionID := r.URL.Query().Get("amazonCheckoutSessionId")
		resp, httpResp, err := amazonpayCli.UpdateCheckoutSession(r.Context(), checkoutSessionID, &amazonpay.UpdateCheckoutSessionRequest{
			WebCheckoutDetails: &amazonpay.WebCheckoutDetails{
				CheckoutResultReturnURL: "http://localhost:8000/confirm",
			},
			PaymentDetails: &amazonpay.PaymentDetails{
				PaymentIntent:                 "AuthorizeWithCapture",
				CanHandlePendingAuthorization: amazonpay.Bool(false),
				ChargeAmount: &amazonpay.Price{
					Amount:       "100",
					CurrencyCode: "JPY",
				},
			},
			MerchantMetadata: &amazonpay.MerchantMetadata{
				NoteToBuyer: "Testing plan",
			},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		switch httpResp.StatusCode {
		case http.StatusOK, http.StatusCreated:
			http.Redirect(w, r, resp.WebCheckoutDetails.AmazonPayRedirectURL, http.StatusFound)
		default:
			http.Error(w, resp.ErrorResponse.ReasonCode+" | "+resp.ErrorResponse.Message, http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/confirm", func(w http.ResponseWriter, r *http.Request) {
		checkoutSessionID := r.URL.Query().Get("amazonCheckoutSessionId")
		resp, httpResp, err := amazonpayCli.CompleteCheckoutSession(r.Context(), checkoutSessionID, &amazonpay.CompleteCheckoutSessionRequest{
			ChargeAmount: &amazonpay.Price{
				Amount:       "100",
				CurrencyCode: "JPY",
			},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		switch httpResp.StatusCode {
		case http.StatusOK, http.StatusCreated:
			log.Println("confirm: " + resp.StatusDetails.State)
			switch resp.StatusDetails.State {
			case "Open":
			case "Completed":
				// TODO should save to database
				log.Println(resp.ChargeID)
				log.Println(resp.ChargePermissionID)
				chargePermissionID = resp.ChargePermissionID
			case "Canceled":
			}
			data := struct{}{}
			if err := template.Must(template.ParseFiles("confirm.html")).Execute(w, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		default:
			http.Error(w, resp.ErrorResponse.ReasonCode+" | "+resp.ErrorResponse.Message, http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/recurring", func(w http.ResponseWriter, r *http.Request) {
		cpResp, httpResp, err := amazonpayCli.GetChargePermission(r.Context(), chargePermissionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		switch httpResp.StatusCode {
		case http.StatusOK, http.StatusCreated:
			log.Println("recurring: " + cpResp.StatusDetails.State)
			switch cpResp.StatusDetails.State {
			case "Chargeable":
				cResp, httpResp, err := amazonpayCli.CreateCharge(r.Context(), &amazonpay.CreateChargeRequest{
					ChargePermissionID: chargePermissionID,
					ChargeAmount: &amazonpay.Price{
						Amount:       "100",
						CurrencyCode: "JPY",
					},
					CaptureNow:                    amazonpay.Bool(true),
					CanHandlePendingAuthorization: amazonpay.Bool(false),
				})
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				log.Println(httpResp.StatusCode)
				switch httpResp.StatusCode {
				case http.StatusOK, http.StatusCreated:
					log.Println("Success /recurring")
					w.WriteHeader(http.StatusOK)
				default:
					http.Error(w, cResp.ErrorResponse.ReasonCode+" | "+cResp.ErrorResponse.Message, http.StatusInternalServerError)
				}
			case "NonChargeable":
			case "Closed":
			}
		default:
			http.Error(w, cpResp.ErrorResponse.ReasonCode+" | "+cpResp.ErrorResponse.Message, http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/recurring/close", func(w http.ResponseWriter, r *http.Request) {
		resp, httpResp, err := amazonpayCli.CloseChargePermission(r.Context(), chargePermissionID, &amazonpay.CloseChargePermissionRequest{
			ClosureReason:        "closing reason",
			CancelPendingCharges: amazonpay.Bool(false),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		switch httpResp.StatusCode {
		case http.StatusOK, http.StatusCreated:
			log.Println("Success /recurring/close")
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, resp.ErrorResponse.ReasonCode+" | "+resp.ErrorResponse.Message, http.StatusInternalServerError)
		}
	})

	fmt.Println("http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
