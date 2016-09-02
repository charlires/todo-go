ToDo
========

This application has been written for a talk on Go, TDD, and REST APIs.

TODO provides a simple task management application.

The purpose of this project is to show how some general techniques can be applied to Go code.

##### Test Driven Development

The task package has been implemented using TDD techniques, starting writing tests, seeing them failing, writing the code to make the test pass.

Once all tests pass the code is refactored.

##### Implementation of a REST API in Go

The server package handles requests on `"/task/"` providing a REST API.

This REST API is then exposed as a stand-alone http server.
Please take into account that the TaskManager is only in-memory at this point, which means that tasks are
lost if the application is restarted.

