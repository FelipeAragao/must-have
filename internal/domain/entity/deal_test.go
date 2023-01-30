package entity

import (
	"github.com/FelipeAragao/must-have/pkg/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDeal(t *testing.T) {
	typeDeal := VENDA
	value := 100000.00
	description := "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt"
	tradeFor := "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt"
	location := LocationDeal{}
	urgency := Urgency{}
	photos := []string{"https://www.google.com.br", "https://www.google.com.br"}

	location.Lat = -67.868968
	location.Lng = -12.943469
	location.Address = "460 Hamilton Avenue, Oberlin, Michigan, 2608"
	location.City = "São Luis"
	location.State = "MA"
	location.ZipCode = 11000000

	urgency.Type = MEDIA
	urgency.LimitDate, _ = time.Parse("2021-01-01", "2021-01-01")

	userID, _ := entity.ParseID("5f9f1c5b-7b1a-4b5d-8c1a-1c5b7b1a4b5d")
	deal, err := NewDeal(userID, typeDeal, value, description, tradeFor, location, urgency, photos)
	assert.Nil(t, err)
	assert.NotNil(t, deal)
	assert.Equal(t, deal.Type, typeDeal)
	assert.Equal(t, deal.Value, value)
	assert.Equal(t, deal.Description, description)
	assert.Equal(t, deal.TradeFor, tradeFor)
	assert.Equal(t, deal.Location, location)
	assert.Equal(t, deal.Urgency, urgency)
	assert.Equal(t, deal.Photos, photos)
}

func TestDeal_Modify(t *testing.T) {
	typeDeal := VENDA
	value := 100000.00
	description := "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt"
	tradeFor := "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt"
	location := LocationDeal{}
	urgency := Urgency{}
	photos := []string{"https://www.google.com.br", "https://www.google.com.br"}

	location.Lat = -67.868968
	location.Lng = -12.943469
	location.Address = "460 Hamilton Avenue, Oberlin, Michigan, 2608"
	location.City = "São Luis"
	location.State = "MA"
	location.ZipCode = 11000000

	urgency.Type = MEDIA
	urgency.LimitDate, _ = time.Parse("2021-01-01", "2021-01-01")

	userID, _ := entity.ParseID("5f9f1c5b-7b1a-4b5d-8c1a-1c5b7b1a4b5d")
	deal, _ := NewDeal(userID, typeDeal, value, description, tradeFor, location, urgency, photos)

	deal.Type = TROCA
	deal.Value = 200000.00
	deal.Description = "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt"

	err := deal.Modify()
	assert.Nil(t, err)
	assert.Equal(t, deal.Type, TROCA)
	assert.Equal(t, deal.Value, 200000.00)
	assert.Equal(t, deal.Description, "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt")
}
