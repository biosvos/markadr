$(function () {
    $(".task").on("dragstart", function (event) {
        event.originalEvent.dataTransfer.setData("text", event.target.id);
    });

    let blocks = $(".kanban-block");
    blocks.on("drop", function (event) {
        console.log("hi");
        event.preventDefault();
        let data = event.originalEvent.dataTransfer.getData("text");
        event.originalEvent.currentTarget.appendChild(document.getElementById(data));
    });
    blocks.on("dragover", function (event) {
        event.preventDefault();
    });
})

