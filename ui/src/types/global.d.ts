export {};

declare global {
  interface Window {
    goAdd(a: number, b: number): number;
  }
}