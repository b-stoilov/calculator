# What does the app do?
The calculator app is a simple web service takes an input in the form of a string. The string is an expression in the form of : 
"What is ***{number} {operation} {number}***?" where operation could be one of the following ***\{"plus", "minus", "multiplied by", "divided by"\}***.
The input can take also even more numbers and operations. It then can either evaulate the expression or validate it.
- If the expression is evaluated, then the output should be displayed, e.g.:
> What is 5 multiplied by 5 plus 2? -> 27</li>
- If the expression is validated, then it should return information if the expression is valid (the correct syntax and correct operations used). If not, then it should display the type of failed validation.
The types of failed validation can be one of the following:
  - Unsupported operations ("What is 52 cubed?")</li>
  - Non-math questions ("Who is the President of the United States?")
  - Expressions with invalid syntax ("What is 1 plus plus 2?")
- If there are any errors during the validation command, save them and display them (only in-memory store)
    
# Rules of calculation
Since these are verbal word problems, the evaluation of the expression is done from left-to-right, ignoring the typical order of operations. (See the above example)

# How to run the app? (Windows)
THe app has a server side part and a CLI that acts as a client. Follow these steps to start the CLI:
1. in the /server/ directory ->  build the the server module into an .exe file called `server.exe`
2. Go to the /cli/ directory and build the module in to `calc.exe`
3. Start the calc CLI by typing `./calc`
   - to start the server, type `./calc start`. You should get a log when this step is successful
   - to evaluate an expression: `./calc evaluate {expression}`
   - to validate an expression: `./calc validate {expression}`
   - to fetch the errors, caught during the validation operation: `./calc errors` 
