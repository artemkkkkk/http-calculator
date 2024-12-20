package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/artemkkkkk/http-calculator/pkg/calculate"
	"github.com/artemkkkkk/http-calculator/pkg/custom_errors"
)

func ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var expr Expression

	if err := json.NewDecoder(r.Body).Decode(&expr); err != nil {
		http.Error(w, "Failed to encoding JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	result, err := calculate.Eval(expr.Expression)

	if err == nil {
		var resp = map[string]string{"result": fmt.Sprint(result)}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)

	} else if errors.Is(err, custom_errors.InvalidExpression) || errors.Is(err, custom_errors.UnknownOperator) {
		var resp = map[string]string{"error": "Expression is not valid"}

		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(resp)

	} else {
		var resp = map[string]string{"error": "Internal server error"}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(resp)
	}
}
