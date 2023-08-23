# Running Instructions:


To launch the engine; Simply run the following snippet:

```
go run cmd/main.go
```

# Documentation

## Device Management

Events are distributed and managed at a device level and is thus, truly event-driven.
The device manager is responsible for managing the devices and their respective events.
`client.Run` launches all listeners on a device level. This is easier for maintenance and debugging.
Additionally, it clearly draws the distinction between engine-specific callback listeners and server-specific actions.

## Essence

Every time a source code is compiled, the analyser generates an essence. This is a condensed version of the source code that contains only the information that is relevant to the analyser. 
This essence is then passed through the generator to attach the functional methods to it. 
The idea is that, when the engine has an appropriate memory manager capable of caching and a graph comparison method, we can reuse the same essence without having to recompile the code for re-running.

## Lookup 

Presently, two individual lookup tables are constructed to generate the code in code generator. 
- The actions lookup table generates action methods.
- The callbacks lookup table generates callback methods.