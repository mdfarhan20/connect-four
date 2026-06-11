import { useEffect, useRef, useState } from "react"
import Loader from "./components/Loader/Loader";
import { initializeWasm, wasm } from "./services/wasm_client";

function App() {
  const columnRef = useRef<HTMLInputElement | null>(null);

  const [wasmLoaded, setWasmLoaded] = useState(false);

  useEffect(() => {
    initializeWasm().then(() => {
      setWasmLoaded(true);
    });
  }, []);

  const handleStartGame =  async () => {
    console.log("Game Started", await wasm.startGame());
  }

  const handleMakeMove = async () => {
    if (!columnRef.current) return;
    console.log(`Move made: ${columnRef.current.value}`);
    console.log(await wasm.makePlayerMove(Number(columnRef.current.value)))
  }

  const handleBotMove = async () => {
    console.log("Bot move", await wasm.makeBotMove());
  }

  const handleReset = async () => {
    console.log("Game reset", await wasm.resetGame());
  }

  if (!wasmLoaded) return <Loader />;

  return (
    <div>
      <h1>Connect Four</h1>

      <button onClick={handleStartGame}>
        Start Game
      </button>

      <input type="number" min={0} max={6} ref={columnRef} />
      <button onClick={handleMakeMove}>Make move</button>

      <button onClick={handleBotMove}>Make Bot move</button>

      <button onClick={handleReset}>Reset game</button>
    </div>
  )
}

export default App
