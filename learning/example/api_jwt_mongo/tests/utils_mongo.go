// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// src: https://github.com/mongodb/mongo-go-driver/blob/master/internal/testutil/config.go

package tests

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"go.mongodb.org/mongo-driver/x/mongo/driver/session"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
)

var connectionString connstring.ConnString
var connectionStringOnce sync.Once
var connectionStringErr error
var monitoredSessionPool *session.Pool

// ConnString gets the globally configured connection string.
func ConnString(t *testing.T) connstring.ConnString {
	connectionStringOnce.Do(func() {
		connectionString, connectionStringErr = GetConnString()
		mongodbURI := os.Getenv("MONGO_URI")
		if mongodbURI == "" {
			mongodbURI = MONGO_URI
		}

		mongodbURI = AddTLSConfigToURI(mongodbURI)
		mongodbURI = AddCompressorToUri(mongodbURI)

		var err error
		connectionString, err = connstring.Parse(mongodbURI)
		// connectionString, err = connstring.ParseAndValidate(mongodbURI)
		if err != nil {
			connectionStringErr = err
		}
	})
	if connectionStringErr != nil {
		t.Fatal(connectionStringErr)
	}

	return connectionString
}

func GetConnString() (connstring.ConnString, error) {
	mongodbURI := os.Getenv("MONGO_URI")
	if mongodbURI == "" {
		mongodbURI = MONGO_URI
	}

	mongodbURI = AddTLSConfigToURI(mongodbURI)

	// cs, err := connstring.ParseAndValidate(mongodbURI)
	cs, err := connstring.Parse(mongodbURI)
	if err != nil {
		return connstring.ConnString{}, err
	}

	return cs, nil
}

// DBName gets the globally configured database name.
func DBName(t *testing.T) string {
	return GetDBName(ConnString(t))
}

func GetDBName(cs connstring.ConnString) string {
	if cs.Database != "" {
		return cs.Database
	}

	// return fmt.Sprintf("mongo-go-driver-%d", os.Getpid())
	return MONGO_DBNAME
}

// ColName gets a collection name that should be unique
// to the currently executing test.
func ColName(t *testing.T) string {
	// Get this indirectly to avoid copying a mutex
	v := reflect.Indirect(reflect.ValueOf(t))
	name := v.FieldByName("name")
	return name.String()
}

// GlobalMonitoredSessionPool returns the globally configured session pool.
// Must be called after GlobalMonitoredTopology()
func GlobalMonitoredSessionPool() *session.Pool {
	return monitoredSessionPool
}

// AddOptionsToURI appends connection string options to a URI.
func AddOptionsToURI(uri string, opts ...string) string {
	if !strings.ContainsRune(uri, '?') {
		if uri[len(uri)-1] != '/' {
			uri += "/"
		}

		uri += "?"
	} else {
		uri += "&"
	}

	for _, opt := range opts {
		uri += opt
	}

	return uri
}

// AddTLSConfigToURI checks for the environmental variable indicating that the tests are being run
// on an SSL-enabled server, and if so, returns a new URI with the necessary configuration.
func AddTLSConfigToURI(uri string) string {
	caFile := os.Getenv("MONGO_GO_DRIVER_CA_FILE")
	if len(caFile) == 0 {
		return uri
	}

	return AddOptionsToURI(uri, "ssl=true&sslCertificateAuthorityFile=", caFile)
}

// AddCompressorToUri checks for the environment variable indicating that the tests are being run with compression
// enabled. If so, it returns a new URI with the necessary configuration
func AddCompressorToUri(uri string) string {
	comp := os.Getenv("MONGO_GO_DRIVER_COMPRESSOR")
	if len(comp) == 0 {
		return uri
	}

	return AddOptionsToURI(uri, "compressors=", comp)
}
