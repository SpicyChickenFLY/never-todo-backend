function getTasks() {
    $.post(
        "localhost:8080/todo/tasklist/all",
        function(data) {
            alert(data)
            status = data.status
        }
    );
}

function getTags() {}

$(document).ready(function () {
    getTasks();
    getTags();

    $("button#popup").click(function () {
        $(".popup").show();
    });

    $("button#popdown").click(function() {
        $(".popup").hide();
    });



});

