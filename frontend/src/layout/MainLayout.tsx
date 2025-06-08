// src/layouts/MainLayout.tsx
import { from, JSX } from "solid-js";
import Sidebar from "../components/Sidebar";
import "../index.css";

export default function MainLayout(props: { children: JSX.Element }) {
  return (
    <div>
        <Sidebar/>
      <main class="p-4">{props.children}</main>
    </div>
  );
}
