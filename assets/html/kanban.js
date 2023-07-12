const tasks = document.getElementsByClassName("task");
Array.prototype.forEach.call(tasks, (task) => {
    task.addEventListener("dragstart", (event) => {
        event.dataTransfer.setData("text", event.target.id);
    }, false);
});

const kanbanBlocks = document.getElementsByClassName("kanban-block");
Array.prototype.forEach.call(kanbanBlocks, (kanban) => {
    kanban.addEventListener("drop", (event) => {
        event.preventDefault();
        let data = event.dataTransfer.getData("text");
        event.currentTarget.appendChild(document.getElementById(data));
    }, false);
});

Array.prototype.forEach.call(kanbanBlocks, (kanban) => {
    kanban.addEventListener("dragover", (event) => {
        event.preventDefault();
    }, false);
});

