# Five


## Supported Types

### 

```go
five := components.CreateFiveStr(string{/* values here */})
```

### 

```go
five := components.CreateFiveInteger(int64{/* values here */})
```

### 

```go
five := components.CreateFiveNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch five.Type {
	case components.FiveTypeStr:
		// five.Str is populated
	case components.FiveTypeInteger:
		// five.Integer is populated
	case components.FiveTypeNumber:
		// five.Number is populated
}
```
