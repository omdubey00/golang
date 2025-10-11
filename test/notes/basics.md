## Understanding a go program structuring . 

###  Writing a test is just like writing a function, with a few rules
### This is the important mindset write test , fail them , correct your code , make test pass , refactor it to improve it maintaining the passing of tests. 
## %v is the default placeholder and can be used for arrays. 

- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T ( this t is like a hook to our testing framework and lets us 
~ do many things like t.fail or t.Errorf that we used to display our error message if test fails) 

######  A few new concepts:

- In our function signature we have made a named return value (prefix
- string).
- This will create a variable called prefix in your function.
- It will be assigned the "zero" value. This depends on the type, for example
- ints are 0 and for strings it is "".
- You can return whatever it's set to by just calling return rather than
- return prefix.
- This will display in the Go Doc for your function so it can make the intent
- of your code clearer.
- default in the switch case will be branched to if none of the other case
- statements match.
- The function name starts with a lowercase letter. In Go public functions
- start with a capital letter and private ones start with a lowercase. We don't
- want the internals of our algorithm to be exposed to the world, so we made
- this function private.
