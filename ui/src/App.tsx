import { useCallback, useEffect, useState } from "react"
import GameScreen from "./components/GameScreen/GameScreen";
import HowToPlay from "./components/HowToPlay/HowToPlay";
import Layout from "./components/Layout/Layout";
import Loader from "./components/Loader/Loader";
import { initializeWasm } from "./services/wasm_client";

function App() {
  const [wasmLoaded, setWasmLoaded] = useState(false);
  const [screen, setScreen] = useState<"game" | "how-to-play">("game");

  useEffect(() => {
    initializeWasm().then(() => {
      setWasmLoaded(true);
    });
  }, []);

  const handleHelp = useCallback(() => setScreen("how-to-play"), []);
  const handleCode = useCallback(() => {
    window.open("https://github.com/mdfarhan20/connect-four", "_blank", "noopener");
  }, []);

  if (!wasmLoaded) return <Loader />;

  if (screen === "how-to-play") {
    return (
      <Layout onCode={handleCode}>
        <HowToPlay onBack={() => setScreen("game")} />
      </Layout>
    );
  }

  return (
    <Layout onHelp={handleHelp} onCode={handleCode}>
      <GameScreen />
    </Layout>
  )
}

export default App
