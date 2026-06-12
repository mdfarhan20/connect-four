/// <reference lib="webworker" />

import * as Comlink from "comlink";
import type { WasmApi } from "../types/wasm";
import "../scripts/wasm_exec.js";


const go = new Go();

const wasmReady = (async () => {
  const result = await WebAssembly.instantiateStreaming(
    fetch(`${import.meta.env.BASE_URL}main.wasm`),
    go.importObject
  );

  go.run(result.instance);
})();

const wasmApi: WasmApi = {
  async ready() {
    await wasmReady;
    return true;
  },

  async startGame() {
    await wasmReady;
    return globalThis.Game.startGame()
  },

  async makePlayerMove(col: number) {
    await wasmReady;
    return globalThis.Game.makePlayerMove(col);
  },

  async makeBotMove() {
    await wasmReady;
    return globalThis.Game.makeBotMove();
  },

  async resetGame() {
    await wasmReady;
    return globalThis.Game.resetGame();
  },
};

Comlink.expose(wasmApi);