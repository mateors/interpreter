# Interpreter | Token | Lexer | Parser

Students I am going to write a brand new interpreter in golang from scratch. If you really care about learning something new and from scratch then this is the place to start. You can start following my GitHub profile and get my project status on regular basis. I push almost every day and multiple times with detailed instructions which are helpful for learners.

## When you tokenize
* Everything (source code, your input) must be tokenize
* Each character must match a token
* Unsupported characters are usually tokenize as ILLEGAL

## Different types of expression

### Expressions involving prefix operators:
* -5
* !true
* !false

### Infix operators (or “binary operators”) Expressions:
* 5 + 5
* 5 - 5
* 5 / 5
* 5 * 5

### Comparison operators Expressions:
* foo == bar
* foo != bar
* foo < bar
* foo > bar

> We can use parentheses to group expressions and influence the order of evaluation:

### Grouped expressions:
* 5 * (5 + 5)
* ((5 + 5) * 5) * 5

### call expressions:
* add(2, 3)
* add(add(2, 3), add(5, 10))
* max(5, add(5, (5 * 5)))

### Identifiers are expressions too:
* foo * bar / foobar
* add(foo, bar)

### function literal is just the expression in the statement:
> `let add = fn(x, y) { return x + y };`


### function literal in place of an identifier:
* fn(x, y) { return x + y }(5, 5)
* (fn(x) { return x }(5) + 10 ) * 10

### if expressions:
* let result = if (10 > 5) { true } else { false };
* result // => true

Looking at all these different forms of expressions it becomes clear that we need a really good approach to parse them correctly and in an understandable and extendable way. And that is where `Vaughan Pratt` comes in.


## Pratt Parser

### Integer Literals
Integer literals are expressions. The value they produce is the integer itself.

imagine in which places integer literals can occur to understand why they are expressions:

* let x = 5;
* add(5, 10);
* 5 + 5 + 5;

### Prefix Operator | Prefix expressions
* -5;
* !foobar;
* 5 + -10;

#### Prefix operator structure:
> `<prefix operator><expression>;`

Any expression can follow a prefix operator as operand.
* `!isGreaterThanZero(2);`
* `5 + -add(5, 5);`


## HOW PRATT PARSING WORKS

> Statement : `1 + 2 * 3`

We are going take a close look at what the parser does as soon as `parseExpressionStatement` is called for the first time.

#### ParseProgram() | Here is what happens when we parse `1 + 2 * 3;`
    1. parseStatement()
        1. parseExpressionStatement()
            1. `parseExpression(LOWEST:1)`
                1. curToken= 1, peekToken= +, curToken.Type= INTEGAR
                2. leftExp= parseIntegerLiteral()
                    leftExp= &ast.IntegerLiteral{Token: {INTEGAR,1}, Value: 1}

                3. precedence=1, peekPrecedence=4, peekToken.Type= PLUS
                4. infix=parseInfixExpression
                5. NextToken>> curToken= +, peekTone= 2, curToken.Type= PLUS, peekToken.Type= INTEGAR
                6. leftExp= `parseInfixExpression(leftExp)`
                    1. leftExp= `&ast.InfixExpression{
                                    Token: {PLUS,+} 
                                    Operator: + 
                                    Left: &ast.IntegerLiteral{Token: {INTEGAR,1}, Value: 1} 
                                    Right: ?? 
                                }`
                    2. precedence=4
                    3. `NextToken>> curToken= 2, peekTone= *, curToken.Type= INTEGAR, peekToken.Type= ASTERISK`
                    4. `parseExpression(precedence:4)`
                        1. prefixFunction=parseIntegerLiteral
                        2. leftExp= parseIntegerLiteral() 
                        leftExp= &ast.IntegerLiteral{Token: {INTEGAR,2}, Value: 2}

                    5. precedence=4, peekPrecedence=5, peekToken.Type= ASTERISK
                        1. infix=parseInfixExpression
                        2. NextToken>> curToken= *, peekTone= 3, curToken.Type= ASTERISK, peekToken.Type= INTEGAR
                        3. leftExp= parseInfixExpression(`&ast.IntegerLiteral{Token: {INTEGAR,2}, Value: 2}`)
                            1. &ast.InfixExpression{Token: {ASTERISK,*} Operator: * Left: &ast.IntegerLiteral{Token: {INTEGAR,2}, Value: 2} Right: ?? }
                            2. precedence=5
                            3. NextToken>> curToken= 3, peekTone= , curToken.Type= INTEGAR, peekToken.Type= 
                            4. Right::parseExpression(precedence:5)
                                    1. prefixFunction=parseIntegerLiteral
                                    2. leftExp= parseIntegerLiteral()
                                       leftExp= &ast.IntegerLiteral{Token: {INTEGAR,3}, Value: 3}
                            5. Right= &ast.InfixExpression{
                                    Token: {ASTERISK,*} Operator: * 
                                    Left: &ast.IntegerLiteral{Token: {INTEGAR,2}, Value: 2} 
                                    Right: &ast.IntegerLiteral{Token: {INTEGAR,3}, Value: 3} 
                                }