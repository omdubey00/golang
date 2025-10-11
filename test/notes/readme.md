### This one is for the TDD book to learn go lang and here are the all the notes for it . 

## Intro 

- Refactoring is the process of converting  one form of program to another form  of program without changing the underlying logic of the 
  program . Think it like writing tedious code into an efficient one without changing the method as you will have to 
  change your tests also for the solution --> This is refactoring means you dont change the underlying logic . 
  The working still remains the same. 

- We need testing to be able to check whether I have refactored my code in the right way or I have changed it's working. 
- That is why we will learn go by writing tests along the way so we understand how can we write generalized test for the logic even 
  if we refactor the code. That means I dont have to refactor the tests in order to test the current code 

# Advice --> Always write the smaller version of your project rather than spending time in creating the perfect one 
             that might eventually fail . 

#### Steps for doing test driven developmnet . 

- Write a small test for a small amount of desired behaviour
- Check the test fails with a clear error (red)
- Write the minimal amount of code to make the test pass (green)
- Refactor
- Repeat  
