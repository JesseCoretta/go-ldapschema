package indexer

import (
	"fmt"
	"testing"
)

func ExampleMatchingRuleProperties_Resolve_reverseLookup() {
	numericOID, descriptor, _ := exampleIndex.MR.Resolve(`2.5.13.17`)
	fmt.Printf("Principal name for %s: %q", numericOID, descriptor)
	// Output: Principal name for 2.5.13.17: "octetStringMatch"
}

func ExampleMatchingRuleProperties_Resolve_forwardLookup() {
	numericOID, descriptor, _ := exampleIndex.MR.Resolve(`octetStringMatch`)
	fmt.Printf("Principal name for %s: %q", numericOID, descriptor)
	// Output: Principal name for 2.5.13.17: "octetStringMatch"
}

func BenchmarkMatchingRuleCalls(b *testing.B) {
	var mrule []string
	for k := range exampleIndex.MR.O2D {
		// list of matching rule OIDs
		mrule = append(mrule, k)
	}

	b.StopTimer()
	maxIdx := len(mrule)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = exampleIndex.MR.D2O[mrule[i%maxIdx]]
		_, _ = exampleIndex.MR.Obsolete[mrule[i%maxIdx]]
		_, _ = exampleIndex.MR.Type[mrule[i%maxIdx]]
		_, _ = exampleIndex.MR.Applies[mrule[i%maxIdx]]
		_, _ = exampleIndex.MR.SrcIndex[mrule[i%maxIdx]]
	}
}
