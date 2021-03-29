function getFullTasks() {
    $.ajax({
        type: "GET",
        url: "/api/v1/todo/fulltask/",
        dataType: "json",
        success: function (data, textStatus) {
            $("#full-tasks pre").html(
                syntaxHighlight(data)
            ) 
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("error");
        }
    });
}

function getFullTasksByContent() {
    let urlStr = 
        "/api/v1/todo/fulltask/content/" + 
        $("button.full-task.rtr").siblings("input.task.content").val()
    $.ajax({
        type: "GET",
        url: urlStr,
        dataType: "json",
        success: function (data, textStatus) {
            $("#full-tasks pre").html(
                syntaxHighlight(data)
            ) 
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("error");
        }
    });
}

function getFullTasksByTagID() {
    $.ajax({
        type: "GET",
        url: "/api/v1/todo/fulltask/",
        dataType: "json",
        success: function (data, textStatus) {
            $("#full-tasks pre").html(
                syntaxHighlight(data)
            ) 
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("error");
        }
    });
}

function addFulltask() {
    let dataStr = 
        '{"TaskContent":"' + 
        $("button.full-task.add").siblings("input.task.content").val() + 
        '", "TagsID":[' + 
        $("button.full-task.add").siblings("input.tag.id").val() + ']}';
    alert("post:" + dataStr);

    $.ajax({
        type: "POST",
        url: "/api/v1/todo/fulltask/",
        contentType: "application/json",
        data: dataStr,
        dataType: "json",
        success: function(data, textStatus) {
            alert(JSON.stringify(data));
            getFullTasks();
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("Request Error [addTask]");
        }
    });
}

function delFulltask() {
    let urlStr = 
        "/api/v1/todo/fulltask/" + $("button.full-task.del").siblings("input.task.id").val();
    $.ajax({
        type: "DELETE",
        url: urlStr,
        contentType: "application/json",
        success: function(data, textStatus) {
            alert(JSON.stringify(data));
            getFullTasks();
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("Request Error [delTask]");
        }
    });
}

function updFulltask() {
    let dataStr = 
        '{"TaskID":' +
        $("button.full-task.upd").siblings("input.task.id").val() +
        ', "TaskContent":"' + 
        $("button.full-task.upd").siblings("input.task.content").val() + 
        '", "TagsID":[' + 
        $("button.full-task.upd").siblings("input.tag.id").val() + ']}';
    alert("post:" + dataStr);

    $.ajax({
        type: "PUT",
        url: "/api/v1/todo/fulltask/",
        contentType: "application/json",
        data: dataStr,
        dataType: "json",
        success: function(data, textStatus) {
            alert(JSON.stringify(data));
            getFullTasks();
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("Request Error [UpdateTask]");
        }
    });
}

function getTags() {
    $.ajax({
        type: "GET",
        url: "/api/v1/todo/tag",
        dataType: "json",
        success: function (data, textStatus) {
            $("#tags pre").html(
                syntaxHighlight(data)
            ) 
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("Request Error [getTag]");
        }
    });
}

function addTag() {
    let dataStr = 
        '{"TagContent":"' + 
        $("button.tag.add").siblings("input.tag.content").val() + 
        '", "TagDesc":"' + 
        $("button.tag.add").siblings("input.tag.desc").val() + 
        '","TagColor":"' +
        $("button.tag.add").siblings("input.tag.color").val() + '"}';
    alert("post:" + dataStr);

    $.ajax({
        type: "POST",
        url: "/api/v1/todo/tag",
        contentType: "application/json",
        data: dataStr,
        dataType: "json",
        success: function(data, textStatus) {
            alert(JSON.stringify(data));
            getTags();
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("Request Error [addTag]");
        }
    });
}

function delTag() {
    let urlStr = 
        "/api/v1/todo/tag" + $("button.tag.del").siblings("input.tag.id").val();

    $.ajax({
        type: "DELETE",
        url: urlStr,
        contentType: "application/json",
        success: function(data, textStatus) {
            alert(JSON.stringify(data));
            getTags();
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("Request Error [delTag]");
        }
    });
}

function updTag() {
    let dataStr = 
        '{"TagID":' +
        $("button.tag.upd").siblings("input.tag.id").val() +
        ', "TagContent":"' + 
        $("button.tag.upd").siblings("input.tag.content").val() + 
        '", "TagDesc":"' + 
        $("button.tag.upd").siblings("input.tag.desc").val() + 
        '","TagColor":"' +
        $("button.tag.add").siblings("input.tag.color").val() + '"}';
    alert("post:" + dataStr);

    $.ajax({
        type: "PUT",
        url: "/api/v1/todo/tag",
        contentType: "application/json",
        data: dataStr,
        dataType: "json",
        success: function(data, textStatus) {
            alert(JSON.stringify(data));
            getTags();
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("Request Error [UpdateTag]");
        }
    });
}

$(document).ready(function () {
    getFullTasks();
    getTags();

    $("button.full-task.add").click(addFulltask);
    $("button.full-task.del").click(delFulltask);
    $("button.full-task.upd").click(updFulltask);

    $("button.tag.add").click(addTag);
    $("button.tag.del").click(delTag);
    $("button.tag.upd").click(updTag);
});
