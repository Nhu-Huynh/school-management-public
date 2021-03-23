package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["scholy/controllers:ClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:ClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:ClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:ClassController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:ClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:ClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:CourseController"] = append(beego.GlobalControllerRouter["scholy/controllers:CourseController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:CourseController"] = append(beego.GlobalControllerRouter["scholy/controllers:CourseController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:CourseController"] = append(beego.GlobalControllerRouter["scholy/controllers:CourseController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:CourseController"] = append(beego.GlobalControllerRouter["scholy/controllers:CourseController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:CourseController"] = append(beego.GlobalControllerRouter["scholy/controllers:CourseController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentClassController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentClassController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentClassController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentClassController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentClassController"],
        beego.ControllerComments{
            Method: "GetByClass",
            Router: "/by_class",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentClassController"],
        beego.ControllerComments{
            Method: "GetByStudent",
            Router: "/by_student",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentClassController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentClassController"],
        beego.ControllerComments{
            Method: "GetStudentClass",
            Router: "/by_studentclass",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["scholy/controllers:StudentController"] = append(beego.GlobalControllerRouter["scholy/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
