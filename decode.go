// Copyright 2023 OpenEsim. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Author: Johannes Waigel

package asn1go

func Unmarshal(data []byte, v interface{}) error {
	var ds decodeState
	err := checkValid(data, &ds.scan)
	if err != nil {
		return err
	}
	/*
		ds.init(data)
		return ds.unmarshal(v)*/
	return nil
}

// decodeState represents the state while decoding a ASN.1 value.
type decodeState struct {
	scan scanner
}
