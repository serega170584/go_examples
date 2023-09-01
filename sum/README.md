
#Algorithm

Firstly we need order and work strategy when given values have different types
It seems obvious that if we take int64 and float64 values the sum has to be casted to float64 becauseof more wide range of complex
If we take float64 and complex the sum has to be casted to complex becauseof more wide range of complex
It turns out that int64 has to be casted to complex
Now we can consider all the cases and approach to cover these ones
    
- case 1 int + int : return int
- case 2 int + float : return float
- case 3 int + complex : return complex
- case 4 float + int : return float
- case 5 float + float : return float
- case 6 float + complex : return complex
- case 7 complex + int : return complex
- case 8 complex + float : return complex
- case 9 complex + complex : return complex

We pass list of numbers to the method New of sum struct for sum all these ones

In the method we iterate value by value

For resolving we use dynamic programming

We need initialize sum first state before iteration of numbers because we have to know what value zero number of list adds to

We easily assign zero-value int val for first state of sum

Next in the iteration of list we should cover all the cases above

We have 2 components that we need to sum: 1 - sum of previous numbers and 2 - current value

We need fix type of component 1 and type of component 2

Let's try to build-up conditions to cover cases above

If component 2 is complex and component 1 is not complex we easily can cast component 1 to float, pass it to complex function and add to component 2

So we cover case 3, 6

If component 2 is complex and component 1 is complex obviously we cover case 9

If component 1 is complex and component 2 is float we cover case 8

After get worked conditions above we can easily check if component 1 is complex so we cover case 7

After get worked conditions above we check if component 1 and component 2 are floats and sum these one in this case. Therefore we cover case 5

After get worked conditions above we check component 1 is float so we cover case 4

After get worked conditions above we check component 1 is int and component 2 is float so we cover case 2

Case 1 is handled when we couldn't visit conditions above

So we covered all the cases

It depends on type of sum whether we fill intVal, floatVal or complexVal in type sum

Method add check if intVal, floatVal and complexVal not zero null val and casts argument in terms of situation.