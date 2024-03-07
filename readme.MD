# tra-slovnik-sk-go

slovnik.sk from command line

## Description
Uses webpage slovnik.sk to translate between slovak and other languages:  
EN, ES, IT, FR, DE, HU, RU

## Usage
tra [-limit INT] [-l EN|DE|ES|IT|FR|HU|RU] INPUT WORD"  
	Arguments:"  
		-limit | --limit (INT): limit results  
		-l | --language (STR): translation language, default EN, available: EN|DE|ES|IT|FR|HU|RU"

### Example of use
open terminal
write and confirm your query `tra plethora` or `tra bahno`  
choose language to use `tra -l es porque`  
choose limit of results if word is too general `tra -limit 3 what`

### get result
```
> tra plethora
# English -> Slovak (plethora)
plethora: nadbytok červených krviniek, nadbytok

> tra -limit 1 bahno
# Slovak -> English (bahno)
bahno: dirt bed, clarts, scurf, sleech, slop, morass, bog, pulp, mud, scum
```

## Image
![screenshot](./scrn.png)