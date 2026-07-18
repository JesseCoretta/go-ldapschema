package indexer

import (
	"fmt"
	"testing"
)

func ExampleLDAPSyntaxProperties_Resolve_reverseLookup() {
	numericOID, text, _ := exampleIndex.LS.Resolve(`1.3.6.1.4.1.1466.115.121.1.3`)
	fmt.Printf("Principal name for %s: %q", numericOID, text)
	// Output: Principal name for 1.3.6.1.4.1.1466.115.121.1.3: "Attribute Type Description"
}

func ExampleLDAPSyntaxProperties_Resolve_forwardLookup() {
	numericOID, text, _ := exampleIndex.LS.Resolve(`Attribute Type Description`)
	fmt.Printf("Principal name for %s: %q", numericOID, text)
	// Output: Principal name for 1.3.6.1.4.1.1466.115.121.1.3: "Attribute Type Description"
}

func BenchmarkLDAPSyntaxCalls(b *testing.B) {

	var synk []string
	for k := range exampleIndex.LS.O2D {
		// list of syntax OIDs
		synk = append(synk, k)
	}

	b.StopTimer()
	maxIdx := len(synk)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = exampleIndex.LS.D2O[synk[i%maxIdx]]
		_, _ = exampleIndex.LS.NotHR[synk[i%maxIdx]]
		_, _ = exampleIndex.LS.AT[synk[i%maxIdx]]
		_, _ = exampleIndex.LS.MR[synk[i%maxIdx]]
		_, _ = exampleIndex.LS.SrcIndex[synk[i%maxIdx]]
	}
}
