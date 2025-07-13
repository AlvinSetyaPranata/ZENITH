// app.config.ts
import { defineConfig } from "@solidjs/start/config";
import tailwindcss from "@tailwindcss/vite"; // Import the Tailwind Vite plugin

export default defineConfig({
  // Other Vinxi configurations might be here (e.g., server, dev)

  // This is where you configure the underlying Vite instance
  vite: {
    plugins: [
      tailwindcss(), // Add the Tailwind Vite plugin here
    ],
  },
});