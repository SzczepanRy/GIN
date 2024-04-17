const button = document.querySelector(".button");
const images = document.querySelector(".images");

button.addEventListener("click", async (e) => {
    e.preventDefault();
    const id = document.querySelector(".id");

    let res = await fetch("http://localhost:3000/get", {
        method: "post",
        headers: {
            "Content-type": "application/json",
        },
        body: JSON.stringify({ id: +id.value }),
    });

    let data = await res.json();
    console.log(data);

    if (typeof data.file == "string") {
        let image = new Image();
        image.src = "data:image/png;base64," + data.file;
        images.append(image);
    } else if (Array.isArray(data.file)) {
        let { file } = data;
        // console.log(typeof file);
        for (let fi of file) {
            let image = new Image();
            image.src = "data:image/png;base64," + fi;
            images.append(image);
        }
    }
});
