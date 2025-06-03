import type { Component } from 'solid-js';
import Navbar from './Layout/Navbar';
import Sidebar from './Layout/Sidebar';

const App: Component = () => {
  return (
    <div class=' bg-amber-50'>
    <Navbar/>
    <Sidebar/>
    </div>
  );
};

export default App;
