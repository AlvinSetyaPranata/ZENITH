// src/layouts/AuthLayout.tsx
import { JSX } from "solid-js";
import "../index.css";

export default function AuthLayout(props: { children: JSX.Element }) {
  return (
    
      <div>
        {props.children}
      </div>
  
  );
}
