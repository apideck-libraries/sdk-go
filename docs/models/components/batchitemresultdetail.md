# BatchItemResultDetail

Contains parameter or domain specific information related to the error and why it occurred.


## Supported Types

### 

```go
batchItemResultDetail := components.CreateBatchItemResultDetailStr(string{/* values here */})
```

### Detail2

```go
batchItemResultDetail := components.CreateBatchItemResultDetailDetail2(components.Detail2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch batchItemResultDetail.Type {
	case components.BatchItemResultDetailTypeStr:
		// batchItemResultDetail.Str is populated
	case components.BatchItemResultDetailTypeDetail2:
		// batchItemResultDetail.Detail2 is populated
}
```
