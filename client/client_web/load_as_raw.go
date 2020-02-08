package client_web

// 下边的内容就是当前目录下 index.html 的内容, 只是为了方便使用 raw 调用, 写入到了代码中而已
func LoadTemplate() string {
	return `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">

    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

    <title>golang crontab dashboard</title>
</head>
<body>

<div class="container-fluid">
    <center><h2>golang计划任务</h2></center>
    <div class="row">
        <div class="col-sm-6">

            <center><h2>任务列表</h2></center>
            <div class="input-group ">
                3s周期的测试任务:
                <input id="second" type="number" value="3">
                <div class="input-group-append">
                    <button class="btn btn-sm btn-success" type="button" onclick="create()">创建</button>
                </div>
            </div>
            <table class="table table-dark table-striped">
                <thead>
                <tr>
                    <th>id</th>
                    <th>任务名字</th>
                    <th>运行状态</th>
                    <th>操作</th>
                </tr>
                </thead>
                <tbody id="taskList">
                <!--<tr>
                    <td>start_per_day</td>
                    <td>
                        <span class="badge badge-dark">已停止</span>
                        <span class="badge badge-success">运行中</span>
                    </td>
                    <td>
                        <button type="button" class="btn btn-sm btn-success" onclick="oper('start','id')">启动</button>
                        <button type="button" class="btn btn-sm btn-warning" onclick="oper('stop','id')">停止</button>
                        <button type="button" class="btn btn-sm btn-danger" onclick="oper('remove','id')">删除</button>
                    </td>
                </tr>-->
                </tbody>
            </table>

            <h2>api说明</h2>

            <p>启动所有任务 api: /start</p>
            <p>启动一个任务 api: /start/{id}</p>
            <p>停止一个任务 api: /stop/{id}</p>
            <p>移除一个任务 api: /remove/{id}</p>
            <p>任务列表 api: /taskList</p>
            <p>任务日志 api: /log?limit=20</p>
        </div>


        <div class="col-sm-6">

            <center><h2>运行日志</h2></center>
            <div class="input-group ">
                刷新日志(limit=20):
                <input id="refresh" type="number" value="20">
                <div class="input-group-append">
                    <button class="btn btn-sm btn-success" type="button" onclick="refresh()">刷新</button>
                </div>
            </div>
            <table class="table table-sm table-striped">
                <tbody id="logList">
                <!--<tr>
                    <td>
                        [2020-02-05 12:52:24] [info] 开始任务:statistic_of_per_day
                    </td>
                </tr>
                <tr>
                    <td>
                        [2020-02-05 12:52:24] [info] 开始任务:statistic_of_per_day
                    </td>
                </tr>-->
                </tbody>
            </table>
        </div>
    </div>
</div>

<script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.15.0/umd/popper.min.js"
        integrity="sha384-L2pyEeut/H3mtgCBaUNw7KWzp5n9+4pDQiExs933/5QfaTh8YStYFFkOzSoXjlTb"
        crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"
        integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
        crossorigin="anonymous"></script>

<script>
    taskList()

    var freshDom = $("#refresh")
    var logListDom = $("#logList")
    var taskListDom = $("#taskList")

    function refresh() {
        logList(freshDom.val())
    }

    function logList(limit) {
        $.get("/log", {limit: limit}, function (res) {
            var html = ""
            for (var i = res.length; i > 0; i--) {
                html += ('<tr><td>' + res[i - 1] + '</td></tr>')
            }
            logListDom.html(html)
        })
    }

    function taskList() {
        $.get("/tasklist", function (resp) {
            var html = ""
            for (var i in resp) {
                var res = resp[i]
                var badge = '<span class="badge badge-warning">已停止</span>\n'
                var btns = "<button type=\"button\" class=\"btn btn-sm btn-success\" onclick=\"oper('start'," + res.id + ")\">启动</button> " +
                    "<button type=\"button\" class=\"btn btn-sm btn-danger\" onclick=\"oper('remove'," + res.id + ")\">删除</button>"
                if (res.status == "运行中") {
                    badge = '<span class="badge badge-success">运行中</span>\n'
                    btns = "<button type=\"button\" class=\"btn btn-sm btn-warning\" onclick=\"oper('stop'," + res.id + ")\">停止</button>"

                }
                html += "<tr>\n" +
                    "                    <td>" + res.id + "</td>\n" +
                    "                    <td>" + res.title + "</td>\n" +
                    "                    <td>\n" +
                    badge +
                    "                    </td>\n" +
                    "                    <td>\n" +btns
                "                    </td>\n" +
                    "                </tr>"
            }
            taskListDom.html(html)
            refresh()
        })
    }

    function create() {
        $.get("/new/" + $("#second").val(), {}, function () {
            taskList()
        })
    }

    function oper(api, id) {
        $.get("/" + api + "/" + id, {}, function () {
            taskList()
        })
    }
</script>
</body>
</html>`
}