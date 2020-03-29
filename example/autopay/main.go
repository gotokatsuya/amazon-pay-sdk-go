package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gotokatsuya/amazon-pay-sdk-go/amazonpay"
	"github.com/rs/xid"
)

var (
	accessKey = os.Getenv("AMAZON_PAY_ACCESS_KEY")
	secretKey = os.Getenv("AMAZON_PAY_SECRET_KEY")
	sellerID  = os.Getenv("AMAZON_PAY_SELLER_ID")
	clientID  = os.Getenv("AMAZON_PAY_CLIENT_ID")
)

func main() {
	amazonpayCli, err := amazonpay.New(accessKey, secretKey, sellerID, amazonpay.WithSandbox())
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/payment/buy", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		amazonBillingAgreementID := r.PostFormValue("amazonBillingAgreementId")
		if _, _, err := amazonpayCli.SetBillingAgreementDetails(ctx, &amazonpay.SetBillingAgreementDetailsRequest{
			AmazonBillingAgreementID: amazonBillingAgreementID,
			BillingAgreementAttributes: amazonpay.BillingAgreementAttributes{
				SellerNote: "有料 12ヶ月プラン",
			},
		}); err != nil {
			log.Printf("SetBillingAgreementDetails: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, _, err := amazonpayCli.ConfirmBillingAgreement(ctx, &amazonpay.ConfirmBillingAgreementRequest{
			AmazonBillingAgreementID: amazonBillingAgreementID,
		}); err != nil {
			log.Printf("ConfirmBillingAgreement: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		authorizationReferenceID := xid.New().String()
		authorizeResp, _, err := amazonpayCli.AuthorizeOnBillingAgreement(ctx, &amazonpay.AuthorizeOnBillingAgreementRequest{
			AmazonBillingAgreementID: amazonBillingAgreementID,
			AuthorizationReferenceID: authorizationReferenceID,
			AuthorizationAmount: amazonpay.Price{
				Amount:       "10.00",
				CurrencyCode: "JPY",
			},
			SellerAuthorizationNote: "有料 12ヶ月プラン",
			TransactionTimeout:      0,
			CaptureNow:              true,
		})
		if err != nil {
			log.Printf("AuthorizeOnBillingAgreement: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		amazonCaptureIDs := authorizeResp.AuthorizeOnBillingAgreementResult.AuthorizationDetails.IDList
		log.Printf("AuthorizeOnBillingAgreement.amazonCaptureIDs: %v\n", amazonCaptureIDs)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(&struct {
			OK bool
		}{OK: true}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			AmazonPayClientID string
			AmazonPaySellerID string
		}{
			AmazonPayClientID: clientID,
			AmazonPaySellerID: sellerID,
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
