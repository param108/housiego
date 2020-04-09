# housiego

code to generate housie tickets (its not as easy as you thing)

# setup
`go build`

# run

`./housiego name:number name:number`

name is the name of the person
number is the number of tickets he/she wants

eg
./housiego param:1 manan:2

output:

```
param:
| 5|11|21|  |42|  |63|  |  |
| 7|  |  |33|  |52|70|77|  |
| 8|  |  |39|47|55|  |  |88|


manan:
| 7|20|  |32|43|  |  |  |82|
|  |  |26|  |48|53|  |71|85|
|  |  |28|39|50|  |61|73|  |


|  |11|24|33|  |  |  |73|82|
| 7|  |  |  |42|  |69|74|87|
|10|  |  |40|  |58|70|77|  |
```
