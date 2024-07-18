# Assignment-week12
In this assignemnt the Go function Divide takes two integer parameters, a and b, returning the result of a divided by b. That is, this function does integer division explicitly, hence giving an integer result. 

TestDivide is a test for Divide function using the testing package. It defines a set of test cases with various input values a and b, with their expected results returned back from the Divide function. The test cases will cover all possibilities: positive, negative, and zero values for the variable types.
It calls, for each test case, the function Divide with the input values and checks the result against the expected value. If the result doesn't turn up as expected, it logs an error message using t.Errorf, mentioning what the actual result versus the expected result was.
