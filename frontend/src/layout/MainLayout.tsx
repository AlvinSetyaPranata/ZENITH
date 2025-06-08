// src/layouts/MainLayout.tsx
import { from, JSX } from "solid-js";
import Sidebar from "../components/Sidebar";
import "../index.css";
import Navbar from "../components/Navbar";

export default function MainLayout(props: { children: JSX.Element }) {
  return (
    <div class="flex">
        <Sidebar/>
       
      <main class="flex-1 ">
         <Navbar/>
         {props.children}
        </main>

    </div>
  );
}
