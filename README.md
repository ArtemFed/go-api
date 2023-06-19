# go-api
 
### handlers:
Students

    GET    /api/students/get-all     --> go-api/src/transport/handler.(*Handler).getAllStudents<br/>
    GET    /api/students/get-by-id   --> go-api/src/transport/handler.(*Handler).getStudentById<br/>
    POST   /api/students/create      --> go-api/src/transport/handler.(*Handler).createStudent<br/>
    PUT    /api/students/update      --> go-api/src/transport/handler.(*Handler).updateStudent<br/>
    DELETE /api/students/delete      --> go-api/src/transport/handler.(*Handler).deleteStudent<br/>
    
Grades

    GET    /api/grades/get-by-id     --> go-api/src/transport/handler.(*Handler).getGradeById<br/>
    GET    /api/grades/get-by-student-id --> go-api/src/transport/handler.(*Handler).getGradesByStudentId<br/>
    POST   /api/grades/create        --> go-api/src/transport/handler.(*Handler).createGrade<br/>
    DELETE /api/grades/delete        --> go-api/src/transport/handler.(*Handler).deleteGrade<br/>
