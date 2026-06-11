import type { ReactNode } from "react";
import "./Layout.css";

type LayoutProps = {
  children: ReactNode;
  actions?: ReactNode;
  onHelp?: () => void;
  onCode?: () => void;
};

function Layout({ children, actions, onHelp, onCode }: LayoutProps) {
  return (
    <div className="app-shell">
      <header className="top-bar">
        <div className="top-bar__inner">
          <span className="top-bar__title">Connect Four</span>

          <nav className="top-bar__actions">
            {actions ?? (
              <>
                <button className="top-bar__icon-btn" type="button" aria-label="Help" onClick={onHelp}>
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                    <circle cx="12" cy="12" r="10" />
                    <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3" />
                    <path d="M12 17h.01" />
                  </svg>
                </button>
                <button className="top-bar__icon-btn" type="button" aria-label="Code" onClick={onCode}>
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                    <polyline points="16 18 22 12 16 6" />
                    <polyline points="8 6 2 12 8 18" />
                  </svg>
                </button>
              </>
            )}
          </nav>
        </div>
      </header>

      <main className="app-shell__content">{children}</main>
    </div>
  );
}

export default Layout;
