package server

import (
	"context"
	"encoding/json"
	"fincal-server/.gen/fincal/public/model"
	. "fincal-server/.gen/fincal/public/table"
	"fincal-server/lib"
	"github.com/roger-king/go-getenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net/http"
)

func (s *InternalServer) Setup(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := populateCurrency()

	if err != nil {
		lib.Sugar.Error("Failed to populate currency", err)
		return nil, status.Errorf(codes.Unknown, "%s", err)
	}

	return &emptypb.Empty{}, nil
}

type CurrencyS struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Code   string `json:"code"`
}

type CurrencyReq struct {
	Data map[string]CurrencyS `json:"data"`
}

func populateCurrency() error {
	client := &http.Client{}

	uri := "https://api.freecurrencyapi.com/v1/currencies"

	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("apiKey", getenv.EnvOrDefault("FREE_CURRENCY_API_KEY", getenv.String("")))

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var target CurrencyReq

	err = json.NewDecoder(resp.Body).Decode(&target)

	var items []model.Currency

	for _, e := range target.Data {
		items = append(items, model.Currency{
			Code:         e.Code,
			Symbol:       e.Symbol,
			ExchangeRate: 1,
			LastUpdated:  nil,
		})
	}

	stm := Currency.INSERT(Currency.AllColumns).MODELS(items).ON_CONFLICT().DO_NOTHING()

	_, err = stm.Exec(lib.DB)

	if err != nil {
		return err
	}

	return nil
}
