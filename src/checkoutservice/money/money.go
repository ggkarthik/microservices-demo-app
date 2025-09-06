// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package money

import (
	"errors"

	pb "github.com/GoogleCloudPlatform/microservices-demo/src/checkoutservice/genproto"
)

const (
	nanosMin = -999999999
	nanosMax = +999999999
	nanosMod = 1000000000
)

var (
	// ErrInvalidValue is returned when a money value is invalid
	ErrInvalidValue        = errors.New("one of the specified money values is invalid")
	// ErrMismatchingCurrency is returned when currencies don't match
	ErrMismatchingCurrency = errors.New("mismatching currency codes")
)

// IsValid checks if specified value has a valid units/nanos signs and ranges.
func IsValid(m pb.Money) bool {
	return signMatches(m) && validNanos(m.GetNanos())
}

func signMatches(m pb.Money) bool {
	return m.GetNanos() == 0 || m.GetUnits() == 0 || (m.GetNanos() < 0) == (m.GetUnits() < 0)
}

func validNanos(nanos int32) bool {
	return nanosMin <= nanos && nanos <= nanosMax
}

// IsZero returns true if the specified money value is zero.
func IsZero(m pb.Money) bool {
	return m.GetUnits() == 0 && m.GetNanos() == 0
}

// IsPositive returns true if the specified money value is positive.
func IsPositive(m pb.Money) bool {
	return IsValid(m) && (m.GetUnits() > 0 || (m.GetUnits() == 0 && m.GetNanos() > 0))
}

// IsNegative returns true if the specified money value is negative.
func IsNegative(m pb.Money) bool {
	return IsValid(m) && (m.GetUnits() < 0 || (m.GetUnits() == 0 && m.GetNanos() < 0))
}

// AreSameCurrency returns true if values a and b have a currency code and they are the same values.
func AreSameCurrency(a, b pb.Money) bool {
	return a.GetCurrencyCode() == b.GetCurrencyCode() && a.GetCurrencyCode() != ""
}

// AreEquals returns true if values a and b are the same. This only works when values
// are of the same currency or don't have currency code.
func AreEquals(a, b pb.Money) bool {
	return a.GetUnits() == b.GetUnits() && a.GetNanos() == b.GetNanos()
}

// Negate returns the same amount with the sign negated.
func Negate(m pb.Money) (pb.Money, error) {
	if !IsValid(m) {
		return pb.Money{}, ErrInvalidValue
	}
	return pb.Money{
		Units:        -m.GetUnits(),
		Nanos:        -m.GetNanos(),
		CurrencyCode: m.GetCurrencyCode()}, nil
}

// Must panics if the error is not nil. This can be used with other functions like:
// money.Must(money.Sum(a,b))
func Must(v pb.Money, err error) pb.Money {
	if err != nil {
		panic(err)
	}
	return v
}

// Sum adds two values. Returns an error if currency codes are not matching or resulting value is invalid.
func Sum(a, b pb.Money) (pb.Money, error) {
	if !IsValid(a) || !IsValid(b) {
		return pb.Money{}, ErrInvalidValue
	} else if a.GetCurrencyCode() != b.GetCurrencyCode() {
		return pb.Money{}, ErrMismatchingCurrency
	}

	units := a.GetUnits() + b.GetUnits()
	nanos := a.GetNanos() + b.GetNanos()

	if (units == 0 && nanos == 0) || (units > 0 && nanos >= 0) || (units < 0 && nanos <= 0) {
		// same sign <units, nanos>
		units += int64(nanos / nanosMod)
		nanos = nanos % nanosMod
	} else {
		// different sign. nanos guaranteed to not be 0 by this time.
		if units > 0 {
			units--
			nanos += nanosMod
		} else {
			units++
			nanos -= nanosMod
		}
	}

	return pb.Money{
		Units:        units,
		Nanos:        nanos,
		CurrencyCode: a.GetCurrencyCode()}, nil
}

// MultiplySlow is a slow multiplication operation done through adding the value to itself n-1 times.
func MultiplySlow(m pb.Money, n uint32) (pb.Money, error) {
	if !IsValid(m) {
		return pb.Money{}, ErrInvalidValue
	}
	if n == 0 {
		return pb.Money{
			Units:        0,
			Nanos:        0,
			CurrencyCode: m.GetCurrencyCode(),
		}, nil
	}
	if n == 1 {
		return m, nil
	}

	product, err := m, error(nil)
	for i := uint32(1); i < n; i++ {
		product, err = Sum(product, m)
		if err != nil {
			return pb.Money{}, err
		}
	}
	return product, nil
}