package vm

import (
	"fmt"
	"monkey/ast"
	"monkey/compiler"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

type vmTestCase struct {
	input    string
	expected interface{}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"50 / 2 * 2 + 10 - 5 * 2", 50},
		{"5 * (2 + 10)", 60},
		{"-5", -5},
		{"-50 + 100 + -50", 0},
		{"(5 + 10 * 2 + 15 / 3)", 30},
	}
	runVmTests(t, tests)
}

func TestBooleanExpressions(t *testing.T) {
	tests := []vmTestCase{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"!true", false},
		{"!false", true},
		{"!5", false}, // In our specification, everything other than False are treated as truthy.
		{"!!5", true},
		{"!(if(false){ 5; })", true},                // To test !Null to be treated as true
		{"if (if(false){ 5; }) {10} else {20}", 20}, // To test nil to be treated as false

	}
	runVmTests(t, tests)
}

func TestConditionals(t *testing.T) {
	tests := []vmTestCase{
		{"if (true) { 10 }", 10},
		{"if (true) { 10 } else { 20 }", 10},
		{"if (false) { 10 } else { 20 }", 20},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 > 2) { 10 }", Null},
		{"if (false) { 10 }", Null},
	}
	runVmTests(t, tests)
}

func TestGlobalLetStatements(t *testing.T) {
	tests := []vmTestCase{
		{"let one = 1; one", 1},
		{"let one = 1; let two = 2; one + two", 3},
		{"let one = 1; let two = one + one; one + two", 3},
	}
	runVmTests(t, tests)
}

func TestStringExpressions(t *testing.T) {
	tests := []vmTestCase{
		{`"monkey"`, "monkey"},
		{`"mon" + "key"`, "monkey"},
		{`"mon" + "key" + "banana"`, "monkeybanana"},
	}
	runVmTests(t, tests)
}

func TestArrayLiterals(t *testing.T) {
	tests := []vmTestCase{
		{"[]", []int{}},
		{"[1,2,3]", []int{1, 2, 3}},
		{"[1+2, 3*4, 5+6]", []int{3, 12, 11}},
	}
	runVmTests(t, tests)
}

func TestHashLiterals(t *testing.T) {
	tests := []vmTestCase{
		{
			"{}", map[object.HashKey]int64{},
		},
		{
			"{1: 2, 2: 3}",
			map[object.HashKey]int64{
				(&object.Integer{Value: 1}).HashKey(): 2,
				(&object.Integer{Value: 2}).HashKey(): 3,
			},
		},
		{
			"{1+1: 2*2, 3+3: 4*4}",
			map[object.HashKey]int64{
				(&object.Integer{Value: 2}).HashKey(): 4,
				(&object.Integer{Value: 6}).HashKey(): 16,
			},
		},
	}
	runVmTests(t, tests)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []vmTestCase{
		{"[1,2,3][1]", 2},
		{"[1,2,3][0+2]", 3},
		{"[[1,1,1]][0][0]", 1},
		{"[][0]", Null},
		{"[1,2,3][3]", Null},
		{"[1][-1]", Null},
	}
	runVmTests(t, tests)
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []vmTestCase{
		{"{1:1, 2:2}[1]", 1},
		{"{1:1, 2:2}[2]", 2},
		{"{1: 1}[0]", Null},
		{"{}[0]", Null},
	}
	runVmTests(t, tests)
}

func TestCallingFunctionsWithoutArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let fivePlusTen = fn() {5+10;};
			fivePlusTen();
			`,
			expected: 15,
		},
		{
			input: `
			let one = fn(){1;};
			let two = fn(){2;};
			one() + two()
			`,
			expected: 3,
		},
		{
			input: `
			let a = fn(){1};
			let b = fn(){a()+1};
			let c = fn(){b()+1};
			c();
			`,
			expected: 3,
		},
	}
	runVmTests(t, tests)
}

func TestCallingFunctionsWithReturnStatements(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let earlyExit = fn(){return 99; 100;};
			earlyExit()
			`,
			expected: 99,
		},
		{
			input: `
			let earlyExit = fn(){return 99; return 100;};
			earlyExit()
			`,
			expected: 99,
		},
	}
	runVmTests(t, tests)
}

