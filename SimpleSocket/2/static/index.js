const socket = new WebSocket("ws://localhost:3000/ws");

const ul = document.querySelector("ul");
const input = document.querySelector("input");
const form = document.querySelector("form");

console.log("js works");

window.addEventListener("load", () => {
    socket.onopen = () => {
        const li = document.createElement("li");
        li.innerText = "CONNECTED";
        ul.append(li);
    };
});

form.addEventListener("submit", (e) => {
    e.preventDefault();
    if (input.value) {
        socket.send(input.value);
        input.value = "";
    }
    input.focus();
});
socket.onmessage = (e) => {
    console.log(e.data);
    const li = document.createElement("li");
    li.innerText = e.data;
    ul.append(li);
};
