$(function () {
    let tasks = $(".task");
    tasks.on("dragstart", function (event) {
        event.originalEvent.dataTransfer.setData("text", event.target.id);
    });
    tasks.on("dblclick", function () {
        window.location = `/pages/${$(this).attr("id")}`
    })

    let blocks = $(".kanban-block");
    blocks.on("drop", function (event) {
        event.preventDefault();
        let data = event.originalEvent.dataTransfer.getData("text");
        event.originalEvent.currentTarget.appendChild(document.getElementById(data));
    });
    blocks.on("dragover", function (event) {
        event.preventDefault();
    });
})

