package aper

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// TEST BIT STRING
func TestSingleBitStringMarshal(t *testing.T) {
	runBitStringTest1Marshal(t)
	runBitStringTest2Marshal(t)
	runBitStringTest3Marshal(t)
	runBitStringTest4Marshal(t)
	runBitStringTest5Marshal(t)
	runBitStringTest6Marshal(t)
	runBitStringTest7Marshal(t)
	runBitStringTest8Marshal(t)
}

func TestSingleBitStringUnmarshal(t *testing.T) {
	runBitStringTest1Unmarshal(t)
	runBitStringTest2Unmarshal(t)
	runBitStringTest3Unmarshal(t)
	runBitStringTest4Unmarshal(t)
	runBitStringTest5Unmarshal(t)
	runBitStringTest6Unmarshal(t)
	runBitStringTest7Unmarshal(t)
	runBitStringTest8Unmarshal(t)
}

type bitStringTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled BitStringTest1
}

var bitStringTestData1 = []bitStringTestDataStruct1{
	{[]byte{0x17, 0xD4, 0xA5, 0x4A}, bitStringTest1Data[0]},
	{[]byte{0x18, 0xD4, 0xA5, 0x4B}, bitStringTest1Data[1]},
	{[]byte{0x20, 0x1F, 0xD4, 0xA5, 0x7F}, bitStringTest1Data[2]},
}

