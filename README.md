A Cloud Native implementation of the important Fizz Buzz algorithm

# Enterprise Cloud Native Architecture Reference (WIP)

 - Query CRD: Represents a single Query of the Cloud Native Fizz Buzz Cluster. Implemented as a Kubernetes CRD with the single 'Input' type. The status: either 'Fizz' or 'Buzz' is reflected in the `status` of the CRD after calculation
 - Calculation CRD: Represents an Arbitrary calculation. The Query CRD creates a Calculation to determine whether the `Input` modulo 3 or 5 is 0. 
