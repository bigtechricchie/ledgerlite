package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Parse pounds and pence without converting through floating-point.
func parseAmount(input string) (int64, error) {
	if strings.HasPrefix(input, "-") {
		return 0, errors.New(
			"transaction amount must be greater than zero",
		)
	}

	parts := strings.Split(input, ".")

	if len(parts) > 2 {
		return 0, errors.New(
			"transaction amount must be a valid currency amount",
		)
	}

	if parts[0] == "" {
		return 0, errors.New(
			"transaction amount must include pounds",
		)
	}

	pounds, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, errors.New(
			"transaction amount must be a valid currency amount",
		)
	}

	var pence int64

	if len(parts) == 2 {
		penceText := parts[1]

		if len(penceText) == 0 {
			penceText = "00"
		}

		if len(penceText) == 1 {
			penceText += "0"
		}

		if len(penceText) > 2 {
			return 0, errors.New(
				"transaction amount cannot have more than two decimal places",
			)
		}

		pence, err = strconv.ParseInt(penceText, 10, 64)
		if err != nil {
			return 0, errors.New(
				"transaction amount must be a valid currency amount",
			)
		}
	}

	amount := pounds*100 + pence

	if amount <= 0 {
		return 0, errors.New(
			"transaction amount must be greater than zero",
		)
	}

	return amount, nil
}

func formatAmount(amount int64) string {
	pounds := amount / 100
	pence := amount % 100

	return fmt.Sprintf("£%d.%02d", pounds, pence)
}
