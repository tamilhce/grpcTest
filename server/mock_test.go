/*
Title : Mock Test GO GRPC SERVER
Description: Mock Test for GRPC SERVER to valdiate the Compute Service
Author: TAMILHCE
version : 1
*/
package main

import (
	"context"
	"testing"

	"github.com/tamilhce/grpcTest/proto"
)

func TestComputeService(t *testing.T) {

	s := server{}
	// Mock Data Sets
	tests := []struct {
		a              int64
		b              int64
		addResult      int64
		multiplyResult int64
	}{
		{
			a:              10,
			b:              20,
			addResult:      30,
			multiplyResult: 200,
		},
		{
			a:              10,
			b:              0,
			addResult:      10,
			multiplyResult: 0,
		},
	}

	for _, tt := range tests {
		req := &proto.Request{A: tt.a, B: tt.b}

		// Testing the ComputeAdd func
		addResp, addErr := s.ComputeAdd(context.Background(), req)
		if addErr != nil {
			t.Errorf("ComputeAdd got unexpected error")
		}
		if addResp.Result != tt.addResult {
			t.Errorf("ComputeADD (%v)+(%v)=%v, wanted %v", tt.a, tt.b, addResp.Result, tt.addResult)
		} else {
			t.Logf("ComputeADD (%v)+(%v)=%v, wanted %v", tt.a, tt.b, addResp.Result, tt.addResult)
		}
		// Testing the ComputeMulitply func
		mulitResp, mulitErr := s.ComputeMultiply(context.Background(), req)
		if mulitErr != nil {
			t.Errorf("ComputeMultiply got unexpected error")
		}
		if mulitResp.Result != tt.multiplyResult {
			t.Errorf("ComputeMultiply (%v)+(%v)=%v, wanted %v", tt.a, tt.b, mulitResp.Result, tt.multiplyResult)
		} else {
			t.Logf("ComputeMultiply (%v)+(%v)=%v, wanted %v", tt.a, tt.b, mulitResp.Result, tt.multiplyResult)
		}
	}
}
