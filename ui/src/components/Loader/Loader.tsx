import "./Loader.css";

function Loader() {
  return (
    <main className="loader" aria-busy="true" aria-live="polite">
      <div className="loader__chip" aria-hidden="true" />
      <p className="body-md">Loading game engine</p>
    </main>
  );
}

export default Loader;
