// slice as stack
var excepted = map[byte]byte {
	')' : '(',
	']' : '[',
	'}' : '{',
}
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])			// push
		default:
			n := len(stack)
			if n > 0 && stack[n-1] == excepted[s[i]] {
				stack = stack[:len(stack)-1]	// pop
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}


// my stack (and error)
var excepted = map[byte]byte {
	')' : '(',
	']' : '[',
	'}' : '{',
}

const INIT_STACK_SIZE = 10
const INCREASE_STACK_SIZE =  10

type Stack struct {
	base []byte
	cap int
	top int
}

func createStack() Stack {
	s := new(Stack)
	s.base = make([]byte, INIT_STACK_SIZE, INIT_STACK_SIZE)
	s.cap = INIT_STACK_SIZE
	s.top = -1
	return *s
}

func (s *Stack) pop() (byte, error) {
	if s.top < 0 {
		return '0', newError(-1, "cannot pop")
	}
	s.top--
	return s.base[s.top+1], nil
}

func (s *Stack) push(ch byte) {
	if s.top == s.cap - 1 {
		newBase := make([]byte, s.cap + INCREASE_STACK_SIZE, s.cap + INCREASE_STACK_SIZE)
		copy(newBase, s.base)
		s.base = newBase
		s.cap += INCREASE_STACK_SIZE
	}
	s.top++
	s.base[s.top] = ch
}

func (s Stack) peek() (byte, error) {
	if s.top < 0 {
		return '0', newError(-2, "cannot peek")
	}
	return s.base[s.top], nil
}

func (s Stack) isEmpty() bool {
	return s.top == -1
}

type Error struct {
	errCode int
	errMsg string
}

func (err *Error) Error() string {
	return fmt.Sprintf( "stack error %v: " + err.errMsg, err.errCode)
}

func newError(code int, msg string) *Error {
	return &Error{code, msg}
}

func isValid(s string) bool {
	stack := createStack()
	for i := range s {
		switch s[i] == ')' || s[i] == ']' || s[i] == '}' {
		case true:
			if ch, err := stack.peek(); err == nil && ch == excepted[s[i]] {
				_, _ = stack.pop()
			} else {
				return false
			}
		default:
			stack.push(s[i])
		}
	}
	return stack.isEmpty()
}