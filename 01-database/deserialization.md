#The problem
How to deserialize the messages until the last one ?


#The solution
Store the number of bytes used to store the message, just before the message data.
(Encode the length of the message in the data)