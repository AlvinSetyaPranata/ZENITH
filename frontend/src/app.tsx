import { Router } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Suspense } from "solid-js";
import "./app.css";
import RootLayout from "./components/layouts/RootLayout";
import { AuthProvider } from "./contexts/auth-context";

export default function App() {
  return (
    <Router
      root={(props) => (
        <Suspense>
          <AuthProvider>{props.children}</AuthProvider>
        </Suspense>
      )}
    >
      <FileRoutes />
    </Router>
  );
}
