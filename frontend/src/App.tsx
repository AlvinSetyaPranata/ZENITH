import type { Component } from 'solid-js';
import Navbar from './Layout/Navbar';
import Sidebar from './Layout/Sidebar';

const App: Component = () => {
  return (
    <div class='max-w-full  bg-base-100'>
     <Navbar/>
      <Sidebar/>
    
   
    </div>
  );
};

export default App;
