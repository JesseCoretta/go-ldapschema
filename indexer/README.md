# indexer

Package indexer wraps [JesseCoretta/go-ldapschema](https://github.com/JesseCoretta/go-ldapschema) to generate fast schema lookup tables, vastly improving response time as compared to use of a native `*schema.SubentrySubschema` instance.

## Example

```go
package main

import (
	"github.com/JesseCoretta/go-ldapschema"
	"github.com/JesseCoretta/go-ldapschema/indexer"
)

func main() {
	sch, err := schema.NewSubschemaSubentry(true)
	if err != nil {
		panic(errr)
	}

	//
	// Populate your schema as you normally would.
	//

	// Feed schema to indexer.New.
	idx, err := indexer.New(sch)
	if err != nil {
		panic(errr)
	}

	// You now have a complete, high-performance schema index.
	//
	// Resolve an attribyte by numeric OID or descriptor. Most
	// map indices use numeric OID.
	numericOID, descriptor, altNames := r.AT.Resolve(`cn`)      // "2.5.4.3" descriptor
	numericOID, descriptor, altNames := r.AT.Resolve(`2.5.4.3`) // "cn" numeric OID
	
	// Get the effective syntax for an attribute
	syn := r.AT.EFS[numericOID]   // directory string

	// Get the super type of an attribute
	super := r.AT.Sup[numericOID] // "name" is super type
}
```
