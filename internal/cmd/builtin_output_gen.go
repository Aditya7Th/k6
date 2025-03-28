// Code generated by "enumer -type=builtinOutput -trimprefix builtinOutput -transform=kebab -output builtin_output_gen.go"; DO NOT EDIT.

package cmd

import (
	"fmt"
	"strings"
)

const _builtinOutputName = "cloudcsvdatadogexperimental-prometheus-rwinfluxdbjsonkafkastatsdexperimental-opentelemetrysummary"

var _builtinOutputIndex = [...]uint8{0, 5, 8, 15, 41, 49, 53, 58, 64, 90, 97}

const _builtinOutputLowerName = "cloudcsvdatadogexperimental-prometheus-rwinfluxdbjsonkafkastatsdexperimental-opentelemetrysummary"

func (i builtinOutput) String() string {
	if i >= builtinOutput(len(_builtinOutputIndex)-1) {
		return fmt.Sprintf("builtinOutput(%d)", i)
	}
	return _builtinOutputName[_builtinOutputIndex[i]:_builtinOutputIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _builtinOutputNoOp() {
	var x [1]struct{}
	_ = x[builtinOutputCloud-(0)]
	_ = x[builtinOutputCSV-(1)]
	_ = x[builtinOutputDatadog-(2)]
	_ = x[builtinOutputExperimentalPrometheusRW-(3)]
	_ = x[builtinOutputInfluxdb-(4)]
	_ = x[builtinOutputJSON-(5)]
	_ = x[builtinOutputKafka-(6)]
	_ = x[builtinOutputStatsd-(7)]
	_ = x[builtinOutputExperimentalOpentelemetry-(8)]
	_ = x[builtinOutputSummary-(9)]
}

var _builtinOutputValues = []builtinOutput{builtinOutputCloud, builtinOutputCSV, builtinOutputDatadog, builtinOutputExperimentalPrometheusRW, builtinOutputInfluxdb, builtinOutputJSON, builtinOutputKafka, builtinOutputStatsd, builtinOutputExperimentalOpentelemetry, builtinOutputSummary}

var _builtinOutputNameToValueMap = map[string]builtinOutput{
	_builtinOutputName[0:5]:        builtinOutputCloud,
	_builtinOutputLowerName[0:5]:   builtinOutputCloud,
	_builtinOutputName[5:8]:        builtinOutputCSV,
	_builtinOutputLowerName[5:8]:   builtinOutputCSV,
	_builtinOutputName[8:15]:       builtinOutputDatadog,
	_builtinOutputLowerName[8:15]:  builtinOutputDatadog,
	_builtinOutputName[15:41]:      builtinOutputExperimentalPrometheusRW,
	_builtinOutputLowerName[15:41]: builtinOutputExperimentalPrometheusRW,
	_builtinOutputName[41:49]:      builtinOutputInfluxdb,
	_builtinOutputLowerName[41:49]: builtinOutputInfluxdb,
	_builtinOutputName[49:53]:      builtinOutputJSON,
	_builtinOutputLowerName[49:53]: builtinOutputJSON,
	_builtinOutputName[53:58]:      builtinOutputKafka,
	_builtinOutputLowerName[53:58]: builtinOutputKafka,
	_builtinOutputName[58:64]:      builtinOutputStatsd,
	_builtinOutputLowerName[58:64]: builtinOutputStatsd,
	_builtinOutputName[64:90]:      builtinOutputExperimentalOpentelemetry,
	_builtinOutputLowerName[64:90]: builtinOutputExperimentalOpentelemetry,
	_builtinOutputName[90:97]:      builtinOutputSummary,
	_builtinOutputLowerName[90:97]: builtinOutputSummary,
}

var _builtinOutputNames = []string{
	_builtinOutputName[0:5],
	_builtinOutputName[5:8],
	_builtinOutputName[8:15],
	_builtinOutputName[15:41],
	_builtinOutputName[41:49],
	_builtinOutputName[49:53],
	_builtinOutputName[53:58],
	_builtinOutputName[58:64],
	_builtinOutputName[64:90],
	_builtinOutputName[90:97],
}

// builtinOutputString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func builtinOutputString(s string) (builtinOutput, error) {
	if val, ok := _builtinOutputNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _builtinOutputNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to builtinOutput values", s)
}

// builtinOutputValues returns all values of the enum
func builtinOutputValues() []builtinOutput {
	return _builtinOutputValues
}

// builtinOutputStrings returns a slice of all String values of the enum
func builtinOutputStrings() []string {
	strs := make([]string, len(_builtinOutputNames))
	copy(strs, _builtinOutputNames)
	return strs
}

// IsAbuiltinOutput returns "true" if the value is listed in the enum definition. "false" otherwise
func (i builtinOutput) IsAbuiltinOutput() bool {
	for _, v := range _builtinOutputValues {
		if i == v {
			return true
		}
	}
	return false
}
