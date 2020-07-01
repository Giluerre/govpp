//  Copyright (c) 2020 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package binapigen

import (
	"git.fd.io/govpp.git/examples/binapi/interfaces"
	"git.fd.io/govpp.git/examples/binapi/ip_types"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"git.fd.io/govpp.git/binapigen/vppapi"
)

const testOutputDir = "test_output_directory"

func GenerateFromFile(file, outputDir string, opts Options) error {
	apifile, err := vppapi.ParseFile(file)
	if err != nil {
		return err
	}

	g, err := New(opts, []*vppapi.File{apifile})
	if err != nil {
		return err
	}
	for _, file := range g.Files {
		if !file.Generate {
			continue
		}
		GenerateBinapi(g, file, outputDir)
		if file.Service != nil {
			GenerateRPC(g, file, outputDir)
		}
	}

	if err = g.Generate(); err != nil {
		return err
	}

	return nil
}

func TestGenerateFromFile(t *testing.T) {
	RegisterTestingT(t)

	// remove directory created during test
	defer os.RemoveAll(testOutputDir)

	err := GenerateFromFile("vppapi/testdata/acl.api.json", testOutputDir, Options{FilesToGenerate: []string{"acl"}})
	Expect(err).ShouldNot(HaveOccurred())
	fileInfo, err := os.Stat(testOutputDir + "/acl/acl.ba.go")
	Expect(err).ShouldNot(HaveOccurred())
	Expect(fileInfo.IsDir()).To(BeFalse())
	Expect(fileInfo.Name()).To(BeEquivalentTo("acl.ba.go"))
}

func TestGenerateFromFileInputError(t *testing.T) {
	RegisterTestingT(t)

	err := GenerateFromFile("vppapi/testdata/nonexisting.json", testOutputDir, Options{})
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).To(ContainSubstring("unsupported"))
}

func TestGenerateFromFileReadJsonError(t *testing.T) {
	RegisterTestingT(t)

	err := GenerateFromFile("vppapi/testdata/input-read-json-error.json", testOutputDir, Options{})
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).To(ContainSubstring("unsupported"))
}

func TestGenerateFromFileGeneratePackageError(t *testing.T) {
	RegisterTestingT(t)

	// generate package throws panic, recover after it
	defer func() {
		if recovery := recover(); recovery != nil {
			t.Logf("Recovered from panic: %v", recovery)
		}
		os.RemoveAll(testOutputDir)
	}()

	err := GenerateFromFile("vppapi/testdata/input-generate-error.json", testOutputDir, Options{})
	Expect(err).Should(HaveOccurred())
}

func TestGeneratedParseAddress(t *testing.T) {
	RegisterTestingT(t)

	var data = []struct {
		input  string
		result ip_types.Address
	}{
		{"192.168.0.1", ip_types.Address{
			Af: ip_types.ADDRESS_IP4,
			Un: ip_types.AddressUnionIP4(ip_types.IP4Address{192, 168, 0, 1}),
		}},
		{"aac1:0:ab45::", ip_types.Address{
			Af: ip_types.ADDRESS_IP6,
			Un: ip_types.AddressUnionIP6(ip_types.IP6Address{170, 193, 0, 0, 171, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		}},
	}

	for _, entry := range data {
		t.Run(entry.input, func(t *testing.T) {
			parsedAddress, err := ip_types.ParseAddress(entry.input)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(parsedAddress).To(Equal(entry.result))

			originAddress := parsedAddress.ToString()
			Expect(originAddress).To(Equal(entry.input))
		})
	}
}

func TestGeneratedParseAddressError(t *testing.T) {
	RegisterTestingT(t)

	_, err := ip_types.ParseAddress("malformed_ip")
	Expect(err).Should(HaveOccurred())
}

func TestGeneratedParsePrefix(t *testing.T) {
	RegisterTestingT(t)

	var data = []struct {
		input  string
		result ip_types.Prefix
	}{
		{"192.168.0.1/24", ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{192, 168, 0, 1}),
			},
			Len: 24,
		}},
		{"192.168.0.1", ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{192, 168, 0, 1}),
			},
			Len: 32,
		}},
		{"aac1:0:ab45::/96", ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP6,
				Un: ip_types.AddressUnionIP6(ip_types.IP6Address{170, 193, 0, 0, 171, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
			},
			Len: 96,
		}},
		{"aac1:0:ab45::", ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP6,
				Un: ip_types.AddressUnionIP6(ip_types.IP6Address{170, 193, 0, 0, 171, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
			},
			Len: 128,
		}},
	}

	for _, entry := range data {
		t.Run(entry.input, func(t *testing.T) {
			parsedAddress, err := ip_types.ParsePrefix(entry.input)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(parsedAddress).To(Equal(entry.result))

			// Parsed IP without prefix receives a default one
			// so the input data must be adjusted
			if entry.result.Address.Af == ip_types.ADDRESS_IP4 && !strings.Contains(entry.input, "/") {
				entry.input = entry.input + "/32"
			}
			if entry.result.Address.Af == ip_types.ADDRESS_IP6 && !strings.Contains(entry.input, "/") {
				entry.input = entry.input + "/128"
			}
			originAddress := parsedAddress.ToString()
			Expect(originAddress).To(Equal(entry.input))
		})
	}
}

