However, there is one thing to be wary of when determining whether you should utilize a Pool: 
if the code that utilizes the Pool requires things that are not roughly homogenous, 
you may spend more time converting what you’ve retrieved from the Pool than it would have taken to just instantiate it in the first place.

For instance, if your program requires slices of random and variable length, a Pool isn’t going to help you much. 
The probability that you’ll receive a slice the length you require is low.


When working with a Pool, just remember the following points:
• When instantiating sync.Pool, give it a New member variable that is thread-safe when called.
• When you receive an instance from Get, make no assumptions regarding the state of the object you receive back.
• Make sure to call Put when you’re finished with the object you pulled out of the pool. Otherwise, the Pool is useless. Usually this is done with defer.
• Objects in the pool must be roughly uniform in makeup.
