waitgroup
WaitGroup is a great way to wait for a set of concurrent operations to complete. 
The waitgroup can be used if 1) we are not bothered about the result of the concurrent operations or 2) we can collect the results in some othe way. if both the conditions are not satisfied then channels has to be used instead of waitgroups.

mutex
mutex is way to gaurd the critical section of the code.A critical section is the code which requires exclusive access to 
the shared resources.

channel

direction channel 

select
