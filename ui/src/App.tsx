import { useEffect, useState } from "react"
import GameScreen from "./components/GameScreen/GameScreen";
import Layout from "./components/Layout/Layout";
import Loader from "./components/Loader/Loader";
import { initializeWasm } from "./services/wasm_client";

function App() {
  const [wasmLoaded, setWasmLoaded] = useState(false);

  useEffect(() => {
    initializeWasm().then(() => {
      setWasmLoaded(true);
    });
  }, []);

  if (!wasmLoaded) return <Loader />;

  return (
    <Layout
      status={<span className="status-badge status-badge-turn-red">Your turn</span>}
      actions={
        <>
          <button className="top-bar__action" type="button">New game</button>
          <button className="top-bar__action" type="button">Reset</button>
        </>
      }
    >
      <GameScreen />
    </Layout>
  )
}

export default App
