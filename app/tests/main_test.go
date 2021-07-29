package main_tests

import (
	"card_api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Api_DeckHandler_GET_returns_cards(t *testing.T) {

	req, err := http.NewRequest("GET", "/deck", nil)
	if err != nil {
		t.Fatal(err)
	}

	requstRecorder := httptest.NewRecorder()

	handler := handlers.NewStandardDeckHandler()

	reqHandler := http.HandlerFunc(handler.HandleDeck)

	reqHandler.ServeHTTP(requstRecorder, req)

	if status := requstRecorder.Code; status != http.StatusOK {
		t.Errorf("handler got other status code than expected: got %v expected %v", status, http.StatusOK)
	}

	expectedresult := `[{"ID":1,"Suit":"heart","Value":"ace"},{"ID":2,"Suit":"heart","Value":"two"},{"ID":3,"Suit":"heart","Value":"three"},{"ID":4,"Suit":"heart","Value":"four"},{"ID":5,"Suit":"heart","Value":"five"},{"ID":6,"Suit":"heart","Value":"six"},{"ID":7,"Suit":"heart","Value":"seven"},{"ID":8,"Suit":"heart","Value":"eight"},{"ID":9,"Suit":"heart","Value":"nine"},{"ID":10,"Suit":"heart","Value":"ten"},{"ID":11,"Suit":"heart","Value":"jack"},{"ID":12,"Suit":"heart","Value":"queen"},{"ID":13,"Suit":"heart","Value":"king"},{"ID":14,"Suit":"diamond","Value":"ace"},{"ID":15,"Suit":"diamond","Value":"two"},{"ID":16,"Suit":"diamond","Value":"three"},{"ID":17,"Suit":"diamond","Value":"four"},{"ID":18,"Suit":"diamond","Value":"five"},{"ID":19,"Suit":"diamond","Value":"six"},{"ID":20,"Suit":"diamond","Value":"seven"},{"ID":21,"Suit":"diamond","Value":"eight"},{"ID":22,"Suit":"diamond","Value":"nine"},{"ID":23,"Suit":"diamond","Value":"ten"},{"ID":24,"Suit":"diamond","Value":"jack"},{"ID":25,"Suit":"diamond","Value":"queen"},{"ID":26,"Suit":"diamond","Value":"king"},{"ID":27,"Suit":"clover","Value":"ace"},{"ID":28,"Suit":"clover","Value":"two"},{"ID":29,"Suit":"clover","Value":"three"},{"ID":30,"Suit":"clover","Value":"four"},{"ID":31,"Suit":"clover","Value":"five"},{"ID":32,"Suit":"clover","Value":"six"},{"ID":33,"Suit":"clover","Value":"seven"},{"ID":34,"Suit":"clover","Value":"eight"},{"ID":35,"Suit":"clover","Value":"nine"},{"ID":36,"Suit":"clover","Value":"ten"},{"ID":37,"Suit":"clover","Value":"jack"},{"ID":38,"Suit":"clover","Value":"queen"},{"ID":39,"Suit":"clover","Value":"king"},{"ID":40,"Suit":"spade","Value":"ace"},{"ID":41,"Suit":"spade","Value":"two"},{"ID":42,"Suit":"spade","Value":"three"},{"ID":43,"Suit":"spade","Value":"four"},{"ID":44,"Suit":"spade","Value":"five"},{"ID":45,"Suit":"spade","Value":"six"},{"ID":46,"Suit":"spade","Value":"seven"},{"ID":47,"Suit":"spade","Value":"eight"},{"ID":48,"Suit":"spade","Value":"nine"},{"ID":49,"Suit":"spade","Value":"ten"},{"ID":50,"Suit":"spade","Value":"jack"},{"ID":51,"Suit":"spade","Value":"queen"},{"ID":52,"Suit":"spade","Value":"king"}]`

	if len(requstRecorder.Body.String()) <= 0 {
		t.Errorf("body was empty should have contained something")
	}

	if requstRecorder.Body.String() != expectedresult {
		t.Errorf("expected %s but got %s", expectedresult, requstRecorder.Body.String())
	}
}
func Test_Api_DeckHandler_id_is_one_GET_returns_correct_card(t *testing.T) {
	req, err := http.NewRequest("GET", "/deck/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	requstRecorder := httptest.NewRecorder()

	handler := handlers.NewStandardDeckHandler()

	reqHandler := http.HandlerFunc(handler.DeckOptions)

	reqHandler.ServeHTTP(requstRecorder, req)

	if status := requstRecorder.Code; status != http.StatusOK {
		t.Errorf("handler got other status code than expected: got %v expected %v", status, http.StatusOK)
	}

	expectedResponse := `{"ID":1,"Suit":"heart","Value":"ace"}`

	if requstRecorder.Body.String() != expectedResponse {
		t.Errorf("expected %s but got %s", expectedResponse, requstRecorder.Body.String())
	}
}

func Test_Api_DeckHandler_id_is_0_GET_returns_not_found(t *testing.T) {
	req, err := http.NewRequest("GET", "/deck/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	requstRecorder := httptest.NewRecorder()

	handler := handlers.NewStandardDeckHandler()

	reqHandler := http.HandlerFunc(handler.DeckOptions)

	reqHandler.ServeHTTP(requstRecorder, req)

	if status := requstRecorder.Code; status != http.StatusNotFound {
		t.Errorf("handler got other status code than expected: got %v expected %v", status, http.StatusNotFound)
	}

}
