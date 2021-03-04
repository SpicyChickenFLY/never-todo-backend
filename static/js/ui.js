function getFullTasks() {
    var urlStr = "/todo/fulltask/all";
    $.ajax({
        type: "POST",
        url: urlStr,
        dataType: "json",
        success: function (data, textStatus) {
            $("#full-tasks pre").html(
                syntaxHighlight(data)
            )
            alert("Post\nurl:" + urlStr + "\nresult: success");
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("error");
        }
    });
}

function addFulltask() {
    var urlStr = "/todo/fulltask/add";
    var dataStr = 
        '{"TaskContent":"' + 
        $("button.full-task.add").siblings("input.task.content").val() + 
        '", "TagsID":[' + 
        $("button.full-task.add").siblings("input.tag.id").val() + ']}';

    $.ajax({
        type: "POST",
        url: urlStr,
        contentType: "application/json",
        data: dataStr,
        dataType: "json",
        success: function(data, textStatus) {
            alert("Post\nurl:" + urlStr + "\ndata:" + dataStr + "\nresult:" + JSON.stringify(data));
            getFullTasks();
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("Request Error [addTask]");
        }
    });
}

function delFulltask() {
    var urlStr = "/todo/fulltask/del";
    var dataStr = 
        '{"TaskID":' + 
        $("button.full-task.del").siblings("input.task.id").val() + '}';
    alert("Post\nurl:" + urlStr + "data:\n" + dataStr);

    $.ajax({
        type: "POST",
        url: urlStr,
        contentType: "application/json",
        data: dataStr,
        dataType: "json",
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
    var urlStr = "/todo/fulltask/upd"
    var dataStr = 
        '{"TaskID":' +
        $("button.full-task.upd").siblings("input.task.id").val() +
        ', "TaskContent":"' + 
        $("button.full-task.upd").siblings("input.task.content").val() + 
        '", "TagsID":[' + 
        $("button.full-task.upd").siblings("input.tag.id").val() + ']}';
    alert("Post\nurl:" + urlStr + "data:\n" + dataStr);

    $.ajax({
        type: "POST",
        url: urlStr,
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
    var urlStr = "/todo/tag/all";
    alert("Post\nurl:" + urlStr);
    $.ajax({
        type: "POST",
        url: urlStr,
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
    var dataStr = 
        '{"TagContent":"' + 
        $("button.tag.add").siblings("input.tag.content").val() + 
        '", "TagDesc":"' + 
        $("button.tag.add").siblings("input.tag.desc").val() + '"}';
    alert("Post\nurl:" + urlStr + "data:\n" + dataStr);

    $.ajax({
        type: "POST",
        url: "/todo/tag/add",
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
    var urlStr = "/todo/tag/del";
    var dataStr = 
        '{"TagID":' + 
        $("button.tag.del").siblings("input.tag.id").val() + '}';
    alert("Post\nurl:" + urlStr + "data:\n" + dataStr);

    $.ajax({
        type: "POST",
        url: urlStr,
        contentType: "application/json",
        data: dataStr,
        dataType: "json",
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
    var urlStr = "/todo/tag/upd";
    var dataStr = 
        '{"TagID":' +
        $("button.tag.upd").siblings("input.tag.id").val() +
        ', "TagContent":"' + 
        $("button.tag.upd").siblings("input.tag.content").val() + 
        '", "TagDesc":"' + 
        $("button.tag.upd").siblings("input.tag.desc").val() + '"}';
    alert("Post\nurl:" + urlStr + "data:\n" + dataStr);

    $.ajax({
        type: "POST",
        url: urlStr,
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
    $("button.full-task.all").click(getFullTasks);

    $("button.tag.add").click(addTag);
    $("button.tag.del").click(delTag);
    $("button.tag.upd").click(updTag);
    $("button.tag.all").click(getTags);
});
