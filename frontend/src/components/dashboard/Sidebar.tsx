import { Component } from "solid-js";
 

const Sidebar: Component = () => {
  return (
   <div  class="menu bg-black h-[90vh] w-50">
    <h2 class="font-extrabold text-xl ml-12">ZENITH</h2>
   
    

      
  <nav class="mt-8">
    
    <a href="#" class="block px-3 py-2 rounded hover:bg-base-300">Dashboard</a>

    <details class="group">
      <summary class="flex items-center justify-between px-3 py-2 rounded cursor-pointer hover:bg-base-300">
        <span>Management</span>
        <svg class="w-4 h-4 transition-transform group-open:rotate-180" fill="none" stroke="currentColor"
          stroke-width="2" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
        </svg>
      </summary>
      <ul class="pl-4 mt-1 space-y-1">
        <li><a href="#" class="block px-2 py-1 rounded hover:bg-base-300">Users</a></li>
        <li><a href="#" class="block px-2 py-1 rounded hover:bg-base-300">Projects</a></li>
        <li><a href="#" class="block px-2 py-1 rounded hover:bg-base-300">Tasks</a></li>
      </ul>
    </details>

    <details class="group">
      <summary class="flex items-center justify-between px-3 py-2 rounded cursor-pointer hover:bg-base-300">
        <span>Settings</span>
        <svg class="w-4 h-4 transition-transform group-open:rotate-180" fill="none" stroke="currentColor"
          stroke-width="2" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
        </svg>
      </summary>
      <ul class="pl-4 mt-1 space-y-1">
        <li><a href="#" class="block px-2 py-1 rounded hover:bg-base-300">Profile</a></li>
        <li><a href="#" class="block px-2 py-1 rounded hover:bg-base-300">Security</a></li>
      </ul>
    </details>

    <a href="#" class="block px-3 py-2 rounded hover:bg-base-300">Logout</a>

  </nav>


   


    
  
    
    </div>
  );
};

export default Sidebar;
