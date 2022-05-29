# Esper

A programming language narrowly designed for code sharing across environments.

## Pitch

So, you know how pretty much every modern, garbage-collected language feels eerily like the others lately? I think we can blame this phenomenon on the fact that many of these languages are converging on the same features and concepts; how many languages have added first-class functions, co-routines, data classes, and language-level immutablity recently? The only _tangible_ differences between one language and another are the ecosystems and platforms that they can unlock for you. Go gets you into the cloud and terminal ecosystems, while JS/TS gets you into the browser and to the edge. Swift and Kotlin get you onto phones, and with C# you can ship on an Xbox.

And it got me thinking: **if the languages we use to write our apps are this similar, why on earth are we writing the same logic over and over again?** Why can't we write most of our logic, constants, and types once, and use them anywhere? What if there was a language purely designed to be interoperable with other languages?

Esper is that language.

The intent behind Esper is to incorporate the smallest set of common features from these garbage-collected languages sufficient to:

- Create common types
- Implement business logic
- Declare common constants

Of course, it wouldn't hurt to end up with a language that is pleasant to use and maintain while we're at it.

## Design

Esper is not designed to be fast, or sexy, or interesting, or well-suited for any specific domain. It should fit right into the source coce powering any ol' user interface, backend API, and smart fridge. Esper's guiding design principles, are to be:

- Simple,
- Familar, and
- Practical

Esper should never feel as esoteric and ornate as Rust, but it should feel a smidge more expressive than Go. It should be easy to read, follow, and document like Java; getting out of your way and letting you solve the damn problem like Node.js.

