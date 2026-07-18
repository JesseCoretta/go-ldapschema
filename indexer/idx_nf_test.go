package indexer

import (
	"fmt"
	"testing"
)

func ExampleNameFormProperties_Resolve_reverseLookup() {
	numericOID, descriptor, _ := exampleIndex.NF.Resolve(`1.3.6.1.4.1.56521.999.1234`)
	fmt.Printf("Principal name for %s: %q", numericOID, descriptor)
	// Output: Principal name for 1.3.6.1.4.1.56521.999.1234: "applicationProcessNameForm"
}

func ExampleNameFormProperties_Resolve_forwardLookup() {
	numericOID, descriptor, _ := exampleIndex.NF.Resolve(`applicationProcessNameForm`)
	fmt.Printf("Principal name for %s: %q", numericOID, descriptor)
	// Output: Principal name for 1.3.6.1.4.1.56521.999.1234: "applicationProcessNameForm"
}

func BenchmarkNameFormCalls(b *testing.B) {
	var form []string
	for k := range exampleIndex.NF.O2D {
		// list of name forms
		form = append(form, k)
	}

	b.StopTimer()
	maxIdx := len(form)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = exampleIndex.NF.D2O[form[i%maxIdx]]
		_, _ = exampleIndex.NF.Obsolete[form[i%maxIdx]]
		_, _ = exampleIndex.NF.DS[form[i%maxIdx]]
		_, _ = exampleIndex.NF.Must[form[i%maxIdx]]
		_, _ = exampleIndex.NF.May[form[i%maxIdx]]
		_, _ = exampleIndex.NF.SrcIndex[form[i%maxIdx]]
	}
}
