# minimal-secret-santa-go

GO implementation of a command line secret-santa with **ChatGPT**

[article on my website](https://creativeyann17.github.io/#/article/minimal-secret-santa-go)

## Configuration
The program requires a TSV file as input, example:

|NAME|EXCLUDE|
|-|-|
a|b c
b|a
c
d
e

*Note: EXCLUDE are people you don't want to give gifts to each other. It can contain multiple elements, use white space as separator.*
