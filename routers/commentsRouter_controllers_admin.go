package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/link`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/link/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/link/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/link/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/link/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/link/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/menu`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/menu/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/menu/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/menu/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/menu/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["kangg/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/menu/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
