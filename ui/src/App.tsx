import { useCallback, useEffect, useMemo, useState } from "react";
import GameScreen from "./components/GameScreen/GameScreen";
import HowToPlay from "./components/HowToPlay/HowToPlay";
import Layout from "./components/Layout/Layout";
import Loader from "./components/Loader/Loader";
import { initializeWasm, wasm } from "./services/wasm_client";
import type { BoardState, Player } from "./types/game";
import HomeScreen from "./components/HomeScreen/HomeScreen";
import { CELLS } from "./constants/game";

function App() {
  const [wasmLoaded, setWasmLoaded] = useState(false);
  const [screen, setScreen] = useState<"game" | "how-to-play">("game");
  const [gameStarted, setGameStarted] = useState<boolean>(false);
  const [selectedPlayer, setSelectedPlayer] = useState<Player>(1);
  const [gameState, setGameState] = useState<BoardState | null>(null);

  const [botThinking, setBotThinking] = useState<boolean>(false);
  const [moveLoading, setMoveLoading] = useState<boolean>(false);

  const isGameOver = useMemo(() => {
    if (!gameState) return
    return gameState.winner !== CELLS.EMPTY || gameState.isDraw;
  }, [gameState]);

  useEffect(() => {
    initializeWasm().then(() => {
      setWasmLoaded(true);
    });
  }, []);

  useEffect(() => {
    if (!gameState || isGameOver) return;

    if (gameState.player !== selectedPlayer) {
      handleBotMove();
    }
  }, [selectedPlayer, gameState]);

  const handleGameStart = async () => {
    const res = await wasm.startGame();
    if (res.status === "error") {
      return;
    }

    setGameState(res.data);
    setGameStarted(true);
  };

  const handlePlayerMove = async (col: number) => {
    setMoveLoading(true);
    const res = await wasm.makePlayerMove(col);
    if (res.status === "error") {
      setMoveLoading(false);
      return;
    }

    setGameState(res.data);
    setMoveLoading(false);
  };

  const handleBotMove = async () => {
    setBotThinking(true);
    const res = await wasm.makeBotMove();
    if (res.status === "error") {
      setBotThinking(false);
      return;
    }

    setGameState(res.data);
    setBotThinking(false);
  };

  const handleHelp = useCallback(() => setScreen("how-to-play"), []);
  const handleCode = useCallback(() => {
    window.open(
      "https://github.com/mdfarhan20/connect-four",
      "_blank",
      "noopener",
    );
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
      {gameStarted && gameState ? (
        <GameScreen
          board={gameState.board}
          moveLoading={moveLoading}
          botThinking={botThinking}
          onPlayerMove={handlePlayerMove}
        />
      ) : (
        <HomeScreen
          player={selectedPlayer}
          onPlayerChange={setSelectedPlayer}
          onGameStart={handleGameStart}
        />
      )}
    </Layout>
  );
}

export default App;
