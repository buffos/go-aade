package mydataInvoices

type ContinuationToken struct {
	NextPartitionKey string `xml:"nextPartitionKey"`
	NextRowKey       string `xml:"nextRowKey"`
}
