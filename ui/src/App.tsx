import { useEffect, useState } from "react"
import HomeScreen from "./components/HomeScreen/HomeScreen";
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
    <Layout>
      <HomeScreen />
    </Layout>
  )
}

export default App
