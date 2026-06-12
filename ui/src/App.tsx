import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import toast, { Toaster } from "react-hot-toast";
import GameScreen from "./components/GameScreen/GameScreen";
import HowToPlay from "./components/HowToPlay/HowToPlay";
import Layout from "./components/Layout/Layout";
import Loader from "./components/Loader/Loader";
import { initializeWasm, wasm } from "./services/wasm_client";
import type { BoardState, Player } from "./types/game";
import HomeScreen from "./components/HomeScreen/HomeScreen";
import { CELLS } from "./constants/game";
import coinDropSfx from "./assets/sound/coin-drop.mp3";

function App() {
  const [wasmLoaded, setWasmLoaded] = useState(false);
  const [screen, setScreen] = useState<"game" | "how-to-play">("game");
  const [gameStarted, setGameStarted] = useState<boolean>(false);
  const [selectedPlayer, setSelectedPlayer] = useState<Player>(1);
  const [gameState, setGameState] = useState<BoardState | null>(null);

  const [botThinking, setBotThinking] = useState<boolean>(false);
  const [moveLoading, setMoveLoading] = useState<boolean>(false);

  const coinAudio = useRef<HTMLAudioElement | null>(null);
  const playCoinDrop = useCallback(() => {
    if (!coinAudio.current) {
      coinAudio.current = new Audio(coinDropSfx);
    }
    coinAudio.current.currentTime = 0;
    coinAudio.current.play().catch(() => {});
  }, []);

  const isGameOver = useMemo(() => {
    if (!gameState) return;
    return gameState.winner !== CELLS.EMPTY || gameState.isDraw;
  }, [gameState]);

  const handleGameStart = async () => {
    const res = await wasm.startGame();
    if (res.status === "error") {
      showError(res.message);
      return;
    }

    setGameState(res.data);
    setGameStarted(true);
  };

  const handlePlayerMove = async (col: number) => {
    setMoveLoading(true);
    const res = await wasm.makePlayerMove(col);
    if (res.status === "error") {
      showError(res.message);
      setMoveLoading(false);
      return;
    }

    setGameState(res.data);
    setMoveLoading(false);
    playCoinDrop();
  };

  const handleBotMove = async () => {
    setBotThinking(true);
    const res = await wasm.makeBotMove();
    if (res.status === "error") {
      showError(res.message);
      setBotThinking(false);
      return;
    }

    setGameState(res.data);
    setBotThinking(false);
    playCoinDrop();
  };

  const handleResetGame = async () => {
    const res = await wasm.resetGame();
    if (res.status === "error") {
      showError(res.message);
      return;
    }

    setGameState(res.data);
  };

  const showError = useCallback((message: string) => {
    toast.error(message, { duration: 4000 });
  }, []);

  const handleHelp = useCallback(() => setScreen("how-to-play"), []);
  const handleCode = useCallback(() => {
    window.open(
      "https://github.com/mdfarhan20/connect-four",
      "_blank",
      "noopener",
    );
  }, []);

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

  if (!wasmLoaded) return <Loader />;

  const toaster = (
    <Toaster
      position="bottom-right"
      toastOptions={{
        style: {
          fontFamily: "'Plus Jakarta Sans', system-ui, sans-serif",
          fontSize: "14px",
          fontWeight: 600,
          borderRadius: "var(--rounded-md, 0.75rem)",
          background: "var(--color-error-container, #ffdad6)",
          color: "var(--color-on-error-container, #93000a)",
          padding: "12px 16px",
          boxShadow: "0 4px 12px rgba(186, 26, 26, 0.15)",
        },
        icon: null,
      }}
    />
  );

  if (screen === "how-to-play") {
    return (
      <>
        {toaster}
        <Layout onCode={handleCode}>
          <HowToPlay onBack={() => setScreen("game")} />
        </Layout>
      </>
    );
  }

  return (
    <>
      {toaster}
      <Layout onHelp={handleHelp} onCode={handleCode}>
        {gameStarted && gameState ? (
          <GameScreen
            gameState={gameState}
            moveLoading={moveLoading}
            botThinking={botThinking}
            onPlayerMove={handlePlayerMove}
            onResetGame={handleResetGame}
            playerColor={selectedPlayer}
          />
        ) : (
          <HomeScreen
            player={selectedPlayer}
            onPlayerChange={setSelectedPlayer}
            onGameStart={handleGameStart}
          />
        )}
      </Layout>
    </>
  );
}

export default App;
