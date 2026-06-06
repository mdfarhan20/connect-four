import { useEffect } from "react"

function App() {
  useEffect(() => {
    async function loadWasm() {
      const go = new Go();

      const result = await WebAssembly.instantiateStreaming(
        fetch("/main.wasm"),
        go.importObject,
      );

      go.run(result.instance);

      console.log("WASM", window.goAdd(1, 20))
    }

    loadWasm();
  })

  return (
    <div>
      <h1>Connect Four</h1>
    </div>
  )
}

export default App
