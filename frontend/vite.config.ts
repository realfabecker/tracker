import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@pages": path.resolve(__dirname, "./src/pages"),
      "@core": path.resolve(__dirname, "./src/core"),
      "@store": path.resolve(__dirname, "./src/store"),
    },
  },
});
