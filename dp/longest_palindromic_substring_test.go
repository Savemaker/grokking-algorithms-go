package main

import "testing"

const (
	TEST_ONE   = "abobaaaaaabbaaaaaa"
	TEST_TWO   = "aa"
	TEST_THREE = "abcde"
	TEST_FOUR  = "kayak"
	TEST_FIVE  = "lolkayak"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Run(TEST_ONE, func(t *testing.T) {
		useTestCase(t, TEST_ONE, "aaaaaabbaaaaaa")
	})

	t.Run(TEST_TWO, func(t *testing.T) {
		useTestCase(t, TEST_TWO, "aa")
	})

	t.Run(TEST_THREE, func(t *testing.T) {
		useTestCase(t, TEST_THREE, "a")
	})

	t.Run(TEST_FOUR, func(t *testing.T) {
		useTestCase(t, TEST_FOUR, "kayak")
	})
	t.Run(TEST_FIVE, func(t *testing.T) {
		useTestCase(t, TEST_FIVE, "kayak")
	})
}

func useTestCase(t *testing.T, testCase string, expectedString string) {
	t.Helper()
	res := GetLongestPalindrome(testCase)
	if res != expectedString {
		t.Errorf("Failed for %v: expected %v but got %v", testCase, expectedString, res)
	}
}
