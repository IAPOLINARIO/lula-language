# Lula Language

![lula_language_logo](logo.jpg)

[sarcasm]

Lula programming language was created in order to provide the easiest, faster and a reliable way for developers to create their beloved applications. You will never get disappointed with Lula, I can assure you that. 
If you are tired of all the regular programming languages out there, come and join the LULA community and let's spread the knowledge among those who need. Knowledge must be free and reach all, mainly those in need. 
  
PS - Lula is still being developed, so most of the features bellow are not ready yet. If you want to keep up to date with Lula, just star the project. 

[sarcasm /]

## Usage

Install the Lula using `go get`:

```sh
$ go get -v -u github.com/IAPOLINARIO/lula-language/...
```

Then run REPL:

```sh
$ $GOPATH/bin/lula-language
This is the Lula language!
Feel free to type in commands
>> 
```

Or run a Lula script file (for example `script.lula` file):

```sh
$ $GOPATH/bin/lula-language script.lula
```

## Getting started with Lula

### Variable bindings and number types

You can define variables using `let` keyword. Supported number types are integers and floating-point numbers.

```sh
>> pt a = 1;
>> a
1
>> pt b = 0.5;
>> b
0.5
```

### Arithmetic expressions

You can do usual arithmetic operations against numbers, such as `+`, `-`, `*` and `/`. 

```sh
>> pt a = 10;
>> pt b = a * 2;
>> (a + b) / 2 - 3;
12
>> pt c = 2.5;
>> b + c
22.5
```

### If expressions

You can use `if` and `else` keywords for conditional expressions. The last value in an executed block are returned from the expression.

```sh
>> pt a = 10;
>> pt b = a * 2;
>> pt c = if (b > a) { 99 } else { 100 };
>> c
99
```

### Functions and closures

You can define functions using `companheiro` keyword. All functions are closures in Lula and you must use `pt` along with `companheiro` to bind a closure to a variable. Closures enclose an environment where they are defined, and are evaluated in *the* environment when called. The last value in an executed function body are returned as a return value.

```sh
>> pt multiply = companheiro(x, y) { x * y };
>> multiply(50 / 2, 1 * 2)
50
>> companheiro(x) { x + 10 }(10)
20
>> pt newAdder = companheiro(x) { companheiro(y) { x + y }; };
>> pt addTwo = newAdder(2);
>> addTwo(3);
5
>> pt sub = companheiro(a, b) { a - b };
>> pt applyFunc = companheiro(a, b, companheiro) { companheiro(a, b) };
>> applyFunc(10, 2, sub);
8
```

### Strings

You can build strings using a pair of double quotes `""`. Strings are immutable values just like numbers. You can concatenate strings with `+` operator.

```sh
>> pt makeGreeter = companheiro(greeting) { companheiro(name) { greeting + " " + name + "!" } };
>> pt hello = makeGreeter("Hello");
>> hello("John");
Hello John!
```

#####
Needless to say, but this is just a joke and a way I found to learn GO in a fun way. Don't take it serious, because I don't ;) 
