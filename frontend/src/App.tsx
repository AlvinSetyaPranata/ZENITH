import { Router, Route } from "@solidjs/router";
import Dashboard from "./pages/Dashboard";
import Login from "./pages/Login";
import MainLayout from "./layout/MainLayout";
import AuthLayout from "./layout/AuthLayout";
import type { Component } from "solid-js";

const App: Component = () => {
  return (
    <Router>
      <Route path="/Dashboard" component={() => <MainLayout><Dashboard /></MainLayout>} />
      <Route path="/login" component={() => <AuthLayout><Login /></AuthLayout>} />
    </Router>
  );
};

export default App;