func TestCallingFunctionsWithoutReturnValue(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let noReturn = fn(){};
			noReturn()
			`,
			expected: Null,
		},
		{
			input: `
			let noReturn = fn(){};
			let noReturnTwo = fn(){ noReturn() };
			noReturn();
			noReturnTwo();
			`,
			expected: Null,
		},
	}
	runVmTests(t, tests)
}

func TestCallingFunctionsWithBindings(t *testing.T) {
	tests := []vmTestCase{
		{
			// Test that local binding works in a first place.
			input: `
			let one = fn() { let one = 1; one };
			one();
			`,
			expected: 1,
		},
		{
			// Test multiple local bindings in one function.
			input: `
			let oneAndTwo = fn(){ let one = 1; let two = 2; one + two };
			oneAndTwo()
			`,
			expected: 3,
		},
		{
			// Test multiple bindings in different functions.
			input: `
			let oneAndTwo = fn(){ let one = 1; let two = 2; one + two };
			let threeAndFour = fn(){ let three = 3; let four = 4; three + four };
			oneAndTwo() + threeAndFour();
			`,
			expected: 10,
		},
		{
			// Test that same named local bindings do not affect each other.
			input: `
			let firstFooBar = fn(){ let FooBar = 50; FooBar; };
			let secondFooBar = fn(){ let FooBar = 100; FooBar; };
			firstFooBar() + secondFooBar();
			`,
			expected: 150,
		},
		{
			input: `
			let globalSeed = 50;
			let minusOne = fn(){
				let num = 1;
				globalSeed - num;
			}
			let minusTwo = fn(){
				let num = 2;
				globalSeed - num;
			}
			minusOne() + minusTwo()
			`,
			expected: 97,
		},
	}
	runVmTests(t, tests)
}

func TestCallingFunctionsWithArgumentsAndBindings(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let identity = fn(a){ a; }
			identity(4)
			`,
			expected: 4,
		},
		{
			input: `
			let sum = fn(a, b){ a + b; };
			sum(1, 2);
			`,
			expected: 3,
		},
		{
			input: `
			let globalNum = 10;

			let sum = fn(a, b){ let c = a + b; c + globalNum; };
			let outer = fn() { sum(1, 2) + sum(3, 4) + globalNum; };
			outer() + globalNum;
			`,
			expected: 50,
		},
	}
	runVmTests(t, tests)
}

func TestCallingFunctionsWithWrongNumberOfArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			input:    `fn() { 1; }(1);`,
			expected: `wrong number of arguments: want=0, got=1`,
		},
		{
			input:    `fn(a) { a; }();`,
			expected: `wrong number of arguments: want=1, got=0`,
		},
		{
			input:    `fn(a, b) { a + b; }(1);`,
			expected: `wrong number of arguments: want=2, got=1`,
		},
	}
	for _, tt := range tests {
		program := parse(tt.input)
		comp := compiler.New()

		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.ByteCode())
		err = vm.Run()

		if err == nil {
			t.Fatalf("expected VM error but resulted in none.")
		}

		if err.Error() != tt.expected {
			t.Fatalf("wrong VM error: want=%q, got=%q", tt.expected, err)
		}
	}
}

func TestFirstClassFunctions(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let returnsOne = fn(){1;};
			let returnsOneReturner = fn(){ returnsOne; };
			returnsOneReturner()();
			`,
			expected: 1,
		},
	}
	runVmTests(t, tests)
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []vmTestCase{
		// len
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{
			`len(1)`,
			&object.Error{Message: "argument to `len` not supported, got=INTEGER"},
		},
		{`len([1,2,3])`, 3},
		{`len([])`, 0},
		// puts
		{`puts("hello", "world!")`, Null},
		// first
		{`first([1,2,3])`, 1},
		{`first([])`, Null},
		{`first(1)`, &object.Error{Message: "argument to `first` must be ARRAY, got: INTEGER"}},
		// last
		{`last([1,2,3])`, 3},
		{`last([])`, Null},
		{`last(1)`, &object.Error{Message: "argument to `last` must be ARRAY, got: INTEGER"}},
		// rest
		{`rest([1,2,3])`, []int{2, 3}}, // これ
		{`rest([])`, Null},
		// push
		{`push([], 1)`, []int{1}},
		{`push(1,1)`, &object.Error{Message: "argument to `push` must be ARRAY, got=INTEGER"}},
	}
	runVmTests(t, tests)
}

func TestClosures(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let newClosure = fn(a){ fn(){a;}; }
			let closure = newClosure(99)
			closure()`,
			expected: 99,
		},
		{
			input: `
			let newAddr = fn(a, b){ fn(c){ return a+b+c }; }
			let adder = newAddr(1, 2)
			adder(8)`,
			expected: 11,
		},
		{
			input: `
			let newAdder = fn(a, b){
				let c = a + b;
				fn(d) { c + d };
			}
			let adder = newAdder(1, 2)
			adder(8)`,
			expected: 11,
		},
	}
	runVmTests(t, tests)
}

