## Basic working of a lexer
Here’s an example. This is the input one gives to a lexer:
"let x = 5 + 5;"
And what comes out of the lexer looks kinda like this:
[  
	LET,
     IDENTIFIER("x"),
     EQUAL_SIGN,
     INTEGER(5),
     PLUS_SIGN,
     INTEGER(5),
     SEMICOLON
     
]

## How we achieved functionality?
The isLetter helper function just checks whether the given argument is a letter. That sounds easy enough, but what’s noteworthy about isLetter is that changing this function has a larger impact on the language our interpreter will be able to parse than one would expect from such a small function. As you can see, in our case it contains the check ch == '_', which means that we’ll treat _ as a letter and allow it in identifiers and keywords. That means we can use variable names like foo_bar. Other programming languages even allow ! and ? in identifiers. If you want to allow that too, this is the place to sneak it in.

In the default: branch of the switch statement we use readIdentifier() to set the Literal field of our current token. But what about its Type? Now that we have read identifiers like let, fn or foobar, we need to be able to tell user-defined identifiers apart from language keywords. We need a function that returns the correct TokenType for the token literal we have. What better place than the token package to add such a function?
