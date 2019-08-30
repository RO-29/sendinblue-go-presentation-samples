# Dependency Injection in GO

## GO Dependency injection example

    ➜ go run main.go

    ** Init DI Provider A **<
    ** Init serviceA **
    ** Init DI Provider B **
    ** Init serviceB **
    from A:
    hey there:  testing
    ** Init DI Provider C **
    ** Init serviceC **
    ** Using already init DI Provider B **
    from C:
    hey there:  testing

## Motivation

Go uses struct and interfaces to declare a service.
A service can inherit another service and as many it wants ( **but we must try to keep it minimal for loosely coupled dependency graph**)

This tutorial assumes all the go services must be initialized at the program initiation not at the run time of program.

Now the problem can arise if a service is required by multiple services!
It will essentially means we will be initializing a service as many times as it is required by other services.
It's not an ideal scenario in a case service that is being initialized multiple times has a network related dependency like database initialization or some other complex initialization.

Lets view this problem from the famous diamond perspective but there is no actual inheritance here just the service dependency angle. (**Though go supports inheritance but this is out of scope of this tutorial**)

          A
        B   C
          D
**Incomplete
