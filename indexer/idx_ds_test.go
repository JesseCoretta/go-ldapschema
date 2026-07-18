package indexer

import (
	"fmt"
	"testing"
)

func ExampleDITStructureRuleProperties_Resolve_reverseLookup() {
	ruleID, descriptor, _ := exampleIndex.DS.Resolve(`1`)
	fmt.Printf("Principal name for rule %s: %q", ruleID, descriptor)
	// Output: Principal name for rule 1: "applicationProcessStructure"
}

func ExampleDITStructureRuleProperties_Resolve_forwardLookup() {
	numericOID, descriptor, _ := exampleIndex.DS.Resolve(`applicationProcessStructure`)
	fmt.Printf("Principal name for rule %s: %q", numericOID, descriptor)
	// Output: Principal name for rule 1: "applicationProcessStructure"
}

func ExampleDITStructureRuleProperties_SuperRules() {
	sub := exampleIndex.DS.SuperRules(`2`)
	fmt.Printf("SuperRules of 2: %v", sub)
	// Output: SuperRules of 2: [1 2]
}

func ExampleDITStructureRuleProperties_SubRules() {
	sub := exampleIndex.DS.SubRules(`1`)
	fmt.Printf("SubRules of 2: %v", sub)
	// Output: SubRules of 2: [2]
}

func BenchmarkDITStructureRuleCalls(b *testing.B) {

	var rule []string
	for k := range exampleIndex.DS.O2D {
		// list of structure rule IDs
		rule = append(rule, k)
	}

	b.StopTimer()
	maxIdx := len(rule)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = exampleIndex.DS.D2O[rule[i%maxIdx]]
		_, _ = exampleIndex.DS.Obsolete[rule[i%maxIdx]]
		_, _ = exampleIndex.DS.NF[rule[i%maxIdx]]
		_, _ = exampleIndex.DS.NOC[rule[i%maxIdx]]
		_, _ = exampleIndex.DS.Sup[rule[i%maxIdx]]
		_, _ = exampleIndex.DS.Sub[rule[i%maxIdx]]
		_, _ = exampleIndex.DS.SrcIndex[rule[i%maxIdx]]
	}
}
