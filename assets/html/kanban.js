function addTaskEvent(obj) {
    obj.on("dragstart", function (event) {
        event.originalEvent.dataTransfer.setData("text", event.target.id);
    });
    obj.on("dblclick", function () {
        window.location = `/pages/${$(this).attr("id")}`
    })
}

$(function () {
    function getDivByStatus(status) {
        switch (status) {
            case "Draft":
                return $("#draft");
            case "Proposed":
                return $("#proposed");
            case "Rejected":
                return $("#rejected");
            case "Accepted":
                return $("#accepted");
            case "Deprecated":
                return $("#deprecated");
            case "Superseded":
                return $("#superseded");
        }
    }

    $.ajax({
        contentType: 'application/json',
        url: `/summaries`,
        type: 'get'
    }).done(function (data) {
        data.forEach(function (summary) {
            let div = getDivByStatus(summary["status"]);
            let task = $(`<div class="task" id="${summary["title"]}" draggable="true"><span>${summary["title"]}</span></div>`);
            div.append(task);
            addTaskEvent(task);
        })
    });
})


$(function () {
    let tasks = $(".task");
    addTaskEvent(tasks);

    let blocks = $(".kanban-block");
    blocks.on("drop", function (event) {
        event.preventDefault();
        let data = event.originalEvent.dataTransfer.getData("text");
        event.originalEvent.currentTarget.appendChild(document.getElementById(data));
        $.ajax({
            contentType: 'application/json',
            url: `/pages/${data}`,
            type: 'put',
            data: JSON.stringify({
                status: event.originalEvent.currentTarget.id
            })
        });
    });
    blocks.on("dragover", function (event) {
        event.preventDefault();
    });
});

$(function () {
    const client = mqtt.connect("ws://127.0.0.1:9001");
    client.subscribe("adr");
    client.on("message", function (topic, payload) {
        console.log(payload);
        console.log(JSON.parse(payload));
    });
});

