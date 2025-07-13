// vinxi.config.ts
import { defineConfig } from "@solidjs/start/config";
import solid from "solid-start/vite";

export default defineConfig({
  vite: {
    plugins: [solid()],
    server: {
      hmr: {
        protocol: "ws", // or "wss" for HTTPS
        host: "localhost",
        port: 5173 // default Vite dev server port, adjust if needed
      }
    }
  }
});
