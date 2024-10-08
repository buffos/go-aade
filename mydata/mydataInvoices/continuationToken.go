package mydataInvoices

import "fmt"

type ContinuationToken struct {
	NextPartitionKey string `xml:"nextPartitionKey"`
	NextRowKey       string `xml:"nextRowKey"`
}

func (c *ContinuationToken) IsEmpty() bool {
	return c.NextPartitionKey == "" && c.NextRowKey == ""
}

func (c *ContinuationToken) Print() {
	fmt.Printf("Continuation token: %s\n", c.NextPartitionKey)
	fmt.Printf("Continuation token: %s\n", c.NextRowKey)
}
