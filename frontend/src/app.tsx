import { Router } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Suspense } from "solid-js";
import "./app.css";

import { AuthProvider } from "./contexts/auth-context";
import { SoonerProvider } from "./contexts/sooner-context";
import RootLayout from "./components/layouts/RootLayout";

export default function App() {
  return (
    <Router
      root={(props) => (
        <Suspense>
          <AuthProvider>
            <SoonerProvider>
              <RootLayout>{props.children}</RootLayout>
            </SoonerProvider>
          </AuthProvider>
        </Suspense>
      )}
    >
      <FileRoutes />
    </Router>
  );
}
