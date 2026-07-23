package mathx

import (
	"math"
	"testing"
)

func TestGCD(t *testing.T) {
	if got := GCD(12, 18); got != 6 {
		t.Errorf("got %d", got)
	}
	if got := GCD(17, 5); got != 1 {
		t.Errorf("got %d", got)
	}
}

func TestLCM(t *testing.T) {
	if got := LCM(4, 6); got != 12 {
		t.Errorf("got %d", got)
	}
}

func TestCountPrimes(t *testing.T) {
	if got := CountPrimes(10); got != 4 { // 2,3,5,7
		t.Errorf("got %d", got)
	}
	if got := CountPrimes(0); got != 0 {
		t.Errorf("got %d", got)
	}
	if got := CountPrimes(2); got != 0 {
		t.Errorf("got %d", got)
	}
}

func TestMyPow(t *testing.T) {
	if got := MyPow(2, 10); got != 1024 {
		t.Errorf("got %v", got)
	}
	if got := MyPow(2, -2); math.Abs(got-0.25) > 1e-9 {
		t.Errorf("got %v", got)
	}
}

func TestReverseInteger(t *testing.T) {
	tests := []struct{ in, want int }{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0}, // overflow
	}
	for _, tt := range tests {
		if got := ReverseInteger(tt.in); got != tt.want {
			t.Errorf("ReverseInteger(%d)=%d want %d", tt.in, got, tt.want)
		}
	}
}

func TestIsPalindromeNumber(t *testing.T) {
	if !IsPalindromeNumber(121) {
		t.Error("121 is palindrome")
	}
	if IsPalindromeNumber(-121) {
		t.Error("-121 is not")
	}
	if IsPalindromeNumber(10) {
		t.Error("10 is not")
	}
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{{"III", 3}, {"LVIII", 58}, {"MCMXCIV", 1994}, {"IV", 4}}
	for _, tt := range tests {
		if got := RomanToInt(tt.in); got != tt.want {
			t.Errorf("RomanToInt(%q)=%d want %d", tt.in, got, tt.want)
		}
	}
}

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{{3, "III"}, {58, "LVIII"}, {1994, "MCMXCIV"}, {4, "IV"}}
	for _, tt := range tests {
		if got := IntToRoman(tt.in); got != tt.want {
			t.Errorf("IntToRoman(%d)=%q want %q", tt.in, got, tt.want)
		}
	}
}

func TestIsHappy(t *testing.T) {
	if !IsHappy(19) {
		t.Error("19 is happy")
	}
	if IsHappy(2) {
		t.Error("2 is not happy")
	}
}
