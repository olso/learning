/* A comment */

Js.log("Hello, BuckleScript and Reason!");

Js.log("🔥");
Js.log({js|❤️|js});

let myInt: float = 123.0;

let add = (x, y) => x + y;

type color =
    Red |
    Green |
    Blue;

let myBool = true;

Js.log(if (myBool) "yes" else "no");

/* let abcabc = {
    let abc = "abc";

    abc ++ abc;
} */

let greeting = "hello!";
let someChar = 'H';

/* if (greeting === String) {
} */
/* 
let immuTest = "hello!";
immuTest = "hello2!";
let immuTest = "hello3!"; */

let message: string = {
    let part1 = "hello";
    let part2 = "world";

    part1 ++ " " ++ part2
}
