// src/app.tsx
import { Suspense } from "solid-js";
import { Router } from "@solidjs/router"; // Import Router
import { FileRoutes } from "@solidjs/start/router"; // Correct import for FileRoutes
import './app.css'

// Import your layout components
// Make sure the paths are correct relative to app.tsx
import Sidebar from "./components/ui/sidebar";
import SidebarNav from "./components/molecules/sidebar-nav";

export default function App() {
  return (
    <Router
      root={(props) => ( // This 'root' prop is where your GLOBAL layout goes
        <div class="flex">
          <Sidebar
            title="Universitas Islam Madura"
            logoHref="https://i0.wp.com/www.uim.ac.id/uimv2/wp-content/uploads/2020/10/Ico.png?resize=200%2C200&ssl=1"
          >
            <SidebarNav title="Dashboard" />
          </Sidebar>
          {/* This is where the actual route content will be rendered */}
          <Suspense>{props.children}</Suspense>
        </div>
      )}
    >
      {/* FileRoutes MUST be a direct child of Router */}
      <FileRoutes />
    </Router>
  );
}