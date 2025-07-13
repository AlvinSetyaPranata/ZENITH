import { Router } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Suspense } from "solid-js";
import "./app.css";
import RootLayout from "./components/layouts/RootLayout";

export default function App() {
  return (
    <Router
      root={(props) => (
        <RootLayout>
          <Suspense>{props.children}</Suspense>
        </RootLayout>
      )}
    >
      <FileRoutes />
    </Router>
  );
}
