function getFullTasks() {
    $.ajax({
        type: "POST",
        url: "/todo/fulltask/all",
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

function getTags() {
    $.ajax({
        type: "POST",
        url: "/todo/tag/all",
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

function addFulltask() {
    var dataStr = 
        '{"TaskContent":"' + 
        $("button.full-task.add").siblings("input.task.content").val() + 
        '", "TagsID":[' + 
        $("button.full-task.add").siblings("input.tag.id").val() + ']}';
    alert("post:" + dataStr);

    $.ajax({
        type: "POST",
        url: "/todo/fulltask/add",
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
    var dataStr = 
        '{"TaskID":' + 
        $("button.full-task.del").siblings("input.task.id").val() + '}';
    alert("post:" + dataStr);

    $.ajax({
        type: "POST",
        url: "/todo/fulltask/del",
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
    var dataStr = 
        '{"TaskID":' +
        $("button.full-task.upd").siblings("input.task.id").val() +
        ', "TaskContent":"' + 
        $("button.full-task.upd").siblings("input.task.content").val() + 
        '", "TagsID":[' + 
        $("button.full-task.upd").siblings("input.tag.id").val() + ']}';
    alert("post:" + dataStr);

    $.ajax({
        type: "POST",
        url: "/todo/fulltask/upd",
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

function addTag() {
    var dataStr = 
        '{"TagContent":"' + 
        $("button.tag.add").siblings("input.tag.content").val() + 
        '", "TagDesc":"' + 
        $("button.tag.add").siblings("input.tag.desc").val() + '"}';
    alert("post:" + dataStr);

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
    var dataStr = 
        '{"TagID":' + 
        $("button.tag.del").siblings("input.tag.id").val() + '}';
    alert("post:" + dataStr);

    $.ajax({
        type: "POST",
        url: "/todo/tag/del",
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
    var dataStr = 
        '{"TagID":' +
        $("button.tag.upd").siblings("input.tag.id").val() +
        ', "TagContent":"' + 
        $("button.tag.upd").siblings("input.tag.content").val() + 
        '", "TagDesc":"' + 
        $("button.tag.upd").siblings("input.tag.desc").val() + '"}';
    alert("post:" + dataStr);

    $.ajax({
        type: "POST",
        url: "/todo/tag/upd",
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
