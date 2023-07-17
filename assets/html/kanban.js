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

