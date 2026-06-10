import * as Comlink from "comlink";
import type { WasmApi } from "../types/wasm";

const worker = new Worker(
  new URL("../workers/go.worker.ts", import.meta.url),
  {
    type: 'module'
  }
);

export const wasm = Comlink.wrap<WasmApi>(worker);

let initialized = false;

export async function initializeWasm() {
  if (initialized) return;

  await wasm.ready();
  initialized = true;
}

export function isWasmReady() {
  return initialized;
}