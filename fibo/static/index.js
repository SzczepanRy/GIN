const socket = new WebSocket("ws://localhost:3000/connect");

const ul = document.querySelector(".numbers");
const p = document.querySelector(".connection");

window.addEventListener("load", () => {
    socket.open = () => {
        p.innerHTML = "CONNECTED";
    };
});

socket.onmessage = ({ data }) => {
    const li = document.createElement("li");
    li.innerHTML = data;
    ul.append(li);
};