func TestGeneratedParsePrefixError(t *testing.T) {
	RegisterTestingT(t)

	_, err := ip_types.ParsePrefix("malformed_ip")
	Expect(err).Should(HaveOccurred())
}

func TestGeneratedParseMAC(t *testing.T) {
	RegisterTestingT(t)

	var data = []struct {
		input  string
		result interfaces.MacAddress
	}{
		{"b7:b9:bb:a1:5c:af", interfaces.MacAddress{183, 185, 187, 161, 92, 175}},
		{"47:4b:c7:3e:06:c8", interfaces.MacAddress{71, 75, 199, 62, 6, 200}},
		{"a7:cc:9f:10:18:e3", interfaces.MacAddress{167, 204, 159, 16, 24, 227}},
	}

	for _, entry := range data {
		t.Run(entry.input, func(t *testing.T) {
			parsedMac, err := interfaces.ParseMAC(entry.input)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(parsedMac).To(Equal(entry.result))

			originAddress := parsedMac.ToString()
			Expect(originAddress).To(Equal(entry.input))
		})
	}
}

func TestGeneratedParseMACError(t *testing.T) {
	RegisterTestingT(t)

	_, err := interfaces.ParseMAC("malformed_mac")
	Expect(err).Should(HaveOccurred())
}

/*func TestGetContext(t *testing.T) {
	RegisterTestingT(t)
	outDir := "test_output_directory"
	result, err := newContext("testdata/af_packet.api.json", outDir)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(result).ToNot(BeNil())
	Expect(result.outputFile).To(BeEquivalentTo(outDir + "/af_packet/af_packet.ba.go"))
}

func TestGetContextNoJsonFile(t *testing.T) {
	RegisterTestingT(t)
	outDir := "test_output_directory"
	result, err := newContext("testdata/input.txt", outDir)
	Expect(err).Should(HaveOccurred())
	Expect(err.Error()).To(ContainSubstring("invalid input file name"))
	Expect(result).To(BeNil())
}

func TestGetContextInterfaceJson(t *testing.T) {
	RegisterTestingT(t)
	outDir := "test_output_directory"
	result, err := newContext("testdata/ip.api.json", outDir)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(result).ToNot(BeNil())
	Expect(result.outputFile)
	Expect(result.outputFile).To(BeEquivalentTo(outDir + "/ip/ip.ba.go"))
}*/

/*func TestGeneratePackage(t *testing.T) {
	RegisterTestingT(t)
	// prepare context
	testCtx := new(GenFile)
	testCtx.packageName = "test-package-name"

	// prepare input/output output files
	inputData, err := ioutil.ReadFile("testdata/ip.api.json")
	Expect(err).ShouldNot(HaveOccurred())
	jsonRoot, err := parseInputJSON(inputData)
	Expect(err).ShouldNot(HaveOccurred())
	testCtx.file, err = parseModule(testCtx, jsonRoot)
	Expect(err).ShouldNot(HaveOccurred())
	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	defer os.RemoveAll(outDir)

	// prepare writer
	writer := bufio.NewWriter(outFile)
	Expect(writer.Buffered()).To(BeZero())
	err = generateFileBinapi(testCtx, writer)
	Expect(err).ShouldNot(HaveOccurred())
}

func TestGenerateMessageType(t *testing.T) {
	RegisterTestingT(t)
	// prepare context
	testCtx := new(GenFile)
	testCtx.packageName = "test-package-name"

	// prepare input/output output files
	inputData, err := ioutil.ReadFile("testdata/ip.api.json")
	Expect(err).ShouldNot(HaveOccurred())
	jsonRoot, err := parseInputJSON(inputData)
	Expect(err).ShouldNot(HaveOccurred())
	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	testCtx.file, err = parseModule(testCtx, jsonRoot)
	Expect(err).ShouldNot(HaveOccurred())
	defer os.RemoveAll(outDir)

	// prepare writer
	writer := bufio.NewWriter(outFile)

	for _, msg := range testCtx.file.Messages {
		generateMessage(testCtx, writer, &msg)
		Expect(writer.Buffered()).ToNot(BeZero())
	}
}*/

