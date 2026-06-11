import type { ReactNode } from "react";
import "./Layout.css";

type LayoutProps = {
  children: ReactNode;
  status?: ReactNode;
  actions?: ReactNode;
};

function Layout({ children, status, actions }: LayoutProps) {
  return (
    <div className="app-shell">
      <header className="top-bar">
        <a className="top-bar__brand" href="/" aria-label="Connect Four home">
          <span className="top-bar__mark" aria-hidden="true">
            <span />
            <span />
            <span />
            <span />
          </span>
          <span>Connect Four</span>
        </a>

        {status ? <div className="top-bar__status">{status}</div> : null}

        <div className="top-bar__meta">
          {actions ?? <span className="status-badge status-badge-info">AI opponent</span>}
        </div>
      </header>

      <main className="app-shell__content">{children}</main>
    </div>
  );
}

export default Layout;
