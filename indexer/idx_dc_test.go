package indexer

import (
	"fmt"
	"testing"
)

func ExampleDITContentRuleProperties_Resolve_reverseLookup() {
	numericOID, descriptor, _ := exampleIndex.DC.Resolve(`0.9.2342.19200300.100.4.7`)
	fmt.Printf("Principal name for %s: %q", numericOID, descriptor)
	// Output: Principal name for 0.9.2342.19200300.100.4.7: "roomContent"
}

func ExampleDITContentRuleProperties_Resolve_forwardLookup() {
	numericOID, descriptor, _ := exampleIndex.DC.Resolve(`roomContent`)
	fmt.Printf("Principal name for %s: %q", numericOID, descriptor)
	// Output: Principal name for 0.9.2342.19200300.100.4.7: "roomContent"
}

func BenchmarkDITContentRuleCalls(b *testing.B) {

	var rule []string
	for k := range exampleIndex.DC.O2D {
		// list of DIT content rules
		rule = append(rule, k)
	}

	b.StopTimer()
	maxIdx := len(rule)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = exampleIndex.DC.D2O[rule[i%maxIdx]]
		_, _ = exampleIndex.DC.Aux[rule[i%maxIdx]]
		_, _ = exampleIndex.DC.Obsolete[rule[i%maxIdx]]
		_, _ = exampleIndex.DC.Must[rule[i%maxIdx]]
		_, _ = exampleIndex.DC.May[rule[i%maxIdx]]
		_, _ = exampleIndex.DC.Not[rule[i%maxIdx]]
		_, _ = exampleIndex.DC.SrcIndex[rule[i%maxIdx]]
	}
}
