// Package testy implements some helpers for writing tests.
package testy

import (
	"log"
)

// LogAssertNoErr compares err with nil.
func LogAssertNoErr(err error, msg ...interface{}) (ok bool) {
	if err != nil {
		log.Println("ERROR: ", msg, ":", err)

		return false
	}

	return true
}

// LogAssertEqual compares two values.
func LogAssertEqual(v1 interface{}, v2 interface{}, msg ...interface{}) (ok bool) {
	if v1 != v2 {
		log.Println("ERROR: ", msg)

		return false
	}

	return true
}

// LogAssertNotEqual compares two values.
func LogAssertNotEqual(v1 interface{}, v2 interface{}, msg ...interface{}) (ok bool) {
	if v1 == v2 {
		log.Println("ERROR: ", msg)

		return false
	}

	return true
}

// LogAssertTrue compares v with true.
func LogAssertTrue(v bool, msg ...interface{}) bool {
	if !v {
		log.Println("ERROR: ", msg)

		return false
	}

	return true
}

// LogErr just writes the error message to log.
func LogErr(msg ...interface{}) {
	log.Println("ERROR: ", msg)
}

// LogWarn just writes the warning message to log.
func LogWarn(msg ...interface{}) {
	log.Println("WARNING: ", msg)
}
