Help from: https://github.com/NaddiNadja/grpc101

a) What are packages in your implementation? 
a2)What data structure do you use to transmit data and meta-data?



b) Does your implementation use threads or processes? 
b2) Why is it not realistic to use threads?
A thread won't spin across networks so to say. But if our implementation uses processes, we can containerize them, and horizontally scale them, by using the same image multiple times. If we were to use threads, we would have to vertical scale, by adding more CPU and memory, which is so to say not scalable in extreme terms, due to limits of techonology.


c) How do you handle message re-ordering?
*Oh ohhh Spaghettios*


d) How do you handle message loss?
*Oh ohhh Spaghettios*


e) Why is the 3-way handshake important?
Opposed to UDP, TCP is connection-oriented, meaning we want to be sure that we are actually having a conversation with someone, who is answering. So by having the handshakes, we negotiate the starting sequence numbers, and make sure the connection is established before sending data.
