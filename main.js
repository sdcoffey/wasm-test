function setTime(time) {
  document.getElementById('time').innerHTML = time;
}

if (WebAssembly) {
	const go = new Go();
  fetch("go.wasm").then(response => response.arrayBuffer())
    .then(buffer => WebAssembly.instantiate(buffer, go.importObject))
    .then((result) => {
      go.run(result.instance);

      setTime(currentTime());

      everySecond((time) => {
        setTime(time);
      });
    })
} else {
	console.log("WebAssembly is not supported in your browser")
}