/*func TestGenerateMessageName(t *testing.T) {
	RegisterTestingT(t)
	// prepare context
	testCtx := new(context)
	testCtx.packageName = "test-package-name"

	// prepare input/output output files
	inputData, err := readFile("testdata/ip.api.json")
	Expect(err).ShouldNot(HaveOccurred())
	testCtx.inputBuff = bytes.NewBuffer(inputData)
	inFile, _ := parseJSON(inputData)
	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	defer os.RemoveAll(outDir)

	// prepare writer
	writer := bufio.NewWriter(outFile)

	types := inFile.Map("types")
	Expect(types.Len()).To(BeEquivalentTo(1))
	for i := 0; i < types.Len(); i++ {
		typ := types.At(i)
		Expect(writer.Buffered()).To(BeZero())
		err := generateMessage(testCtx, writer, typ, false)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(writer.Buffered()).ToNot(BeZero())

	}
}

func TestGenerateMessageFieldTypes(t *testing.T) {
	// expected results according to acl.api.json in testdata
	expectedTypes := []string{
		"\tIsPermit uint8",
		"\tIsIpv6 uint8",
		"\tSrcIPAddr []byte	`struc:\"[16]byte\"`",
		"\tSrcIPPrefixLen uint8",
		"\tDstIPAddr []byte	`struc:\"[16]byte\"`",
		"\tDstIPPrefixLen uint8",
		"\tProto uint8",
		"\tSrcportOrIcmptypeFirst uint16",
		"\tSrcportOrIcmptypeLast uint16",
		"\tDstportOrIcmpcodeFirst uint16",
		"\tDstportOrIcmpcodeLast uint16",
		"\tTCPFlagsMask uint8",
		"\tTCPFlagsValue uint8"}
	RegisterTestingT(t)
	// prepare context
	testCtx := new(context)
	testCtx.packageName = "test-package-name"

	// prepare input/output output files
	inputData, err := readFile("testdata/acl.api.json")
	Expect(err).ShouldNot(HaveOccurred())
	inFile, err := parseJSON(inputData)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(inFile).ToNot(BeNil())

	// test types
	types := inFile.Map("types")
	fields := make([]string, 0)
	for i := 0; i < types.Len(); i++ {
		for j := 0; j < types.At(i).Len(); j++ {
			field := types.At(i).At(j)
			if field.GetType() == jsongo.TypeArray {
				err := processMessageField(testCtx, &fields, field, false)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(fields[j-1]).To(BeEquivalentTo(expectedTypes[j-1]))
			}
		}
	}
}

func TestGenerateMessageFieldMessages(t *testing.T) {
	// expected results according to acl.api.json in testdata
	expectedFields := []string{"\tMajor uint32", "\tMinor uint32", "\tRetval int32",
		"\tVpePid uint32", "\tACLIndex uint32", "\tTag []byte	`struc:\"[64]byte\"`",
		"\tCount uint32"}
	RegisterTestingT(t)
	// prepare context
	testCtx := new(context)
	testCtx.packageName = "test-package-name"

	// prepare input/output output files
	inputData, err := readFile("testdata/acl.api.json")
	Expect(err).ShouldNot(HaveOccurred())
	inFile, err := parseJSON(inputData)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(inFile).ToNot(BeNil())

	// test message fields
	messages := inFile.Map("messages")
	customIndex := 0
	fields := make([]string, 0)
	for i := 0; i < messages.Len(); i++ {
		for j := 0; j < messages.At(i).Len(); j++ {
			field := messages.At(i).At(j)
			if field.GetType() == jsongo.TypeArray {
				specificFieldName := field.At(1).Get().(string)
				if specificFieldName == "crc" || specificFieldName == "_vl_msg_id" ||
					specificFieldName == "client_index" || specificFieldName == "context" {
					continue
				}
				err := processMessageField(testCtx, &fields, field, false)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(fields[customIndex]).To(BeEquivalentTo(expectedFields[customIndex]))
				customIndex++
				if customIndex >= len(expectedFields) {
					// there is too much fields now for one UT...
					return
				}
			}
		}
	}
}

func TestGeneratePackageHeader(t *testing.T) {
	RegisterTestingT(t)
	// prepare context
	testCtx := new(context)
	testCtx.packageName = "test-package-name"

	// prepare input/output output files
	inputData, err := readFile("testdata/acl.api.json")
	Expect(err).ShouldNot(HaveOccurred())
	inFile, err := parseJSON(inputData)
	Expect(err).ShouldNot(HaveOccurred())
	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	defer os.RemoveAll(outDir)
	// prepare writer
	writer := bufio.NewWriter(outFile)
	Expect(writer.Buffered()).To(BeZero())
	generatePackageHeader(testCtx, writer, inFile)
	Expect(writer.Buffered()).ToNot(BeZero())
}

func TestGenerateMessageCommentType(t *testing.T) {
	RegisterTestingT(t)
	// prepare context
	testCtx := new(context)
	testCtx.packageName = "test-package-name"
	testCtx.inputBuff = bytes.NewBuffer([]byte("test content"))

	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	writer := bufio.NewWriter(outFile)
	defer os.RemoveAll(outDir)
	Expect(writer.Buffered()).To(BeZero())
	generateMessageComment(testCtx, writer, "test-struct", "msg-name", true)
	Expect(writer.Buffered()).ToNot(BeZero())
}

func TestGenerateMessageCommentMessage(t *testing.T) {
	RegisterTestingT(t)
	// prepare context
	testCtx := new(context)
	testCtx.packageName = "test-package-name"
	testCtx.inputBuff = bytes.NewBuffer([]byte("test content"))

	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	writer := bufio.NewWriter(outFile)
	defer os.RemoveAll(outDir)
	Expect(writer.Buffered()).To(BeZero())
	generateMessageComment(testCtx, writer, "test-struct", "msg-name", false)
	Expect(writer.Buffered()).ToNot(BeZero())
}

func TestGenerateMessageNameGetter(t *testing.T) {
	RegisterTestingT(t)
	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	writer := bufio.NewWriter(outFile)
	defer os.RemoveAll(outDir)
	Expect(writer.Buffered()).To(BeZero())
	generateMessageNameGetter(writer, "test-struct", "msg-name")
	Expect(writer.Buffered()).ToNot(BeZero())
}

func TestGenerateTypeNameGetter(t *testing.T) {
	RegisterTestingT(t)
	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	writer := bufio.NewWriter(outFile)
	defer os.RemoveAll(outDir)
	Expect(writer.Buffered()).To(BeZero())
	generateTypeNameGetter(writer, "test-struct", "msg-name")
	Expect(writer.Buffered()).ToNot(BeZero())
}

func TestGenerateCrcGetter(t *testing.T) {
	RegisterTestingT(t)
	outDir := "test_output_directory"
	outFile, err := os.Create(outDir)
	Expect(err).ShouldNot(HaveOccurred())
	writer := bufio.NewWriter(outFile)
	defer os.RemoveAll(outDir)
	Expect(writer.Buffered()).To(BeZero())
	generateCrcGetter(writer, "test-struct", "msg-name")
	Expect(writer.Buffered()).ToNot(BeZero())
}

func TestTranslateVppType(t *testing.T) {
	RegisterTestingT(t)
	context := new(context)
	typesToTranslate := []string{"u8", "i8", "u16", "i16", "u32", "i32", "u64", "i64", "f64"}
	expected := []string{"uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "float64"}
	var translated []string
	for _, value := range typesToTranslate {
		translated = append(translated, convertToGoType(context, value, false))
	}
	for index, value := range expected {
		Expect(value).To(BeEquivalentTo(translated[index]))
	}

}

func TestTranslateVppTypeArray(t *testing.T) {
	RegisterTestingT(t)
	context := new(context)
	translated := convertToGoType(context, "u8", true)
	Expect(translated).To(BeEquivalentTo("byte"))
}

func TestTranslateVppUnknownType(t *testing.T) {
	defer func() {
		if recovery := recover(); recovery != nil {
			t.Logf("Recovered from panic: %v", recovery)
		}
	}()
	context := new(context)
	convertToGoType(context, "?", false)
}

func TestCamelCase(t *testing.T) {
	RegisterTestingT(t)
	// test camel case functionality
	expected := "allYourBaseAreBelongToUs"
	result := camelCaseName("all_your_base_are_belong_to_us")
	Expect(expected).To(BeEquivalentTo(result))
	// test underscore
	expected = "_"
	result = camelCaseName(expected)
	Expect(expected).To(BeEquivalentTo(result))
	// test all lower
	expected = "lower"
	result = camelCaseName(expected)
	Expect(expected).To(BeEquivalentTo(result))
}

func TestCommonInitialisms(t *testing.T) {
	RegisterTestingT(t)

	for key, value := range commonInitialisms {
		Expect(value).ShouldNot(BeFalse())
		Expect(key).ShouldNot(BeEmpty())
	}
}
*/
