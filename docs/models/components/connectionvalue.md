# ConnectionValue


## Supported Types

### 

```go
connectionValue := components.CreateConnectionValueStr(string{/* values here */})
```

### 

```go
connectionValue := components.CreateConnectionValueInteger(int64{/* values here */})
```

### 

```go
connectionValue := components.CreateConnectionValueNumber(float64{/* values here */})
```

### 

```go
connectionValue := components.CreateConnectionValueBoolean(bool{/* values here */})
```

### 

```go
connectionValue := components.CreateConnectionValueArrayOfValue5([]components.Value5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch connectionValue.Type {
	case components.ConnectionValueTypeStr:
		// connectionValue.Str is populated
	case components.ConnectionValueTypeInteger:
		// connectionValue.Integer is populated
	case components.ConnectionValueTypeNumber:
		// connectionValue.Number is populated
	case components.ConnectionValueTypeBoolean:
		// connectionValue.Boolean is populated
	case components.ConnectionValueTypeArrayOfValue5:
		// connectionValue.ArrayOfValue5 is populated
}
```