func TestRecursiveFunctions(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let countDown = fn(x) {
				if (x == 0) {
					return 0;
				} else {
					countDown(x-1)
				}
			};
			countDown(1);`,
			expected: 0,
		},
	}
	runVmTests(t, tests)
}

/*
Tests the top element in the stack.
*/
func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}
		// Assume that the logic till here (till compiling) is right.

		vm := New(comp.ByteCode())
		err = vm.Run()
		if err != nil {
			t.Fatalf("vm error: %s", err)
		}
		stackElem := vm.LastPoppedStackElem()
		testExpectedObject(t, tt.expected, stackElem)
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testExpectedObject(t *testing.T, expected interface{}, actual object.Object) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	case string:
		err := testStringObject(expected, actual)
		if err != nil {
			t.Errorf("testStringObject failed: %s", err)
		}
	case bool:
		err := testBooleanObject(expected, actual)
		if err != nil {
			t.Errorf("testBooleanObject failed: %s", err)
		}
	case *object.Null:
		if actual != Null {
			t.Errorf("object is not Null: %T (%+v)", actual, actual)
		}
	case []int:
		array, ok := actual.(*object.Array)
		if !ok {
			t.Errorf("object not Array: %T (%+v)", actual, actual)
			return
		}

		if len(array.Elements) != len(expected) {
			t.Errorf(
				"wrong num of elements. want=%d, got=%d",
				len(expected), len(array.Elements),
			)
			return
		}

		for i, expectedElem := range expected {
			err := testIntegerObject(int64(expectedElem), array.Elements[i])
			if err != nil {
				t.Errorf("testIntegerObject failed %s", err)
			}
		}
	case map[object.HashKey]int64:
		hash, ok := actual.(*object.Hash)
		if !ok {
			t.Errorf("object is not Hash. got=%T(%+v)", actual, actual)
			return
		}
		if len(hash.Pairs) != len(expected) {
			t.Errorf(
				"hash has wrong number of Pairs. want=%d, got=%d",
				len(expected), len(hash.Pairs),
			)
		}
		for expectedKey, expectedValue := range expected {
			pair, ok := hash.Pairs[expectedKey]
			if !ok {
				t.Errorf("no pair for given key in Pairs.")
			}
			err := testIntegerObject(expectedValue, pair.Value)
			if err != nil {
				t.Errorf("testIntegerObject failed: %s", err)
			}
		}
	case *object.Error:
		errObj, ok := actual.(*object.Error)
		if !ok {
			t.Errorf("object is not Error: %T (%+v)", actual, actual)
			return
		}
		if errObj.Message != expected.Message {
			t.Errorf(
				"wrong error message. expected=%q, got=%q",
				expected.Message, errObj.Message,
			)
		}

	}
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf(
			"object is not Integer. got=%T(%+v)",
			actual, actual,
		)
	}

	if result.Value != expected {
		return fmt.Errorf(
			"object has wrong value. want=%d, got=%d",
			expected, result.Value,
		)
	}
	return nil
}

func testStringObject(expected string, actual object.Object) error {
	result, ok := actual.(*object.String)
	if !ok {
		return fmt.Errorf(
			"object is not String. got=%T(%+v)",
			actual, actual,
		)
	}

	if result.Value != expected {
		return fmt.Errorf(
			"object has wrong value. want=%s, got=%s",
			expected, result.Value,
		)
	}
	return nil
}

func testBooleanObject(expected bool, actual object.Object) error {
	result, ok := actual.(*object.Boolean)
	if !ok {
		return fmt.Errorf(
			"object is not Boolean. got=%T(%+v)",
			actual, actual,
		)
	}

	if result.Value != expected {
		return fmt.Errorf(
			"object has wrong value. want=%t, got=%t",
			expected, result.Value,
		)
	}
	return nil
}
