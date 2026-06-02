type A = keyof any;
type B = string;
type C = number;
type D = Exclude<B, C>;
type E = Extract<B, C>;
type F = string;
type G = Exclude<B, F>;

type N = NonNullable<string | null | undefined>;

const a: unknown = 'hello';
let b: unknown;
b = a;

let c: any;
c = a;

let d: unknown;
d = c;

let e: any;
e = c;
