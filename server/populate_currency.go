package server

import (
	"context"
	"encoding/json"
	. "fincal-server/.gen/fincal/public/table"
	"fincal-server/lib"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/roger-king/go-getenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net/http"
)

func (s *InternalServer) PopulateCurrency(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := populateExchangeRate()

	if err != nil {
		lib.Sugar.Error("Failed to populate exchange rate", err)
		return nil, status.Errorf(codes.Unknown, "%s", err)
	}

	return &emptypb.Empty{}, nil
}

type ExchangeReq struct {
	Data map[string]float64 `json:"data"`
}

func populateExchangeRate() error {
	client := &http.Client{}

	uri := "https://api.freecurrencyapi.com/v1/latest"

	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("apiKey", getenv.EnvOrDefault("FREE_CURRENCY_API_KEY", getenv.String("")))

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	//d, _ := io.ReadAll(resp.Body)
	//
	//print(string(d))

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var target ExchangeReq

	err = json.NewDecoder(resp.Body).Decode(&target)

	for c, e := range target.Data {
		stm := Currency.UPDATE(Currency.ExchangeRate).SET(Float(e)).WHERE(Currency.Code.EQ(String(c)))

		_, err = stm.Exec(lib.DB)

		if err != nil {
			return err
		}
	}

	return nil
}
