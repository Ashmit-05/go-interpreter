## Wikipedia definition of a parser
A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree, abstract syntax tree or other hierarchical structure – giving a structural representation of the input, checking for correct syntax in the process. The parser is often preceded by a separate lexical analyser, which creates tokens from the sequence of input characters;

## Comparison with JSON
```
> var input = '{"name": "Thorsten", "age": 28}'; > var output = JSON.parse(input);
> output
{ name: 'Thorsten', age: 28 }
> output.name
'Thorsten'
> output.age
28
>
```
A JSON parser takes text as input and builds a data structure that represents the input. That’s exactly what the parser of a programming language does. The difference is that in the case of a JSON parser you can see the data structure when looking at the input. Whereas if you look at this :
```
if ((5 + 2 * 3) == 91) { return computeStuff(input1, input2); }
```

In most interpreters and compilers the data structure used for the internal representation of
the source code is called a “syntax tree” or an “abstract syntax tree” (AST for short). The “abstract” is based on the fact that certain details visible in the source code are omitted in the AST. Semicolons, newlines, whitespace, comments, braces, bracket and parentheses – depending on the language and the parser these details are not represented in the AST, but merely guide
the parser when constructing it.
A fact to note is that there is not one true, universal AST format that’s used by every parser. Their implementations are all pretty similar, the concept is the same, but they differ in details. The concrete implementation depends on the programming language being parsed.

## How AST's look
Code :
```
if (3 * 5 > 10) { 
    return "hello";
} else {
    return "goodbye";
}
```
AST :
```
> var input = 'if (3 * 5 > 10) { return "hello"; } else { return "goodbye"; }'; 
> var tokens = MagicLexer.parse(input);
> MagicParser.parse(tokens);
{
    type: "if-statement",
    condition: {
    type: "operator-expression",
    operator: ">",
    left: {
        type: "operator-expression",
        operator: "*",
        left: { type: "integer-literal", value: 3 },
        right: { type: "integer-literal", value: 5 }
    },
    right: { type: "integer-literal", value: 10 }
    },
    consequence: {
        type: "return-statement",
        returnValue: { type: "string-literal", value: "hello" }
    },
    alternative: {
        type: "return-statement",
        returnValue: { type: "string-literal", value: "goodbye" }
    }
}
```

## Note on helper functions
The `expectPeek` method is one of the "assertion functions" nearly all parsers share. Their primary purpose is to enforce the correctness of the order of the tokens by checking the type of the next token. Our `expectPeek` here checks the type of the `peekToken` and only if the type is correct does it advance the tokens by calling `nextToken`.
If it does not get expected type, we return nil, which gets ignored in `ParseProgram`, which results in entire statements being ignored becuase of an error in the input.