func runBitStringTest1Marshal(t *testing.T) {
	for i, test := range bitStringTestData1 {
		pd := NewPerBitData(nil)
		err := pd.WriteBitString(test.unmarshalled.BitString, false, nil, nil)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runBitStringTest1Unmarshal(t *testing.T) {
	for i, test := range bitStringTestData1 {
		pd := NewPerBitData(test.marshalled)
		bs, err := pd.ReadBitString(false, nil, nil)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(bs, test.unmarshalled.BitString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

type bitStringTestDataStruct2 struct {
	marshalled   []byte
	unmarshalled BitStringTest2
}

var bitStringTestData2 = []bitStringTestDataStruct2{
	{[]byte{0xAA, 0x56}, bitStringTest2Data[0]},
	{[]byte{0xAB, 0xd6}, bitStringTest2Data[1]},
}

func runBitStringTest2Marshal(t *testing.T) {
	for i, test := range bitStringTestData2 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 15, 15
		err := pd.WriteBitString(test.unmarshalled.BitString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runBitStringTest2Unmarshal(t *testing.T) {
	for i, test := range bitStringTestData2 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 15, 15
		bs, err := pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(bs, test.unmarshalled.BitString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

type bitStringTestDataStruct3 struct {
	marshalled   []byte
	unmarshalled BitStringTest3
}

var bitStringTestData3 = []bitStringTestDataStruct3{
	{[]byte{0xFF, 0x5E, 0xB0}, bitStringTest3Data[0]},
}

func runBitStringTest3Marshal(t *testing.T) {
	for i, test := range bitStringTestData3 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 20, 20
		err := pd.WriteBitString(test.unmarshalled.BitString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

func runBitStringTest3Unmarshal(t *testing.T) {
	for i, test := range bitStringTestData3 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 20, 20
		bs, err := pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(bs, test.unmarshalled.BitString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

type bitStringTestDataStruct4 struct {
	marshalled   []byte
	unmarshalled BitStringTest4
}

var bitStringTestData4 = []bitStringTestDataStruct4{
	{[]byte{0x17, 0x56, 0x2a, 0xdf}, bitStringTest4Data[0]},
	{[]byte{0x07, 0xd1}, bitStringTest4Data[1]},
}

func runBitStringTest4Marshal(t *testing.T) {
	for i, test := range bitStringTestData4 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 1, 160
		err := pd.WriteBitString(test.unmarshalled.BitString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 4 is FAILED", i+1)
	}
}

func runBitStringTest4Unmarshal(t *testing.T) {
	for i, test := range bitStringTestData4 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 1, 160
		bs, err := pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(bs, test.unmarshalled.BitString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 4 is FAILED", i+1)
	}
}

type bitStringTestDataStruct5 struct {
	marshalled   []byte
	unmarshalled BitStringTest5
}

var bitStringTestData5 = []bitStringTestDataStruct5{
	{[]byte{0x14, 0xF5, 0x5E, 0xB0}, bitStringTest5Data[0]},
	{[]byte{0x1D, 0xF5, 0x5F, 0xFF, 0x58}, bitStringTest5Data[1]},
}

func runBitStringTest5Marshal(t *testing.T) {
	for i, test := range bitStringTestData5 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 0, 255
		err := pd.WriteBitString(test.unmarshalled.BitString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 5 is FAILED", i+1)
	}
}

func runBitStringTest5Unmarshal(t *testing.T) {
	for i, test := range bitStringTestData5 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 0, 255
		bs, err := pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(bs, test.unmarshalled.BitString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 5 is FAILED", i+1)
	}
}

type bitStringTestDataStruct6 struct {
	marshalled   []byte
	unmarshalled BitStringTest6
}

var bitStringTestData6 = []bitStringTestDataStruct6{
	{[]byte{0x00, 0x14, 0xFF, 0x5E, 0xB0}, bitStringTest6Data[0]},
	{[]byte{0x00, 0x1D, 0xF5, 0x5F, 0xFF, 0x58}, bitStringTest6Data[1]},
	{[]byte{0x00, 0x18, 0xD4, 0xA5, 0x4B}, bitStringTest6Data[2]},
	{[]byte{0x00, 0x07, 0xB2}, bitStringTest6Data[3]},
}

func runBitStringTest6Marshal(t *testing.T) {
	for i, test := range bitStringTestData6 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 0, 355
		err := pd.WriteBitString(test.unmarshalled.BitString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 6 is FAILED", i+1)
	}
}

func runBitStringTest6Unmarshal(t *testing.T) {
	for i, test := range bitStringTestData6 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 0, 355
		bs, err := pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(bs, test.unmarshalled.BitString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 6 is FAILED", i+1)
	}
}

type bitStringTestDataStruct7 struct {
	marshalled   []byte
	unmarshalled BitStringTest7
}

var bitStringTestData7 = []bitStringTestDataStruct7{
	{[]byte{0x17, 0xD4, 0xA5, 0x4A}, bitStringTest7Data[0]},
	{[]byte{0x18, 0xD4, 0xA5, 0x4B}, bitStringTest7Data[1]},
	{[]byte{0x20, 0x1F, 0xD4, 0xA5, 0x7F}, bitStringTest7Data[2]},
	{[]byte(bigData), bitStringTest7Data[3]},
}

func runBitStringTest7Marshal(t *testing.T) {
	for i, test := range bitStringTestData7 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 0, 333333
		err := pd.WriteBitString(test.unmarshalled.BitString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 7 is FAILED", i+1)
	}
}

func runBitStringTest7Unmarshal(t *testing.T) {
	for i, test := range bitStringTestData7 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 0, 333333
		bs, err := pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(bs, test.unmarshalled.BitString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 7 is FAILED", i+1)
	}
}

type bitStringTestDataStruct8 struct {
	marshalled   []byte
	unmarshalled BitStringTest8
}

var bitStringTestData8 = []bitStringTestDataStruct8{
	{[]byte{0x80, 0x18, 0x12, 0x3A, 0xAA}, bitStringTest8Data[0]},
	{[]byte{0x40}, bitStringTest8Data[1]},
}

func runBitStringTest8Marshal(t *testing.T) {
	for i, test := range bitStringTestData8 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 1, 1
		err := pd.WriteBitString(test.unmarshalled.BitString, true, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 8 is FAILED", i+1)
	}
}

func runBitStringTest8Unmarshal(t *testing.T) {
	for i, test := range bitStringTestData8 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 1, 1
		bs, err := pd.ReadBitString(true, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(bs, test.unmarshalled.BitString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 8 is FAILED", i+1)
	}
}

// TEST STRUCT BIT STRING STRUCT
func TestStructBitStringMarshal(t *testing.T) {
	runStructBitStringTest1Marshal(t)
	runStructBitStringTest2Marshal(t)
	runStructBitStringTest3Marshal(t)
}

func TestStructBitStringUnmarshal(t *testing.T) {
	runStructBitStringTest1Unmarshal(t)
	runStructBitStringTest2Unmarshal(t)
	runStructBitStringTest3Unmarshal(t)
}

type structBitStringTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled BitStringStructTest1
}

var structBitStringTestData1 = []structBitStringTestDataStruct1{
	{[]byte{0xB4}, BitStringStructTest1Data[0]},
}

func runStructBitStringTest1Marshal(t *testing.T) {
	for i, test := range structBitStringTestData1 {
		var err error
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 3, 3
		if err = pd.WriteBitString(test.unmarshalled.BitString1, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if err = pd.WriteBitString(test.unmarshalled.BitString2, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}

		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runStructBitStringTest1Unmarshal(t *testing.T) {
	for i, test := range structBitStringTestData1 {
		var result BitStringStructTest1
		var err error
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 3, 3
		result.BitString1, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		result.BitString2, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(result, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

type structBitStringTestDataStruct2 struct {
	marshalled   []byte
	unmarshalled BitStringStructTest2
}

var structBitStringTestData2 = []structBitStringTestDataStruct2{
	{[]byte{0xB6}, BitStringStructTest2Data[0]},
}

func runStructBitStringTest2Marshal(t *testing.T) {
	for i, test := range structBitStringTestData2 {
		var err error
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 3, 3
		if err = pd.WriteBitString(test.unmarshalled.BitString1, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 4, 4
		if err = pd.WriteBitString(test.unmarshalled.BitString2, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}

		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runStructBitStringTest2Unmarshal(t *testing.T) {
	for i, test := range structBitStringTestData2 {
		var result BitStringStructTest2
		var err error
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 3, 3
		result.BitString1, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 4, 4
		result.BitString2, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(result, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

type structBitStringTestDataStruct3 struct {
	marshalled   []byte
	unmarshalled BitStringStructTest3
}

var structBitStringTestData3 = []structBitStringTestDataStruct3{
	{[]byte{0xA2, 0x00, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8}, BitStringStructTest3Data[0]},
}

func runStructBitStringTest3Marshal(t *testing.T) {
	var ans [][]byte = [][]byte{
		{0xA2, 0x00, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8},
	}
	for i, test := range structBitStringTestData3 {
		var err error
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 3, 3
		if err = pd.WriteBitString(test.unmarshalled.BitString1, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 125
		if err = pd.WriteBitString(test.unmarshalled.BitString2, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		if err = pd.WriteBitString(test.unmarshalled.BitString3, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 555
		if err = pd.WriteBitString(test.unmarshalled.BitString4, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}

		if bytes.Equal(pd.Bytes(), ans[i]) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

func runStructBitStringTest3Unmarshal(t *testing.T) {
	for i, test := range structBitStringTestData3 {
		var result BitStringStructTest3
		var err error
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 3, 3
		result.BitString1, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 125
		result.BitString2, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		result.BitString3, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 555
		result.BitString4, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}

		if reflect.DeepEqual(result, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

// TEST OCTET STRING
func TestSingleOctetStringMarshal(t *testing.T) {
	runSingleOctetStringTest1Marshal(t)
	runSingleOctetStringTest2Marshal(t)
	runSingleOctetStringTest3Marshal(t)
	runSingleOctetStringTest4Marshal(t)
	runSingleOctetStringTest5Marshal(t)
	runSingleOctetStringTest6Marshal(t)
	runSingleOctetStringTest7Marshal(t)
	runSingleOctetStringTest8Marshal(t)
}

func TestSingleOctetStringUnmarshal(t *testing.T) {
	runSingleOctetStringTest1Unmarshal(t)
	runSingleOctetStringTest2Unmarshal(t)
	runSingleOctetStringTest3Unmarshal(t)
	runSingleOctetStringTest4Unmarshal(t)
	runSingleOctetStringTest5Unmarshal(t)
	runSingleOctetStringTest6Unmarshal(t)
	runSingleOctetStringTest7Unmarshal(t)
	runSingleOctetStringTest8Unmarshal(t)
}

type octetStringTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled oCTETStringTest1
}

var octetStringTestData1 = []octetStringTestDataStruct1{
	{[]byte("\x07free5GC"), oCTETStringTest1Data[0]},
	{[]byte("\x04\x23\x34\x52\x97"), oCTETStringTest1Data[1]},
	{[]byte("\x08Jennifer"), oCTETStringTest1Data[2]},
}

func runSingleOctetStringTest1Marshal(t *testing.T) {
	for i, test := range octetStringTestData1 {
		pd := NewPerBitData(nil)
		err := pd.WriteOctetString(test.unmarshalled.OctetString, false, nil, nil)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d  of TEST 1 is FAILED", i+1)
	}
}

func runSingleOctetStringTest1Unmarshal(t *testing.T) {
	for i, test := range octetStringTestData1 {
		pd := NewPerBitData(test.marshalled)
		os, err := pd.ReadOctetString(false, nil, nil)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(os, test.unmarshalled.OctetString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d  of TEST 1 is FAILED", i+1)
	}
}

type octetStringTestDataStruct2 struct {
	marshalled   []byte
	unmarshalled oCTETStringTest2
}

var octetStringTestData2 = []octetStringTestDataStruct2{
	{[]byte("\xaa\x56"), oCTETStringTest2Data[0]},
	{[]byte("\x43\x12"), oCTETStringTest2Data[1]},
}

func runSingleOctetStringTest2Marshal(t *testing.T) {
	for i, test := range octetStringTestData2 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 2, 2
		err := pd.WriteOctetString(test.unmarshalled.OctetString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runSingleOctetStringTest2Unmarshal(t *testing.T) {
	for i, test := range octetStringTestData2 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 2, 2
		os, err := pd.ReadOctetString(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(os, test.unmarshalled.OctetString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

type octetStringTestDataStruct3 struct {
	marshalled   []byte
	unmarshalled oCTETStringTest3
}

var octetStringTestData3 = []octetStringTestDataStruct3{
	{[]byte("LLpRB9oV8zOkfraw1Nf5"), oCTETStringTest3Data[0]},
}

func runSingleOctetStringTest3Marshal(t *testing.T) {
	for i, test := range octetStringTestData3 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 20, 20
		err := pd.WriteOctetString(test.unmarshalled.OctetString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

func runSingleOctetStringTest3Unmarshal(t *testing.T) {
	for i, test := range octetStringTestData3 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 20, 20
		os, err := pd.ReadOctetString(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(os, test.unmarshalled.OctetString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

type octetStringTestDataStruct4 struct {
	marshalled   []byte
	unmarshalled oCTETStringTest4
}

var octetStringTestData4 = []octetStringTestDataStruct4{
	{[]byte("\x13LLpRB9oV8zOkfraw1Nf5"), oCTETStringTest4Data[0]},
	{[]byte("O1yYPj2WH4Uzex3sU40P1Kq7SgDB2sz0Ksg7fA76zcI5pxVDWtkUrfPti95h7xkzWpAcLaU7fMBBIJ981"), oCTETStringTest4Data[1]},
}

func runSingleOctetStringTest4Marshal(t *testing.T) {
	for i, test := range octetStringTestData4 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 1, 160
		err := pd.WriteOctetString(test.unmarshalled.OctetString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 4 is FAILED", i+1)
	}
}

func runSingleOctetStringTest4Unmarshal(t *testing.T) {
	for i, test := range octetStringTestData4 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 1, 160
		os, err := pd.ReadOctetString(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(os, test.unmarshalled.OctetString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 4 is FAILED", i+1)
	}
}

type octetStringTestDataStruct5 struct {
	marshalled   []byte
	unmarshalled oCTETStringTest5
}

var octetStringTestData5 = []octetStringTestDataStruct5{
	{[]byte("\x14LLpRB9oV8zOkfraw1Nf5"), oCTETStringTest5Data[0]},
	{[]byte("\xFFcGUpp6MH*7@55mntftf$k@eVdd3k2-*dVbGt?BmdTvTvs#ee9cktn6uA5u2g@cvE955P4rUqReG$Ybd83YY" +
		"?r5DqTYqrwDtHzeX+tFVK5RkBmns3GFhU9rPtX-eRfh62+Mmdeav2UFRy$wNghwSm?8RpeqBZTe8W-3Yfm#n=N" +
		"R..r@z6BRXGAX.DMz34ad@-N8Xy-V9AkC-6kPU*Yh$MW7+m-$B6e32!WCCeFe?d-QyV+@z#vKy6meZN87bV2hd"), oCTETStringTest5Data[1]},
}

func runSingleOctetStringTest5Marshal(t *testing.T) {
	for i, test := range octetStringTestData5 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 0, 255
		err := pd.WriteOctetString(test.unmarshalled.OctetString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 5 is FAILED", i+1)
	}
}

func runSingleOctetStringTest5Unmarshal(t *testing.T) {
	for i, test := range octetStringTestData5 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 0, 255
		os, err := pd.ReadOctetString(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(os, test.unmarshalled.OctetString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 5 is FAILED", i+1)
	}
}

type octetStringTestDataStruct6 struct {
	marshalled   []byte
	unmarshalled oCTETStringTest6
}

var octetStringTestData6 = []octetStringTestDataStruct6{
	{[]byte("\x00\x0EI!nGUXiqNpCP&a"), oCTETStringTest6Data[0]},
	{[]byte("\x00\x13u^YlZwgYxf7swQqweqw"), oCTETStringTest6Data[1]},
	{[]byte("\x00\x12iClFlb&YgrS4basdas"), oCTETStringTest6Data[2]},
	{[]byte("\x00\x07wirelab"), oCTETStringTest6Data[3]},
}

func runSingleOctetStringTest6Marshal(t *testing.T) {
	for i, test := range octetStringTestData6 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 0, 355
		err := pd.WriteOctetString(test.unmarshalled.OctetString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 6 is FAILED", i+1)
	}
}

func runSingleOctetStringTest6Unmarshal(t *testing.T) {
	for i, test := range octetStringTestData6 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 0, 355
		os, err := pd.ReadOctetString(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(os, test.unmarshalled.OctetString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 6 is FAILED", i+1)
	}
}

type octetStringTestDataStruct7 struct {
	marshalled   []byte
	unmarshalled oCTETStringTest7
}

var octetStringTestData7 = []octetStringTestDataStruct7{
	{[]byte("\x30\x80"), oCTETStringTest7Data[0]},
	{[]byte("\x80\x07free5GC"), oCTETStringTest7Data[1]},
}

func runSingleOctetStringTest7Marshal(t *testing.T) {
	for i, test := range octetStringTestData7 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 1, 1
		err := pd.WriteOctetString(test.unmarshalled.OctetString, true, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 7 is FAILED", i+1)
	}
}

func runSingleOctetStringTest7Unmarshal(t *testing.T) {
	for i, test := range octetStringTestData7 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 1, 1
		os, err := pd.ReadOctetString(true, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(os, test.unmarshalled.OctetString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 7 is FAILED", i+1)
	}
}

type octetStringTestDataStruct8 struct {
	marshalled   []byte
	unmarshalled oCTETStringTest8
}

var octetStringTestData8 = []octetStringTestDataStruct8{
	{[]byte("\x0EI!nGUXiqNpCP&a"), oCTETStringTest8Data[0]},
	{[]byte(bigOctetData), oCTETStringTest8Data[1]},
}

func runSingleOctetStringTest8Marshal(t *testing.T) {
	for i, test := range octetStringTestData8 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 0, 343434
		err := pd.WriteOctetString(test.unmarshalled.OctetString, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 8 is FAILED", i+1)
	}
}

func runSingleOctetStringTest8Unmarshal(t *testing.T) {
	for i, test := range octetStringTestData8 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 0, 343434
		os, err := pd.ReadOctetString(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(os, test.unmarshalled.OctetString) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 8 is FAILED", i+1)
	}
}

// TEST STRUCT OCTET STRING
func TestStructOctetStringMarshal(t *testing.T) {
	runStructOctetStringTest1Marshal(t)
	runStructOctetStringTest2Marshal(t)
}

func TestStructOctetStringUnmarshal(t *testing.T) {
	runStructOctetStringTest1Unmarshal(t)
	runStructOctetStringTest2Unmarshal(t)
}

type structOctetStringTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled oCTETStringStructTest1
}

var structOctetStringTestData1 = []structOctetStringTestDataStruct1{
	{[]byte("\x30\x8Cbcdef"), oCTETStringStructTest1Data[0]},
	{[]byte("\x30\xA0abcdefghij"), oCTETStringStructTest1Data[1]},
	{[]byte("\x30\x81\x89\x8C"), oCTETStringStructTest1Data[2]},
}

func runStructOctetStringTest1Marshal(t *testing.T) {
	for i, test := range structOctetStringTestData1 {
		pd := NewPerBitData(nil)
		var err error
		var lb, ub uint64 = 1, 1
		if err = pd.WriteOctetString(test.unmarshalled.OctetString1, true, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 20
		if err = pd.WriteOctetString(test.unmarshalled.OctetString2, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 2, 2
		if err = pd.WriteOctetString(test.unmarshalled.OctetString3, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runStructOctetStringTest1Unmarshal(t *testing.T) {
	for i, test := range structOctetStringTestData1 {
		pd := NewPerBitData(test.marshalled)
		var err error
		var result oCTETStringStructTest1
		var lb, ub uint64 = 1, 1
		result.OctetString1, err = pd.ReadOctetString(true, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 20
		result.OctetString2, err = pd.ReadOctetString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 2, 2
		result.OctetString3, err = pd.ReadOctetString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(result, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

type structOctetStringTestDataStruct2 struct {
	marshalled   []byte
	unmarshalled oCTETStringStructTest2
}

var structOctetStringTestData2 = []structOctetStringTestDataStruct2{
	{[]byte("\x30\xB1\x31\x88de"), oCTETStringStructTest2Data[0]},
	{[]byte("\x30\x99\x9A\x105678"), oCTETStringStructTest2Data[1]},
	{[]byte("\x30\x98\x99\x00"), oCTETStringStructTest2Data[2]},
}

func runStructOctetStringTest2Marshal(t *testing.T) {
	for i, test := range structOctetStringTestData2 {
		pd := NewPerBitData(nil)
		var err error
		var lb, ub uint64 = 1, 1
		if err = pd.WriteOctetString(test.unmarshalled.OctetString1, true, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 2, 2
		if err = pd.WriteOctetString(test.unmarshalled.OctetString3, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 20
		if err = pd.WriteOctetString(test.unmarshalled.OctetString2, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

func runStructOctetStringTest2Unmarshal(t *testing.T) {
	for i, test := range structOctetStringTestData2 {
		pd := NewPerBitData(test.marshalled)
		var result oCTETStringStructTest2
		var err error
		var lb, ub uint64 = 1, 1
		result.OctetString1, err = pd.ReadOctetString(true, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 2, 2
		result.OctetString3, err = pd.ReadOctetString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 20
		result.OctetString2, err = pd.ReadOctetString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(result, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

// TEST INTEGER
func TestIntegerMarshal(t *testing.T) {
	runIntegerTest1Marshal(t)
	runIntegerTest2Marshal(t)
	runIntegerTest3Marshal(t)
	runIntegerTest4Marshal(t)
	runIntegerTest5Marshal(t)
	runIntegerTest6Marshal(t)
	runIntegerTest7Marshal(t)
}

func TestIntegerUnmarshal(t *testing.T) {
	runIntegerTest1Unmarshal(t)
	runIntegerTest2Unmarshal(t)
	runIntegerTest3Unmarshal(t)
	runIntegerTest4Unmarshal(t)
	runIntegerTest5Unmarshal(t)
	runIntegerTest6Unmarshal(t)
	runIntegerTest7Unmarshal(t)
}

// 0x00 is an invalid input for unconstrained int,
// but we allow it and decode it as 0 for flexibility
func TestZeroByteIntegerUnmarshal(t *testing.T) {
	pd := NewPerBitData([]byte{0x00})
	val, err := pd.ReadInteger(false, nil, nil)
	require.NoError(t, err)
	require.Equal(t, int64(0), val)
}

type integerTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled intTest1
}

var integerTestData1 = []integerTestDataStruct1{
	{[]byte{0x01, 0x03}, intTest1Data[0]},
	{[]byte{0x03, 0x05, 0x16, 0x15}, intTest1Data[1]},
	{[]byte{0x03, 0xFA, 0xE9, 0xEB}, intTest1Data[2]},
}

func runIntegerTest1Marshal(t *testing.T) {
	for i, test := range integerTestData1 {
		pd := NewPerBitData(nil)
		err := pd.WriteInteger(test.unmarshalled.Value, false, nil, nil)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runIntegerTest1Unmarshal(t *testing.T) {
	for i, test := range integerTestData1 {
		pd := NewPerBitData(test.marshalled)
		val, err := pd.ReadInteger(false, nil, nil)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

type integerTestDataStruct2 struct {
	marshalled   []byte
	unmarshalled intTest2
}

var integerTestData2 = []integerTestDataStruct2{
	{[]byte{0x00}, intTest2Data[0]},
}

func runIntegerTest2Marshal(t *testing.T) {
	for i, test := range integerTestData2 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 3, 3
		err := pd.WriteInteger(test.unmarshalled.Value, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runIntegerTest2Unmarshal(t *testing.T) {
	for i, test := range integerTestData2 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub int64 = 3, 3
		val, err := pd.ReadInteger(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

type integerTestDataStruct3 struct {
	marshalled   []byte
	unmarshalled intTest3
}

var integerTestData3 = []integerTestDataStruct3{
	{[]byte{0x14}, intTest3Data[0]},
	{[]byte{0x18}, intTest3Data[1]},
	{[]byte{0x36}, intTest3Data[2]},
	{[]byte{0xDA}, intTest3Data[3]},
}

func runIntegerTest3Marshal(t *testing.T) {
	for i, test := range integerTestData3 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 1, 110
		err := pd.WriteInteger(test.unmarshalled.Value, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

func runIntegerTest3Unmarshal(t *testing.T) {
	for i, test := range integerTestData3 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub int64 = 1, 110
		val, err := pd.ReadInteger(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

type integerTestDataStruct4 struct {
	marshalled   []byte
	unmarshalled intTest4
}

var integerTestData4 = []integerTestDataStruct4{
	{[]byte{0x8C}, intTest4Data[0]},
}

func runIntegerTest4Marshal(t *testing.T) {
	for i, test := range integerTestData4 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 0, 255
		err := pd.WriteInteger(test.unmarshalled.Value, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 4 is FAILED", i+1)
	}
}

func runIntegerTest4Unmarshal(t *testing.T) {
	for i, test := range integerTestData4 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub int64 = 0, 255
		val, err := pd.ReadInteger(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 4 is FAILED", i+1)
	}
}

type integerTestDataStruct5 struct {
	marshalled   []byte
	unmarshalled intTest5
}

var integerTestData5 = []integerTestDataStruct5{
	{[]byte{0x00, 0x8C}, intTest5Data[0]},
}

func runIntegerTest5Marshal(t *testing.T) {
	for i, test := range integerTestData5 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 0, 65535
		err := pd.WriteInteger(test.unmarshalled.Value, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 5 is FAILED", i+1)
	}
}

func runIntegerTest5Unmarshal(t *testing.T) {
	for i, test := range integerTestData5 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub int64 = 0, 65535
		val, err := pd.ReadInteger(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 5 is FAILED", i+1)
	}
}

type integerTestDataStruct6 struct {
	marshalled   []byte
	unmarshalled intTest6
}

var integerTestData6 = []integerTestDataStruct6{
	{[]byte{0x00, 0x8C}, intTest6Data[0]},
	{[]byte{0xC0, 0xFF, 0xFF, 0xFF, 0xFF}, intTest6Data[1]},
	{[]byte{0x40, 0xFF, 0xFF}, intTest6Data[2]},
	{[]byte{0x80, 0x01, 0x00, 0x00}, intTest6Data[3]},
}

func runIntegerTest6Marshal(t *testing.T) {
	for i, test := range integerTestData6 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 0, 4294967295
		err := pd.WriteInteger(test.unmarshalled.Value, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 6 is FAILED", i+1)
	}
}

func runIntegerTest6Unmarshal(t *testing.T) {
	for i, test := range integerTestData6 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub int64 = 0, 4294967295
		val, err := pd.ReadInteger(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 6 is FAILED", i+1)
	}
}

type integerTestDataStruct7 struct {
	marshalled   []byte
	unmarshalled intTest7
}

var integerTestData7 = []integerTestDataStruct7{
	{[]byte{0x80, 0x02, 0x00, 0x8C}, intTest7Data[0]},
	{[]byte{0x80, 0x04, 0x7F, 0xFF, 0xFF, 0xFF}, intTest7Data[1]},
	{[]byte{0x80, 0x03, 0x00, 0xFF, 0xFF}, intTest7Data[2]},
	{[]byte{0x80, 0x03, 0x01, 0x00, 0x00}, intTest7Data[3]},
	{[]byte{0x42}, intTest7Data[4]},
}

func runIntegerTest7Marshal(t *testing.T) {
	for i, test := range integerTestData7 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 0, 45
		err := pd.WriteInteger(test.unmarshalled.Value, true, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 7 is FAILED", i+1)
	}
}

func runIntegerTest7Unmarshal(t *testing.T) {
	for i, test := range integerTestData7 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub int64 = 0, 45
		val, err := pd.ReadInteger(true, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 7 is FAILED", i+1)
	}
}

// unsigned integer (uint64)
func TestUint64IntegerMarshal(t *testing.T) {
	runUint64IntegerTest1Marshal(t)
	runUint64IntegerTest2Marshal(t)
}

func TestUint64IntegerUnmarshal(t *testing.T) {
	runUint64IntegerTest1Unmarshal(t)
	runUint64IntegerTest2Unmarshal(t)
}

// large ub
type uint64Test1 struct {
	Value uint64 //`aper:"valueLB:0,valueUB:18446744073709551615"`
}

var uint64Test1Data = []uint64Test1{
	{2147483647},
}

type uint64TestDataStruct8 struct {
	marshalled   []byte
	unmarshalled uint64Test1
}

var uint64TestData1 = []uint64TestDataStruct8{
	{[]byte{0x60, 0x7F, 0xFF, 0xFF, 0xFF}, uint64Test1Data[0]},
}

func runUint64IntegerTest1Marshal(t *testing.T) {
	for i, test := range uint64TestData1 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 0, 18446744073709551615
		err := pd.WriteUint64Integer(test.unmarshalled.Value, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runUint64IntegerTest1Unmarshal(t *testing.T) {
	for i, test := range uint64TestData1 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64 = 0, 18446744073709551615
		val, err := pd.ReadUint64Integer(false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

// no lb, no ub
type uint64Test2 struct {
	Value uint64
}

var uint64Test2Data = []uint64Test2{
	{2147483647},
}

type uint64TestDataStruct2 struct {
	marshalled   []byte
	unmarshalled uint64Test2
}

var uint64TestData2 = []uint64TestDataStruct2{
	{[]byte{0x04, 0x7F, 0xFF, 0xFF, 0xFF}, uint64Test2Data[0]},
}

func runUint64IntegerTest2Marshal(t *testing.T) {
	for i, test := range uint64TestData2 {
		pd := NewPerBitData(nil)
		err := pd.WriteUint64Integer(test.unmarshalled.Value, false, nil, nil)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runUint64IntegerTest2Unmarshal(t *testing.T) {
	for i, test := range uint64TestData2 {
		pd := NewPerBitData(test.marshalled)
		val, err := pd.ReadUint64Integer(false, nil, nil)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if val == test.unmarshalled.Value {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

// TEST STRUCT INTEGER
func TestStructIntegerMarshal(t *testing.T) {
	runStructIntegerTest1Marshal(t)
}

func TestStructIntegerUnmarshal(t *testing.T) {
	runStructIntegerTest1Unmarshal(t)
}

type structIntegerTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled intStructTest1
}

var structIntegerTestData1 = []structIntegerTestDataStruct1{
	{[]byte{0x58, 0x7B, 0x80, 0x02, 0x19, 0x2D}, intStructTest1Data[0]},
}

func runStructIntegerTest1Marshal(t *testing.T) {
	for i, test := range structIntegerTestData1 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 1, 110
		err := pd.WriteInteger(test.unmarshalled.Int1, false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		err = pd.WriteInteger(test.unmarshalled.Int2, false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 45
		err = pd.WriteInteger(test.unmarshalled.Int3, true, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runStructIntegerTest1Unmarshal(t *testing.T) {
	for i, test := range structIntegerTestData1 {
		var result intStructTest1
		var err error
		pd := NewPerBitData(test.marshalled)

		var lb, ub int64 = 1, 110
		result.Int1, err = pd.ReadInteger(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		result.Int2, err = pd.ReadInteger(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 45
		result.Int3, err = pd.ReadInteger(true, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

// TEST ENUMERATED
func TestSingleEnumMarshal(t *testing.T) {
	runSingleEnumTest1Marshal(t)
	runSingleEnumTest2Marshal(t)
}

func TestSingleEnumUnmarshal(t *testing.T) {
	runSingleEnumTest1Unmarshal(t)
	runSingleEnumTest2Unmarshal(t)
}

type enumTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled enumTest1
}

var enumTestData1 = []enumTestDataStruct1{
	{[]byte{0x00}, enumTest1Data[0]},
	{[]byte{0x40}, enumTest1Data[1]},
}

func runSingleEnumTest1Marshal(t *testing.T) {
	for i, test := range enumTestData1 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 0, 3
		err := pd.WriteEnumerated(test.unmarshalled.Value, false, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runSingleEnumTest1Unmarshal(t *testing.T) {
	for i, test := range enumTestData1 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub int64 = 0, 3
		enum, err := pd.ReadEnumerated(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(enum, test.unmarshalled.Value) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

type enumTestDataStruct2 struct {
	marshalled   []byte
	unmarshalled enumTest2
}

var enumTestData2 = []enumTestDataStruct2{
	{[]byte{0x10}, enumTest2Data[0]},
	{[]byte{0x20}, enumTest2Data[1]},
	{[]byte{0x80}, enumTest2Data[2]},
}

func runSingleEnumTest2Marshal(t *testing.T) {
	for i, test := range enumTestData2 {
		pd := NewPerBitData(nil)
		var lb, ub int64 = 0, 4
		err := pd.WriteEnumerated(test.unmarshalled.Value, true, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runSingleEnumTest2Unmarshal(t *testing.T) {
	for i, test := range enumTestData2 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub int64 = 0, 4
		enum, err := pd.ReadEnumerated(true, &lb, &ub)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(enum, test.unmarshalled.Value) {
			continue
		}
		fmt.Printf("%d", enum)
		fmt.Printf("%d", test.unmarshalled.Value)
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

// TEST SEQUENCE OF
func TestSingleSequenceOfMarshal(t *testing.T) {
	runSingleSequenceOfTest1Marshal(t)
	runSingleSequenceOfTest2Marshal(t)
	runSingleSequenceOfTest3Marshal(t)
	runSingleSequenceOfTest4Marshal(t)
	runSingleSequenceOfTest5Marshal(t)
	runSingleSequenceOfTest6Marshal(t)
	runSingleSequenceOfTest6Marshal(t)
	runSingleSequenceOfTest7Marshal(t)
}

func TestSingleSequenceOfUnmarshal(t *testing.T) {
	runSingleSequenceOfTest1Unmarshal(t)
	runSingleSequenceOfTest2Unmarshal(t)
	runSingleSequenceOfTest3Unmarshal(t)
	runSingleSequenceOfTest4Unmarshal(t)
	runSingleSequenceOfTest5Unmarshal(t)
	runSingleSequenceOfTest6Unmarshal(t)
	runSingleSequenceOfTest6Unmarshal(t)
	runSingleSequenceOfTest7Unmarshal(t)
}

type singleSeqOfTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled seqofTest1
}

var singleSeqOfTestData1 = []singleSeqOfTestDataStruct1{
	{[]byte{0xC0, 0x01, 0x03, 0x03, 0x05, 0x16, 0x15, 0x03, 0xFA, 0xE9, 0xEB}, seqofTest1Data[0]},
}

func runSingleSequenceOfTest1Marshal(t *testing.T) {
	for i, test := range singleSeqOfTestData1 {
		pd := NewPerBitData(nil)
		var lb, ub uint64 = 0, 3
		err := pd.WriteSequenceOfPreambleBitMap(uint64(len(test.unmarshalled.List)), false, &lb, &ub)
		for _, component := range test.unmarshalled.List {
			err = pd.WriteInteger(component.Value, false, nil, nil)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runSingleSequenceOfTest1Unmarshal(t *testing.T) {
	for i, test := range singleSeqOfTestData1 {
		pd := NewPerBitData(test.marshalled)
		li := []intTest1{}
		var lb, ub uint64 = 0, 3
		numComponents, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		for j := 0; j < int(numComponents); j++ {
			var val intTest1
			if val.Value, err = pd.ReadInteger(false, nil, nil); err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			} else {
				li = append(li, val)
			}
		}

		if reflect.DeepEqual(li, test.unmarshalled.List) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

type singleSeqOfTestDataStruct2 struct {
	marshalled   []byte
	unmarshalled seqofTest2
}

var singleSeqOfTestData2 = []singleSeqOfTestDataStruct2{
	{[]byte{0x0A, 0xC0, 0x7B, 0x80, 0x02, 0x19, 0x2D}, seqofTest2Data[0]},
}

func runSingleSequenceOfTest2Marshal(t *testing.T) {
	var sLb, sUb uint64
	var vLb, vUb int64
	for i, test := range singleSeqOfTestData2 {
		pd := NewPerBitData(nil)
		sLb, sUb = 0, 30
		err := pd.WriteSequenceOfPreambleBitMap(uint64(len(test.unmarshalled.List)), false, &sLb, &sUb)
		for _, component := range test.unmarshalled.List {
			vLb, vUb = 1, 110
			err = pd.WriteInteger(component.Int1, false, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			vLb, vUb = 0, 255
			err = pd.WriteInteger(component.Int2, false, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			vLb, vUb = 0, 45
			err = pd.WriteInteger(component.Int3, true, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runSingleSequenceOfTest2Unmarshal(t *testing.T) {
	var sLb, sUb uint64
	var vLb, vUb int64
	for i, test := range singleSeqOfTestData2 {
		pd := NewPerBitData(test.marshalled)
		li := []intStructTest1{}
		sLb, sUb = 0, 30
		numComponents, err := pd.ReadSequenceOfPreambleBitMap(false, &sLb, &sUb)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		for j := 0; j < int(numComponents); j++ {
			var val intStructTest1
			vLb, vUb = 1, 110
			val.Int1, err = pd.ReadInteger(false, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			vLb, vUb = 0, 255
			val.Int2, err = pd.ReadInteger(false, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			vLb, vUb = 0, 45
			val.Int3, err = pd.ReadInteger(true, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			li = append(li, val)
		}

		if reflect.DeepEqual(li, test.unmarshalled.List) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

type singleSeqOfTestDataStruct3 struct {
	marshalled   []byte
	unmarshalled seqofTest3
}

var singleSeqOfTestData3 = []singleSeqOfTestDataStruct3{
	{[]byte{0x06, 0x88, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8}, seqofTest3Data[0]},
}

func runSingleSequenceOfTest3Marshal(t *testing.T) {
	var lb, ub uint64
	for i, test := range singleSeqOfTestData3 {
		pd := NewPerBitData(nil)
		lb, ub = 0, 50
		err := pd.WriteSequenceOfPreambleBitMap(uint64(len(test.unmarshalled.List)), false, &lb, &ub)
		for _, component := range test.unmarshalled.List {
			lb, ub = 3, 3
			err = pd.WriteBitString(component.BitString1, false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			lb, ub = 0, 125
			err = pd.WriteBitString(component.BitString2, false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			lb, ub = 0, 255
			err = pd.WriteBitString(component.BitString3, false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			lb, ub = 0, 555
			err = pd.WriteBitString(component.BitString4, false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

func runSingleSequenceOfTest3Unmarshal(t *testing.T) {
	var lb, ub uint64
	for i, test := range singleSeqOfTestData3 {
		pd := NewPerBitData(test.marshalled)
		li := []BitStringStructTest3{}
		lb, ub = 0, 50
		numComponents, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		for j := 0; j < int(numComponents); j++ {
			var val BitStringStructTest3
			lb, ub = 3, 3
			val.BitString1, err = pd.ReadBitString(false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			lb, ub = 0, 125
			val.BitString2, err = pd.ReadBitString(false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			lb, ub = 0, 255
			val.BitString3, err = pd.ReadBitString(false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			lb, ub = 0, 555
			val.BitString4, err = pd.ReadBitString(false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}

			li = append(li, val)
		}

		if reflect.DeepEqual(li, test.unmarshalled.List) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

type singleSeqOfTestDataStruct4 struct {
	marshalled   []byte
	unmarshalled seqofTest4
}

var singleSeqOfTestData4 = []singleSeqOfTestDataStruct4{
	{[]byte("\x2C\x02\x00\x8C\x80\x04\x7F\xFF\xFF\xFF\x80\x03\x00\xFF\xFF\x80\x03\x01\x00\x00\x42"), seqofTest4Data[0]},
}

func runSingleSequenceOfTest4Marshal(t *testing.T) {
	var sLb, sUb uint64
	var vLb, vUb int64
	for i, test := range singleSeqOfTestData4 {
		pd := NewPerBitData(nil)
		sLb, sUb = 0, 16
		err := pd.WriteSequenceOfPreambleBitMap(uint64(len(test.unmarshalled.List)), false, &sLb, &sUb)
		for _, component := range test.unmarshalled.List {
			vLb, vUb = 0, 45
			err = pd.WriteInteger(component.Value, true, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 4 is FAILED", i+1)
	}
}

func runSingleSequenceOfTest4Unmarshal(t *testing.T) {
	var sLb, sUb uint64
	var vLb, vUb int64
	for i, test := range singleSeqOfTestData4 {
		pd := NewPerBitData(test.marshalled)
		li := []intTest7{}
		sLb, sUb = 0, 16
		numComponents, err := pd.ReadSequenceOfPreambleBitMap(false, &sLb, &sUb)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		for j := 0; j < int(numComponents); j++ {
			var val intTest7
			vLb, vUb = 0, 45
			val.Value, err = pd.ReadInteger(true, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			li = append(li, val)
		}

		if reflect.DeepEqual(li, test.unmarshalled.List) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 4 is FAILED", i+1)
	}
}

type singleSeqOfTestDataStruct5 struct {
	marshalled   []byte
	unmarshalled seqofTest5
}

var singleSeqOfTestData5 = []singleSeqOfTestDataStruct5{
	{[]byte{0x04, 0x14, 0x30, 0xDE, 0xD0}, seqofTest5Data[0]},
}

func runSingleSequenceOfTest5Marshal(t *testing.T) {
	var sLb, sUb uint64
	var vLb, vUb int64
	for i, test := range singleSeqOfTestData5 {
		pd := NewPerBitData(nil)
		sLb, sUb = 0, 255
		err := pd.WriteSequenceOfPreambleBitMap(uint64(len(test.unmarshalled.List)), false, &sLb, &sUb)
		for _, component := range test.unmarshalled.List {
			vLb, vUb = 1, 110
			err = pd.WriteInteger(component.Value, false, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 5 is FAILED", i+1)
	}
}

func runSingleSequenceOfTest5Unmarshal(t *testing.T) {
	var sLb, sUb uint64
	var vLb, vUb int64
	for i, test := range singleSeqOfTestData5 {
		pd := NewPerBitData(test.marshalled)
		li := []intTest3{}
		sLb, sUb = 0, 255
		numComponents, err := pd.ReadSequenceOfPreambleBitMap(false, &sLb, &sUb)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		for j := 0; j < int(numComponents); j++ {
			var val intTest3
			vLb, vUb = 1, 110
			val.Value, err = pd.ReadInteger(false, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			li = append(li, val)
		}

		if reflect.DeepEqual(li, test.unmarshalled.List) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 5 is FAILED", i+1)
	}
}

type singleSeqOfTestDataStruct6 struct {
	marshalled   []byte
	unmarshalled seqofTest6
}

var singleSeqOfTestData6 = []singleSeqOfTestDataStruct6{
	{[]byte{0x04, 0x14, 0x30, 0xDE, 0xD0}, seqofTest6Data[0]},
}

func runSingleSequenceOfTest6Marshal(t *testing.T) {
	var lb, ub int64
	for i, test := range singleSeqOfTestData6 {
		pd := NewPerBitData(nil)
		err := pd.WriteSequenceOfPreambleBitMap(uint64(len(test.unmarshalled.List)), false, nil, nil)
		for _, component := range test.unmarshalled.List {
			lb, ub = 1, 110
			err = pd.WriteInteger(component.Value, false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 6 is FAILED", i+1)
	}
}

func runSingleSequenceOfTest6Unmarshal(t *testing.T) {
	var lb, ub int64
	for i, test := range singleSeqOfTestData6 {
		pd := NewPerBitData(test.marshalled)
		li := []intTest3{}
		numComponents, err := pd.ReadSequenceOfPreambleBitMap(false, nil, nil)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		for j := 0; j < int(numComponents); j++ {
			var val intTest3
			lb, ub = 1, 110
			val.Value, err = pd.ReadInteger(false, &lb, &ub)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			li = append(li, val)
		}

		if reflect.DeepEqual(li, test.unmarshalled.List) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 6 is FAILED", i+1)
	}
}

type singleSeqOfTestDataStruct7 struct {
	marshalled   []byte
	unmarshalled seqofTest7
}

var singleSeqOfTestData7 = []singleSeqOfTestDataStruct7{
	{[]byte{0x14, 0x30, 0xDE, 0xD0}, seqofTest7Data[0]},
}

func runSingleSequenceOfTest7Marshal(t *testing.T) {
	var sLb, sUb uint64
	var vLb, vUb int64
	for i, test := range singleSeqOfTestData7 {
		pd := NewPerBitData(nil)
		sLb, sUb = 4, 4
		err := pd.WriteSequenceOfPreambleBitMap(uint64(len(test.unmarshalled.List)), false, &sLb, &sUb)
		for _, component := range test.unmarshalled.List {
			vLb, vUb = 1, 110
			err = pd.WriteInteger(component.Value, false, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 7 is FAILED", i+1)
	}
}

func runSingleSequenceOfTest7Unmarshal(t *testing.T) {
	var sLb, sUb uint64
	var vLb, vUb int64
	for i, test := range singleSeqOfTestData7 {
		pd := NewPerBitData(test.marshalled)
		li := []intTest3{}
		sLb, sUb = 4, 4
		numComponents, err := pd.ReadSequenceOfPreambleBitMap(false, &sLb, &sUb)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		for j := 0; j < int(numComponents); j++ {
			var val intTest3
			vLb, vUb = 1, 110
			val.Value, err = pd.ReadInteger(false, &vLb, &vUb)
			if err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			}
			li = append(li, val)
		}

		if reflect.DeepEqual(li, test.unmarshalled.List) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 7 is FAILED", i+1)
	}
}

// TEST CHOICE
func TestChoiceMarshal(t *testing.T) {
	runChoiceTest1Marshal(t)
	runChoiceTest2Marshal(t)
}

func TestChoiceUnmarshal(t *testing.T) {
	runChoiceTest1Unmarshal(t)
	runChoiceTest2Unmarshal(t)
}

// NEW DATA FORMAT
// TEST 1
type newChoiceStruct1 struct {
	Option _NewChoiceStruct1Alt `aper:"valueLB:0,valueUB:2"`
}

func (c *newChoiceStruct1) setOption(index int) error {
	switch index {
	case 0:
		c.Option = &test1Choice1{}
		return nil
	case 1:
		c.Option = &test1Choice2{}
		return nil
	case 2:
		c.Option = &test1Choice3{}
		return nil
	default:
		return fmt.Errorf("Invalid choice index")
	}
}

type _NewChoiceStruct1Alt interface {
	_NewChoiceStruct1Alt()
	Encode(*PerBitData) error
	Decode(*PerBitData) error
}

type test1Choice1 struct { // index = 0
	data []intTest1 `aper:"sizeLB:0,sizeUB:3"`
}

type test1Choice2 struct { // index = 1
	data []intStructTest1 `aper:"sizeLB:0,sizeUB:30"`
}

type test1Choice3 struct { // index = 2
	data []BitStringStructTest3 `aper:"sizeLB:0,sizeUB:50"`
}

// choice 1
func (c *test1Choice1) _NewChoiceStruct1Alt() {}

func (c *test1Choice1) Encode(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 3
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(c.data)), false, &lb, &ub)
	for i := 0; i < len(c.data); i++ {
		err = pd.WriteInteger(c.data[i].Value, false, nil, nil)
	}
	return err
}

func (c *test1Choice1) Decode(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 3
	l, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
	if err != nil {
		return err
	}
	for i := 0; i < int(l); i++ {
		var val intTest1
		val.Value, err = pd.ReadInteger(false, nil, nil)
		c.data = append(c.data, val)
	}
	return err
}

// choice 2
func (c *test1Choice2) _NewChoiceStruct1Alt() {}

func (c *test1Choice2) Encode(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 30
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(c.data)), false, &lb, &ub)
	for i := 0; i < len(c.data); i++ {
		var lb, ub int64
		lb, ub = 1, 110
		err = pd.WriteInteger(c.data[i].Int1, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		err = pd.WriteInteger(c.data[i].Int2, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 45
		err = pd.WriteInteger(c.data[i].Int3, true, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
	}
	return err
}

func (c *test1Choice2) Decode(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 30
	l, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
	for i := 0; i < int(l); i++ {
		var val intStructTest1
		var lb, ub int64
		lb, ub = 1, 110
		val.Int1, err = pd.ReadInteger(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		val.Int2, err = pd.ReadInteger(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 45
		val.Int3, err = pd.ReadInteger(true, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		c.data = append(c.data, val)
	}
	return err
}

// choice 3
func (c *test1Choice3) _NewChoiceStruct1Alt() {}

func (c *test1Choice3) Encode(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 50
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(c.data)), false, &lb, &ub)
	for i := 0; i < len(c.data); i++ {
		var lb, ub uint64
		lb, ub = 3, 3
		err = pd.WriteBitString(c.data[i].BitString1, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 125
		err = pd.WriteBitString(c.data[i].BitString2, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		err = pd.WriteBitString(c.data[i].BitString3, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 555
		err = pd.WriteBitString(c.data[i].BitString4, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
	}
	return err
}

func (c *test1Choice3) Decode(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 50
	l, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
	for i := 0; i < int(l); i++ {
		var val BitStringStructTest3
		var lb, ub uint64
		lb, ub = 3, 3
		val.BitString1, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 125
		val.BitString2, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		val.BitString3, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 555
		val.BitString4, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		c.data = append(c.data, val)
	}
	return err
}

var newChoiceTest1Data = []newChoiceStruct1{
	{&test1Choice1{intTest1Data}},
	{&test1Choice2{intStructTest1Data}},
	{&test1Choice3{BitStringStructTest3Data}},
}

type newChoiceTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled newChoiceStruct1
}

var newChoiceTestData1 = []newChoiceTestDataStruct1{
	{[]byte{0x30, 0x01, 0x03, 0x03, 0x05, 0x16, 0x15, 0x03, 0xFA, 0xE9, 0xEB}, newChoiceTest1Data[0]},
	{[]byte{0x42, 0xB0, 0x7B, 0x80, 0x02, 0x19, 0x2D}, newChoiceTest1Data[1]},
	{[]byte{0x81, 0xA2, 0x00, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8}, newChoiceTest1Data[2]},
}

func runChoiceTest1Marshal(t *testing.T) {
	var indexes []int = []int{
		0,
		1,
		2,
	}
	var pd *PerBitData
	var err error

	// Test 1
	for i, test := range newChoiceTestData1 {
		pd = NewPerBitData(nil)
		var choiceUb, index int64 = 2, int64(indexes[i])
		if err = pd.WriteChoicePreambleBitMap(index, false, &choiceUb); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		err = test.unmarshalled.Option.Encode(pd)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

func runChoiceTest1Unmarshal(t *testing.T) {
	var pd *PerBitData
	// Test 1
	for i, test := range newChoiceTestData1 {
		var val newChoiceStruct1
		pd = NewPerBitData(test.marshalled)
		var choiceUb int64 = 2
		index, err := pd.ReadChoicePreambleBitMap(false, &choiceUb)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		err = val.setOption(int(index))
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		err = val.Option.Decode(pd)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(val, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

// TEST 2
type newChoiceStruct2 struct {
	Option _NewChoiceStruct2Alt //`aper:"valueExt,valueLB:0,valueUB:2"`
}

func (c *newChoiceStruct2) setOption(index int) error {
	switch index {
	// case 0: []intTest1
	// case 1: []intStructTest1
	case 2:
		c.Option = &test2Choice{}
		return nil
	default:
		return fmt.Errorf("Invalid choice index")
	}
}

type _NewChoiceStruct2Alt interface {
	_NewChoiceStruct2Alt()
	Encode(*PerBitData) error
	Decode(*PerBitData) error
}

type test2Choice struct { // index = 2
	data []BitStringStructTest3 //`aper:"sizeLB:0,sizeUB:50"`
}

func (c *test2Choice) _NewChoiceStruct2Alt() {}

func (c *test2Choice) Encode(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 50
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(c.data)), false, &lb, &ub)
	for i := 0; i < len(c.data); i++ {
		var lb, ub uint64
		lb, ub = 3, 3
		err = pd.WriteBitString(c.data[i].BitString1, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 125
		err = pd.WriteBitString(c.data[i].BitString2, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		err = pd.WriteBitString(c.data[i].BitString3, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 555
		err = pd.WriteBitString(c.data[i].BitString4, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
	}
	return err
}

func (c *test2Choice) Decode(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 50
	l, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
	for i := 0; i < int(l); i++ {
		var val BitStringStructTest3
		var lb, ub uint64
		lb, ub = 3, 3
		val.BitString1, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 125
		val.BitString2, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		val.BitString3, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 555
		val.BitString4, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		c.data = append(c.data, val)
	}
	return err
}

var newChoiceTest2Data = []newChoiceStruct2{
	{&test2Choice{BitStringStructTest3Data}},
}

type newChoiceTestDataStruct2 struct {
	marshalled   []byte
	unmarshalled newChoiceStruct2
}

var newChoiceTestData2 = []newChoiceTestDataStruct2{
	{[]byte{0x40, 0xD1, 0x00, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8}, newChoiceTest2Data[0]},
}

func runChoiceTest2Marshal(t *testing.T) {
	var indexes []int = []int{
		2,
	}
	var pd *PerBitData
	var err error

	// Test 2
	for i, test := range newChoiceTestData2 {
		pd = NewPerBitData(nil)
		var choiceUb, index int64 = 2, int64(indexes[i])
		if err = pd.WriteChoicePreambleBitMap(index, true, &choiceUb); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		err = test.unmarshalled.Option.Encode(pd)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runChoiceTest2Unmarshal(t *testing.T) {
	var pd *PerBitData
	// Test 1
	for i, test := range newChoiceTestData2 {
		var val newChoiceStruct2
		pd = NewPerBitData(test.marshalled)
		var choiceUb int64 = 2
		index, err := pd.ReadChoicePreambleBitMap(true, &choiceUb)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		err = val.setOption(int(index))
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		err = val.Option.Decode(pd)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(val, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

// TEST String: PrintableString, VisibleString, UTF8String

// PrintableString
type printableStringTestDataStruct struct {
	marshalled   []byte
	unmarshalled PrintableString
}

// unconstrained
var printableStringTestData1 = []printableStringTestDataStruct{
	{[]byte{0x00}, PrintableString("")},
	{[]byte{0x05, 0x61, 0x62, 0x63, 0x64, 0x65}, PrintableString("abcde")},
}

// sizeExt, lb=1, ub=1
var printableStringTestData2 = []printableStringTestDataStruct{
	{[]byte{0x30, 0x80}, PrintableString("a")},
	{[]byte{0x80, 0x02, 0x61, 0x61}, PrintableString("aa")},
}

// lb=0, ub=20
var printableStringTestData3 = []printableStringTestDataStruct{
	{[]byte{0x00}, PrintableString("")},
	{[]byte{0x20, 0x61, 0x61, 0x62, 0x62}, PrintableString("aabb")},
	{[]byte{
		0xA0, 0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62,
		0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62,
	}, PrintableString("aabbaabbaabbaabbaabb")},
}

func TestPrintableStringMarshal(t *testing.T) {
	runPrintableStringTest1Marshal(t)
	runPrintableStringTest2Marshal(t)
	runPrintableStringTest3Marshal(t)
}

func TestPrintableStringUnmarshal(t *testing.T) {
	runPrintableStringTest1Unmarshal(t)
	runPrintableStringTest2Unmarshal(t)
	runPrintableStringTest3Unmarshal(t)
}

func runPrintableStringTest1Marshal(t *testing.T) {
	for i, test := range printableStringTestData1 {
		pd := NewPerBitData(nil)
		if err := pd.WritePrintableString(test.unmarshalled, false, nil, nil); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runPrintableStringTest2Marshal(t *testing.T) {
	for i, test := range printableStringTestData2 {
		pd := NewPerBitData(nil)
		var lb, ub uint64
		lb, ub = 1, 1
		if err := pd.WritePrintableString(test.unmarshalled, true, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runPrintableStringTest3Marshal(t *testing.T) {
	for i, test := range printableStringTestData3 {
		pd := NewPerBitData(nil)
		var lb, ub uint64
		lb, ub = 0, 20
		if err := pd.WritePrintableString(test.unmarshalled, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

func runPrintableStringTest1Unmarshal(t *testing.T) {
	for i, test := range printableStringTestData1 {
		pd := NewPerBitData(test.marshalled)
		ps, err := pd.ReadPrintableString(false, nil, nil)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(ps, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runPrintableStringTest2Unmarshal(t *testing.T) {
	for i, test := range printableStringTestData2 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64
		lb, ub = 1, 1
		ps, err := pd.ReadPrintableString(true, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(ps, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runPrintableStringTest3Unmarshal(t *testing.T) {
	for i, test := range printableStringTestData3 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64
		lb, ub = 0, 20
		ps, err := pd.ReadPrintableString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(ps, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

// VisibleString
type visibleStringTestDataStruct struct {
	marshalled   []byte
	unmarshalled VisibleString
}

// unconstrained
var visibleStringTestData1 = []visibleStringTestDataStruct{
	{[]byte{0x00}, VisibleString("")},
	{[]byte{0x05, 0x61, 0x62, 0x63, 0x64, 0x65}, VisibleString("abcde")},
}

// sizeExt, lb=1, ub=1
var visibleStringTestData2 = []visibleStringTestDataStruct{
	{[]byte{0x30, 0x80}, VisibleString("a")},
	{[]byte{0x80, 0x02, 0x61, 0x61}, VisibleString("aa")},
}

// lb=0, ub=20
var visibleStringTestData3 = []visibleStringTestDataStruct{
	{[]byte{0x00}, VisibleString("")},
	{[]byte{0x20, 0x61, 0x61, 0x62, 0x62}, VisibleString("aabb")},
	{[]byte{
		0xA0, 0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62,
		0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62,
	}, VisibleString("aabbaabbaabbaabbaabb")},
}

func TestVisibleStringMarshal(t *testing.T) {
	runVisibleStringTest1Marshal(t)
	runVisibleStringTest2Marshal(t)
	runVisibleStringTest3Marshal(t)
}

func TestVisibleStringUnmarshal(t *testing.T) {
	runVisibleStringTest1Unmarshal(t)
	runVisibleStringTest2Unmarshal(t)
	runVisibleStringTest3Unmarshal(t)
}

func runVisibleStringTest1Marshal(t *testing.T) {
	for i, test := range visibleStringTestData1 {
		pd := NewPerBitData(nil)
		if err := pd.WriteVisibleString(test.unmarshalled, false, nil, nil); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runVisibleStringTest2Marshal(t *testing.T) {
	for i, test := range visibleStringTestData2 {
		pd := NewPerBitData(nil)
		var lb, ub uint64
		lb, ub = 1, 1
		if err := pd.WriteVisibleString(test.unmarshalled, true, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runVisibleStringTest3Marshal(t *testing.T) {
	for i, test := range visibleStringTestData3 {
		pd := NewPerBitData(nil)
		var lb, ub uint64
		lb, ub = 0, 20
		if err := pd.WriteVisibleString(test.unmarshalled, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

func runVisibleStringTest1Unmarshal(t *testing.T) {
	for i, test := range visibleStringTestData1 {
		pd := NewPerBitData(test.marshalled)
		vs, err := pd.ReadVisibleString(false, nil, nil)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(vs, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runVisibleStringTest2Unmarshal(t *testing.T) {
	for i, test := range visibleStringTestData2 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64
		lb, ub = 1, 1
		vs, err := pd.ReadVisibleString(true, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(vs, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runVisibleStringTest3Unmarshal(t *testing.T) {
	for i, test := range visibleStringTestData3 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64
		lb, ub = 0, 20
		vs, err := pd.ReadVisibleString(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(vs, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

// UTF8String
type UTF8StringTestDataStruct struct {
	marshalled   []byte
	unmarshalled UTF8String
}

// unconstrained
var UTF8StringTestData1 = []UTF8StringTestDataStruct{
	{[]byte{0x00}, UTF8String("")},
	{[]byte{0x05, 0x61, 0x62, 0x63, 0x64, 0x65}, UTF8String("abcde")},
}

// sizeExt, lb=1, ub=1
var UTF8StringTestData2 = []UTF8StringTestDataStruct{
	{[]byte{0x01, 0x61}, UTF8String("a")},
	{[]byte{0x02, 0x61, 0x61}, UTF8String("aa")},
}

// lb=0, ub=20
var UTF8StringTestData3 = []UTF8StringTestDataStruct{
	{[]byte{0x00}, UTF8String("")},
	{[]byte{0x04, 0x61, 0x61, 0x62, 0x62}, UTF8String("aabb")},
	{[]byte{
		0x14, 0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62,
		0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62, 0x61, 0x61, 0x62, 0x62,
	}, UTF8String("aabbaabbaabbaabbaabb")},
}

func TestUTF8StringMarshal(t *testing.T) {
	runUTF8StringTest1Marshal(t)
	runUTF8StringTest2Marshal(t)
	runUTF8StringTest3Marshal(t)
}

func TestUTF8StringUnmarshal(t *testing.T) {
	runUTF8StringTest1Unmarshal(t)
	runUTF8StringTest2Unmarshal(t)
	runUTF8StringTest3Unmarshal(t)
}

func runUTF8StringTest1Marshal(t *testing.T) {
	for i, test := range UTF8StringTestData1 {
		pd := NewPerBitData(nil)
		if err := pd.WriteUTF8String(test.unmarshalled, false, nil, nil); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runUTF8StringTest2Marshal(t *testing.T) {
	for i, test := range UTF8StringTestData2 {
		pd := NewPerBitData(nil)
		var lb, ub uint64
		lb, ub = 1, 1
		if err := pd.WriteUTF8String(test.unmarshalled, true, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		fmt.Printf("% x\n", pd.Bytes())
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runUTF8StringTest3Marshal(t *testing.T) {
	for i, test := range UTF8StringTestData3 {
		pd := NewPerBitData(nil)
		var lb, ub uint64
		lb, ub = 0, 20
		if err := pd.WriteUTF8String(test.unmarshalled, false, &lb, &ub); err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

func runUTF8StringTest1Unmarshal(t *testing.T) {
	for i, test := range UTF8StringTestData1 {
		pd := NewPerBitData(test.marshalled)
		utf8s, err := pd.ReadUTF8String(false, nil, nil)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(utf8s, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 1 is FAILED", i+1)
	}
}

func runUTF8StringTest2Unmarshal(t *testing.T) {
	for i, test := range UTF8StringTestData2 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64
		lb, ub = 1, 1
		utf8s, err := pd.ReadUTF8String(true, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(utf8s, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 2 is FAILED", i+1)
	}
}

func runUTF8StringTest3Unmarshal(t *testing.T) {
	for i, test := range UTF8StringTestData3 {
		pd := NewPerBitData(test.marshalled)
		var lb, ub uint64
		lb, ub = 0, 20
		utf8s, err := pd.ReadUTF8String(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		if reflect.DeepEqual(utf8s, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d of TEST 3 is FAILED", i+1)
	}
}

// TEST Open Type
func TestOpenTypeMarshal(t *testing.T) {
	runOpenTypeMarshal(t)
}

func TestOpenTypeUnmarshal(t *testing.T) {
	runOpenTypeUnmarshal(t)
}

type newOpenTypeTest1 struct {
	ID int64 // `aper:"valueLB:0,valueUB:255"`
	// Value newOpenTypeStruct //`aper:"openType,referenceFieldName:ID"`
	List1 *openType1 // `aper:"sizeLB:0,sizeUB:3,referenceFieldValue:2"`
	List2 *openType2 // `aper:"sizeLB:0,sizeUB:30,referenceFieldValue:3"`
	List3 *openType3 // `aper:"sizeLB:0,sizeUB:50,referenceFieldValue:5"`
}

type openType1 struct {
	data []intTest1 //`aper:"sizeLB:0,sizeUB:3"`
}

type openType2 struct {
	data []intStructTest1 //`aper:"sizeLB:0,sizeUB:30"`
}

type openType3 struct {
	data []BitStringStructTest3 //`aper:"sizeLB:0,sizeUB:50"`
}

func (ot *openType1) Write(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 3
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(ot.data)), false, &lb, &ub)
	for i := 0; i < len(ot.data); i++ {
		err = pd.WriteInteger(ot.data[i].Value, false, nil, nil)
	}
	return err
}

func (ot *openType1) Read(pd *PerBitData) error {
	var lb, ub uint64 = 0, 3
	l, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
	if err != nil {
		return err
	}

	// SEQUENCE DATA
	for i := 0; i < int(l); i++ {
		var val intTest1
		if val.Value, err = pd.ReadInteger(false, nil, nil); err != nil {
			return err
		} else {
			ot.data = append(ot.data, val)
		}
	}

	return err
}

func (ot *openType2) Write(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 30
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(ot.data)), false, &lb, &ub)
	for i := 0; i < len(ot.data); i++ {
		var lb, ub int64
		lb, ub = 1, 110
		err = pd.WriteInteger(ot.data[i].Int1, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		err = pd.WriteInteger(ot.data[i].Int2, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 45
		err = pd.WriteInteger(ot.data[i].Int3, true, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
	}
	return err
}

func (ot *openType2) Read(pd *PerBitData) error {
	var lb, ub uint64 = 0, 30
	l, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
	for i := 0; i < int(l); i++ {
		var val intStructTest1
		var lb, ub int64
		lb, ub = 1, 110
		val.Int1, err = pd.ReadInteger(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		val.Int2, err = pd.ReadInteger(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 45
		val.Int3, err = pd.ReadInteger(true, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		ot.data = append(ot.data, val)
	}
	return err
}

func (ot *openType3) Write(pd *PerBitData) error {
	var err error
	var lb, ub uint64 = 0, 50
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(ot.data)), false, &lb, &ub)
	for i := 0; i < len(ot.data); i++ {
		var lb, ub uint64
		lb, ub = 3, 3
		err = pd.WriteBitString(ot.data[i].BitString1, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 125
		err = pd.WriteBitString(ot.data[i].BitString2, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		err = pd.WriteBitString(ot.data[i].BitString3, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 555
		err = pd.WriteBitString(ot.data[i].BitString4, false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
	}
	return err
}

func (ot *openType3) Read(pd *PerBitData) error {
	var lb, ub uint64 = 0, 50
	l, err := pd.ReadSequenceOfPreambleBitMap(false, &lb, &ub)
	for i := 0; i < int(l); i++ {
		var val BitStringStructTest3
		var lb, ub uint64
		lb, ub = 3, 3
		val.BitString1, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 125
		val.BitString2, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 255
		val.BitString3, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}
		lb, ub = 0, 555
		val.BitString4, err = pd.ReadBitString(false, &lb, &ub)
		if err != nil {
			return fmt.Errorf("  [ERROR]: %s", err.Error())
		}

		ot.data = append(ot.data, val)
	}
	return err
}

var newOpenTypeTest1Data = []newOpenTypeTest1{
	{2, &openType1{intTest1Data}, nil, nil},
	{3, nil, &openType2{intStructTest1Data}, nil},
	{5, nil, nil, &openType3{BitStringStructTest3Data}},
}

type newOpenTypeTestDataStruct struct {
	marshalled   []byte
	unmarshalled newOpenTypeTest1
}

var newOpenTypeTestData = []newOpenTypeTestDataStruct{
	{[]byte{0x02, 0x0B, 0xC0, 0x01, 0x03, 0x03, 0x05, 0x16, 0x15, 0x03, 0xFA, 0xE9, 0xEB}, newOpenTypeTest1Data[0]},
	{[]byte{0x03, 0x07, 0x0A, 0xC0, 0x7B, 0x80, 0x02, 0x19, 0x2D}, newOpenTypeTest1Data[1]},
	{[]byte{0x05, 0x08, 0x06, 0x88, 0xFE, 0x06, 0xEC, 0x00, 0x05, 0xD8}, newOpenTypeTest1Data[2]},
}

func runOpenTypeMarshal(t *testing.T) {
	var pd *PerBitData
	var pdOpenType *PerBitData
	var lb, ub int64 = 0, 255
	var err error

	for i, test := range newOpenTypeTestData {
		pd = NewPerBitData(nil)
		pdOpenType = NewPerBitData(nil)

		// ID
		err = pd.WriteInteger(newOpenTypeTest1Data[i].ID, false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}

		// Open Type
		if newOpenTypeTest1Data[i].List1 != nil {
			err = newOpenTypeTest1Data[i].List1.Write(pdOpenType)
		} else if newOpenTypeTest1Data[i].List2 != nil {
			err = newOpenTypeTest1Data[i].List2.Write(pdOpenType)
		} else if newOpenTypeTest1Data[i].List3 != nil {
			err = newOpenTypeTest1Data[i].List3.Write(pdOpenType)
		} else {
			err = fmt.Errorf("Unknown Open Type")
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else {
			if err = pd.WriteOpenType(pdOpenType.Bytes()); err != nil {
				t.Errorf("  [ERROR]: %s", err.Error())
			} else if bytes.Equal(pd.Bytes(), test.marshalled) {
				continue
			}
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

func runOpenTypeUnmarshal(t *testing.T) {
	var pd *PerBitData
	var pdOpenType *PerBitData
	var lb, ub int64 = 0, 255

	for i, test := range newOpenTypeTestData {
		var result newOpenTypeTest1
		var err error
		pd = NewPerBitData(test.marshalled)

		// ID
		result.ID, err = pd.ReadInteger(false, &lb, &ub)
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}

		// Open Type - referenceField: ID
		openTypeBytes, err := pd.ReadOpenType()
		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		}
		pdOpenType = NewPerBitData(openTypeBytes)
		if result.ID == 2 {
			result.List1 = &openType1{}
			err = result.List1.Read(pdOpenType)
		} else if result.ID == 3 {
			result.List2 = &openType2{}
			err = result.List2.Read(pdOpenType)
		} else if result.ID == 5 {
			result.List3 = &openType3{}
			err = result.List3.Read(pdOpenType)
		} else {
			err = fmt.Errorf("Unknown reference field value %d", result.ID)
		}

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else {
			if reflect.DeepEqual(result, test.unmarshalled) {
				continue
			}
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

// TEST BOOLEAN
func TestBoolMarshal(t *testing.T) {
	runBoolTest1Marshal(t)
}

func TestBoolUnarshal(t *testing.T) {
	runBoolTest1Unmarshal(t)
}

type boolTestDataStruct1 struct {
	marshalled   []byte
	unmarshalled boolTest1
}

var boolTestData1 = []boolTestDataStruct1{
	{[]byte{0x00}, boolTest1Data[0]},
	{[]byte{0x80}, boolTest1Data[1]},
}

func runBoolTest1Marshal(t *testing.T) {
	for i, test := range boolTestData1 {
		pd := NewPerBitData(nil)
		err := pd.WriteBool(test.unmarshalled.Value)

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if bytes.Equal(pd.Bytes(), test.marshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d is FAILED", i+1)
	}
}

func runBoolTest1Unmarshal(t *testing.T) {
	for i, test := range boolTestData1 {
		pd := NewPerBitData(test.marshalled)

		var result boolTest1
		var err error
		result.Value, err = pd.ReadBool()

		if err != nil {
			t.Errorf("  [ERROR]: %s", err.Error())
		} else if reflect.DeepEqual(result, test.unmarshalled) {
			continue
		}
		t.Errorf("[FAIL]\n")
		t.Errorf("TEST %d is FAILED", i+1)
	}
}
