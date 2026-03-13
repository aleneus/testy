# CHANGELOG

## 0.10.1

* fix: floats: NaN

## 0.10.0

* Added `DecodeHex()` for creating well-documented byte slices from hex strings (supports spaces and multi-line input).

## 0.9.0

* + `AssertRelEqualFloat32()`, `AssertRelNotEqualFloat32()` for relative comparison of float32
* + `AssertRelEqualFloat64()`, `AssertRelNotEqualFloat64()` for relative comparison of float64
* fix: `...NotEqual...`: message

## 0.8.0

* + `AssertEqualComplex()`

## 0.7.3

* refactor: floats comparison

## 0.7.2

* refactor: + `t.Helper()`

## 0.7.1

* fix: floats: NaN

## 0.7.0

* + `AssertErrorIs()`

## 0.6.0

* + `AssertEqualFloat32()`
* + `AssertEqualFloat64()`
* + `AssertNotEqualFloat32()`
* + `AssertNotEqualFloat64()`

## 0.5.0

* + `Writer` structure
* + `AssertNotSubstr()`
* - log stuff

## 0.4.0

* + `AssertPanic()`
* + `AssertSubstr()`

## 0.3.0

* + `Assert()`
* + `AssertNotNil()`

## 0.2.0

* Helpers for working with files in tests

## 0.1

* A couple of assert functions to reduce the amount of code in tests
