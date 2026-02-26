# UpdateConsentRequestResources


## Supported Types

### 

```go
updateConsentRequestResources := components.CreateUpdateConsentRequestResourcesMapOfMapOf1(map[string]map[string]components.One{/* values here */})
```

### Two

```go
updateConsentRequestResources := components.CreateUpdateConsentRequestResourcesTwo(components.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateConsentRequestResources.Type {
	case components.UpdateConsentRequestResourcesTypeMapOfMapOf1:
		// updateConsentRequestResources.MapOfMapOf1 is populated
	case components.UpdateConsentRequestResourcesTypeTwo:
		// updateConsentRequestResources.Two is populated
}
```
