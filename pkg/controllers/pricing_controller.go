package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dieunb/go_training/pkg/models"
)

type PricingIndexData struct {
	Description string           `json:"description"`
	Data        []models.Pricing `json:"data"`
}

func PricingIndex(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(pricingData())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func pricingData() PricingIndexData {
	data := PricingIndexData{
		Description: "Quickly build an effective pricing table for your potential customers with this Bootstrap example. It’s built with default Bootstrap components and utilities with little customization.",
		Data: []models.Pricing{
			models.Pricing{
				Title:  "Free",
				Amount: 0,
				Features: []models.Feature{
					models.Feature{
						Name: "10 users included",
					},
					models.Feature{
						Name: "2 GB of storage",
					},
					models.Feature{
						Name: "Email support",
					},
					models.Feature{
						Name: "Help center access",
					},
				},
			},
			models.Pricing{
				Title:  "Pro",
				Amount: 15,
				Features: []models.Feature{
					models.Feature{
						Name: "20 users included",
					},
					models.Feature{
						Name: "10 GB of storage",
					},
					models.Feature{
						Name: "Priority email support",
					},
					models.Feature{
						Name: "Help center access",
					},
				},
			},
			models.Pricing{
				Title:  "Enterprise",
				Amount: 29,
				Features: []models.Feature{
					models.Feature{
						Name: "30 users included",
					},
					models.Feature{
						Name: "15 GB of storage",
					},
					models.Feature{
						Name: "Phone and email support",
					},
					models.Feature{
						Name: "Help center access",
					},
				},
			},
		},
	}

	return data
}