Described above is language that will be difficult to design, and even harder to implement. My hope in all of this, at the very least, is to move the [Overton window](https://en.wikipedia.org/wiki/Overton_window) in a direction that we bet the programming world would enjoy.

Wish me luck.

### Lineage

As Esper is designed to feel familiar, it borrows heavily from some popular programming languages/runtimes:

- Dependency management from [Deno](https://deno.land/)
- Module system and batteries-included standard library championed by [Go](https://go.dev/)
- Syntax largely stolen from [Rust](https://www.rust-lang.org/) with a few tricks from [TypeScript](https://www.typescriptlang.org/) and [Python](https://www.python.org/) included
- Concurrency model inspired by [Dart](https://dart.dev/)
- Testing is an hommage to [Dart](https://dart.dev/), [JavaScript](https://www.javascript.com/), and [Go](https://go.dev/)

#### Syntax

There is always where you want to start with a new language - what will it look like?

```rust
// Hey look ma! Pythonic imports with Deno/Go-style dependency management.
// And yes, and you have probably gathered, Esper comments = Rust comments.

from "github.com/abc/xyz@v0.9.0" use { hello, world };
from "github.com/foo/bar@v1.2.0-beta/nested/module" use * as module;

// Relative imports are a thing too. Notice how the version isn't specified - this is
// because relatively imported modules always share the version of the importer.

from "some/sub/module" use { something };
from "../bing/bang" use { boom as büm };

// You can also re-export in a similar way.

from "github.com/abc/xyz@v0.9.0" show { hello };
from "github.com/foo/bar@v1.2.0-beta/nested/module" show *;

from "../bing/bang" show { boom as büm };

// Rusty variable declaration with type inference. That means
// `some_immutable_integer` and other variables declared like this cannot be mutated.
let some_immutable_integer = 12;

// As far as primitives go, we plan to support a few:
// - `bool`
// - `char`
// - `double`
// - `int` (no unsigned ints for now)
// - `string`
//
// Operators are fairly typical too:
// - `&&`, `||`, `==`, and `!=` all do what you think they do
// - `+`, `-`, `*`, `/`, and `%` all do what you think they do
//   (except that they only apply to numbers)
// - `**`, the exponentiation operator, is stolen from [ES2017](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Exponentiation)
// - `~/`, truncate divide, is stolen from [Dart](https://api.flutter.dev/flutter/dart-core/num/operator_truncate_divide.html)
//
// You might be wondering where all the bitwise operators are - there are none! lol.
// Good riddance.


// `mut` makes the declared variable mutable.
let mut some_mutable_double = -7.1e2;
some_mutable_double = some_mutable_double + 1;

// Like in Rust, you can export stuff with `pub`. Anything not declared with a `pub`
// will only be visible to things in its own module.
pub let stuff_outside_of_this_module_can_see_me = true;

/// Doc comments! I bet this looks familar to you Rustaceans out there.
///
/// # Markdown
///
/// Who *doesn't* _love_ some `sweet` [Markdown](https://www.markdownguide.org/)
/// formatting?
///
/// # Code links
///
/// All you have to do to reference somthing in your code is wrap it in `[]`.
/// For instance, [some_mutable_double] is a variable we defined above.
/// (This was 100% stolen from Dart).
let a_string = "this is a string";

// Rusty functions everybody!
// No need to specify if a function is `void`, just say nothing at all:

fn print_hello_world() {
  // By the way, printing works like it does in Dart.
  print("hello world");
}

// There is just one wrinkle with functions - Esper does not have positional
// arguments. Args must be labeled unless a variable is passed along sharing
// the name of an argument. There is one exception to this rule: if a function
// has just a single argument, no label is necessary:

fn multiply_by_two(num: double) -> double {
  num * 2
}

print(multiply_by_two(3)); // prints "6"

fn multiply(a: double, b: double) -> double {
  a * b
}

print(multiply(a: 2, b: 5)); // prints "10"

let b = 5;

print(multiply(a: 2, b)); // prints "10"

// Oh! And one more thing, all you have to make an argument optional is give
// it a default value:

fn i_cant_wait_to(action /* `: string` is inferred here */ = "take a nap") {
  // Now seems like a good time to mention we stole string interpolation from
  // Dart.
  print("Time to $action!");
}

print(i_cant_wait_to("eat donuts")); // prints "Time to eat donuts!"
print(i_cant_wait_to()); // prints "Time to take a nap!"

// Esper even borrows Rust's syntax for lambda expressions (Rust calls them closures):
let lambda_annotated = |i: i32| -> i32 { i + 1 };
let lambda_inferred  = |i     |          i + 1  ;

// You may now be wondering how more complex data structures are created and
// managed in Esper. I'm sure you are shocked to find out that we (mostly) stole
// Rust syntax here too /s.

struct User {
  active: bool;
  coolness_rating: int;
  email: string;
  username: string;
}

// We can make some or even all of the `struct` or its fields public using the
// `pub` keyowrd.

pub struct Company {
  pub name: string;
  pub phone_number: string;
  cash_in_the_bank: double;
}

// You may notice that we opted to go with `;` to terminate field declarations
// instead of `,`. This is mostly to make adding methods to structs feel more
// natural. Esper sugarifies Rust's default `impl` by simply allowing you to
// declare methods within the `struct` itself.

struct Car {
  make: string;
  model: string;
  owner: User;
  still_works: bool;
}

// One important thing to note here is that, in general, Esper data structures
// are immutable by default.

let my_car = Car {
  make: "Mazda",
  model: "Miata",
  owner: some_user,
};

my_car.make = "Toyota"; // Compile time error.

// To "change" an immutable value, we have to first clone it first as a mutable
// value.

let mut my_mut_car = my_car.mut;

my_mut_car.make = "Toyota"; // This is a-ok.

// Luckily, Esper has some syntactic sugar to make this look a little cleaner:

let my_first_car = my_car.mut {
  make: "Toyota",
  model: "Camry",
  year: 2008,
};

// Notice that `my_first_car` is mutable, but you may not always want this
// side-effect.

my_first_car.make = "Ferrari"; // This isn't quite right, but is not an error.

// Often times, mutation is needs to happen deeper in the the struct. Esper
// allows for this use case with some more syntax sugar.

let my_next_car = my_immutable_first_car.mut {
  make: "Rivian",
  model: "R1T",
  user: mut {
    coolness_rating: self.coolness_rating + 1,
  },
  year: 2023,
};

// Sometimes, you just gotta mutate structs directly. This is fairly simple to
// do in Esper. All you have to do in order to create a mutable `struct` is use
// the `mut` at creation time:

struct Donut {
  is_tasty: bool;
}

let disappointing = mut Donut { is_tasty: false };

disappointing.is_tasty = true; // This is a-ok.

```

#### Module system

As mentioned above, Esper modules work a lot like Go modules. Each directory, and all of the source files within it, act as a single module. This means that all source files in the same directory act sort of like one big source file. Additionally, from outside of an Esper module, there is not visibility into anything lacking a `pub` keyword.

TODO(skeswa): continue noodling

## Prototype

I think it might be a good idea to check out something like [lalrpop](http://lalrpop.github.io/lalrpop/).
