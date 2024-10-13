# Documentation for requesters in mydata package

## Overview

The `mydata` package provides functionality to interact with the myDATA API, allowing users to request various types of documents, VAT information, and income/expense data. The package includes both direct request functions and iterator functions for paginated results.

## Client Initialization

Before using any functions, you need to initialize a `Client` instance:

```go
c := NewClient(userID, subscriptionKey, timeoutInSeconds, isProduction)
```

- `userID`: Your user ID for authentication.
- `subscriptionKey`: Your subscription key for accessing the API.
- `timeoutInSeconds`: Timeout for the operation
- `isProduction`: If we use the Aade in production env or not.

### Error handling and validation

There is also a parameter on what action to do when there is an invalid invoice while we are sending info.
It has three values:

- ErrorOnInvalid
- PassThroughOnInvalid
- FilterOnInvalid

Default values is `ErrorOnInvalid` and you can set it by `c.SetOnInvalid(value)`

Before sending invoice, `SendInvoices` function tries to do basic validation and these parameters handle what to do.
If `ErrorOnInvalid`, the function returns an error before sending any invoices
If `PassThroughOnInvalid` then we send the invoices even if there are errors
If `FilterOnInvalid` we remove those invoices that are invalid and send the rest to the api.

## Request Functions

### 1. RequestDocs

Fetches requested documents submitted by others.

```go
code, docs, err := c.RequestDocs(mydataInvoices.RequestDocsParams{
    Mark: "0",
    DateFrom: time.Now().AddDate(0, -16, 0), // last 16 days
    DateTo: time.Now(),
})
```

There are some convenience functions

#### 1.1. RequestDocsPastDays

Fetches documents submitted in the past specified number of days.

```go
code, docs, err := c.RequestDocsPastDays(7) // last 7 days
```

#### 1.2. RequestDocWithMark

Fetches documents with a specific mark.

```go
code, docs, err := c.RequestDocWithMark(1) // mark 1
```

### 2. RequestVatInfo

Fetches VAT information based on specified parameters.

```go
code, docs, err := c.RequestVatInfo(mydataInvoices.RequestVatInfoParams{
        DateFrom: time.Now().UTC().AddDate(0, -12, 0),
        DateTo: time.Now().UTC(),
})
```

### 3. RequestMyIncome

Requests invoices that have income characterization for the user for a given period.

```go
code, docs, err := c.RequestMyIncome(mydataInvoices.RequestMyIncomeParams{
    DateFrom: time.Now().UTC().AddDate(0, -12, 0),
    DateTo: time.Now().UTC(),
})
```

### 4. RequestMyExpenses

Requests invoices that have expense characterization for the user for a given period.

```go
code, docs, err := c.RequestMyExpenses(mydataInvoices.RequestMyExpensesParams{
    DateFrom: time.Now().UTC().AddDate(0, -3, 0),
    DateTo: time.Now().UTC(),
})
```

### 5. RequestTransmittedDocs

Fetches invoices that the user has submitted.

```go
code, docs, err := c.RequestTransmittedDocs(mydataInvoices.RequestDocsParams{
    Mark: "1",
})
```

There are some convenience functions defined

#### 5.1. RequestTransmittedDocsPastDays

Fetches transmitted documents for the past specified number of days.

```go
code, docs, err := c.RequestTransmittedDocsPastDays(1) // last 1 day
```

#### 5.2. RequestTransmittedDocWithMark

Fetches transmitted documents with a specific mark.

```go
code, docs, err := c.RequestTransmittedDocWithMark(1) // mark 1
```

#### 5.3. RequestTransmittedDocBetweenMarks

Fetches transmitted documents between two specified marks.

```go
code, docs, err := c.RequestTransmittedDocBetweenMarks(1, 5) // between marks 1 and 5
```

## Iterator Functions

Iterators allow for paginated requests, fetching data in chunks.

### 1. NewIncomeIterator

Creates an iterator for income documents.

```go
params := mydataInvoices.RequestMyIncomeParams{
    DateFrom: time.Now().UTC().AddDate(0, -12, 0),
    DateTo:   time.Now().UTC(),
}
iter := NewIncomeIterator(c, params)
for i, res := range iter.Iterate() {
    res.Print() // process each result
}
```

### 2. NewVatInfoIterator

Creates an iterator for VAT information.

```go
params := mydataInvoices.RequestVatInfoParams{
    DateFrom: time.Now().UTC().AddDate(0, -12, 0),
    DateTo:   time.Now().UTC(),
}
iter := NewVatInfoIterator(c, params)
for i, res := range iter.Iterate() {
    res.Print() // process each result
}
```

### 3. NewExpensesIterator

Creates an iterator for expenses documents.

```go
params := mydataInvoices.RequestMyExpensesParams{
    DateFrom: time.Now().UTC().AddDate(0, -12, 0),
    DateTo:   time.Now().UTC(),
}
iter := NewExpensesIterator(c, params)
for i, res := range iter.Iterate() {
    res.Print() // process each result
}
```

### 4. NewTransmittedDocsIterator

Creates an iterator for transmitted documents.

```go
params := mydataInvoices.RequestDocsParams{
    Mark:     "0",
    DateFrom: time.Now().UTC().AddDate(0, -3, 0),
    DateTo:   time.Now().UTC(),
}
iter := NewTransmittedDocsIterator(c, params)
for i, res := range iter.Iterate() {
    res.Print() // process each result
}
```

### 5. NewDocsIterator

Creates an iterator for requested documents.

```go
params := mydataInvoices.RequestDocsParams{
    Mark:     "0",
    DateFrom: time.Now().UTC().AddDate(0, -3, 0),
    DateTo:   time.Now().UTC(),
}
iter := NewDocsIterator(c, params)
for i, res := range iter.Iterate() {
    res.Print() // process each result
}
```

### 6. Pull iterator

For every iterator defined, except `iter.Iterate()`, that is used with a range function to iterate over the results
you can use the `iter.Pull()` to pull values from the iterator. For example you can iterate as follows

```go
iter := NewDocsIterator(c, params)
next, stop := iter.Pull()
defer stop()
for {
    v, ok := next()
    if !ok {
        return
    }

}

## Conclusion

This documentation provides a comprehensive overview of how to use the functions and iterators in the `mydata` package. Each function is designed to interact with the myDATA API, allowing users to retrieve various types of documents and information efficiently. Use the provided examples to guide your implementation.
```
