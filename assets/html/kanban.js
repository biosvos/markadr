let adrs = {};

function addTaskEvent(obj) {
    obj.on("dragstart", function (event) {
        event.originalEvent.dataTransfer.setData("text", event.target.id);
    });
    obj.on("dblclick", function () {
        window.location = `/pages/${$(this).attr("id")}`
    })
}

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

function getStatusByID(id) {
    switch (id) {
        case "draft":
            return "Draft";
        case "proposed":
            return "Proposed";
        case "rejected":
            return "Rejected";
        case "accepted":
            return "Accepted";
        case "deprecated":
            return "Deprecated";
        case "superseded":
            return "Superseded";
    }
}

$(function () {
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

            adrs[summary["title"]] = summary;
        })
    });
})

$(function () {
    let tasks = $(".task");
    addTaskEvent(tasks);

    let blocks = $(".kanban-block");
    blocks.on("drop", function (event) {
        event.preventDefault();
        let title = event.originalEvent.dataTransfer.getData("text");
        event.originalEvent.currentTarget.appendChild(document.getElementById(title));
        let statusID = event.originalEvent.currentTarget.id;
        let status = getStatusByID(statusID)

        $.ajax({
            contentType: 'application/json',
            url: `/pages/${title}`,
            type: 'put',
            data: JSON.stringify({
                status: statusID
            })
        }).done(function () {
            adrs[title] = {
                "title": title,
                "status": status
            };
        })
    });
    blocks.on("dragover", function (event) {
        event.preventDefault();
    });
});

$(function () {
    const client = mqtt.connect("ws://127.0.0.1:9001");
    client.subscribe("adr");
    client.on("message", function (topic, payload) {
        let adr = JSON.parse(payload);
        let title = adr["title"];
        let status = adr["status"];
        let adrElement = $(`#${title}`);
        if (status === "") { // deleted
            adrElement.detach();
            adrs.delete(`${title}`);
            return;
        }

        let div = getDivByStatus(adr["status"]);
        if (!Object.keys(adrs).includes(title)) { // created
            let task = $(`<div class="task" id="${title}" draggable="true"><span>${title}</span></div>`);
            div.append(task);
            addTaskEvent(task);
            adrs[title] = adr;
            return;
        }

        // updated
        if (adrs[title]["status"] === adr["status"]) { // same
            return;
        }
        adrElement.detach();
        div.append(adrElement);
        adrs[title] = adr;
    });
});
