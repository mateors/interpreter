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