# Value5


## Supported Types

### 

```go
value5 := components.CreateValue5Str(string{/* values here */})
```

### 

```go
value5 := components.CreateValue5Integer(int64{/* values here */})
```

### 

```go
value5 := components.CreateValue5Number(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value5.Type {
	case components.Value5TypeStr:
		// value5.Str is populated
	case components.Value5TypeInteger:
		// value5.Integer is populated
	case components.Value5TypeNumber:
		// value5.Number is populated
}
```
