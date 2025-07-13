import { Router } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Suspense } from "solid-js";
import "./app.css";
import Sidebar from "./components/ui/sidebar";

export default function App() {
  return (
    <Router
      root={props => (
        <div class="flex">
          <Sidebar title="ZENITH" logoHref="https://i0.wp.com/www.uim.ac.id/uimv2/wp-content/uploads/2020/10/Ico.png?resize=200%2C200&ssl=1" />
          <Suspense>{props.children}</Suspense>
        </div>
      )}
      >
      <FileRoutes />
    </Router>
  );
}
